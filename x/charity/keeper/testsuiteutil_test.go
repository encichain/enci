package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	enciapp "github.com/encichain/enci/app"
	charitytypes "github.com/encichain/enci/x/charity/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type CharityTestSuite struct {
	suite.Suite

	app *enciapp.EnciApp
	ctx sdk.Context
}

func (suite *CharityTestSuite) initKeepersWithmAccPerms() (authkeeper.AccountKeeper, bankkeeper.BaseKeeper) {
	app := suite.app
	maccPerms := enciapp.GetMaccPerms()
	appCodec := enciapp.MakeTestEncodingConfig().Codec
	blackListAddrs := map[string]bool{
		authtypes.FeeCollectorName:        true,
		stakingtypes.NotBondedPoolName:    true,
		stakingtypes.BondedPoolName:       true,
		distrtypes.ModuleName:             true,
		minttypes.ModuleName:              true,
		charitytypes.CharityCollectorName: true,
		charitytypes.BurnAccName:          false,
	}

	authKeeper := authkeeper.NewAccountKeeper(
		appCodec, app.GetKey(authtypes.StoreKey), app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount, maccPerms,
	)
	bankkeeper := bankkeeper.NewBaseKeeper(
		appCodec, app.GetKey(banktypes.StoreKey), authKeeper,
		app.GetSubspace(banktypes.ModuleName), blackListAddrs,
	)

	return authKeeper, bankkeeper
}

// returns context and app with params set on account keeper
func CreateTestApp(isCheckTx bool) (*enciapp.EnciApp, sdk.Context) {
	app := enciapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.CharityKeeper.SetParams(ctx, charitytypes.DefaultParamsSet)
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	app.MintKeeper.SetParams(ctx, minttypes.DefaultParams())
	return app, ctx
}

// SetupTest setups a new test, with new app, context, and antehandler.
func (suite *CharityTestSuite) SetupTest(isCheckTx bool) {
	//tempDir := suite.T().TempDir()
	suite.app, suite.ctx = CreateTestApp(isCheckTx)
	suite.ctx = suite.ctx.WithBlockHeight(1)
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")

	mintAcc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, minttypes.ModuleName)
	faucetAcc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, "faucet")
	charityTaxAcc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)

	suite.app.AccountKeeper.SetModuleAccount(suite.ctx, faucetAcc)
	suite.app.AccountKeeper.SetModuleAccount(suite.ctx, charityTaxAcc)
	suite.app.AccountKeeper.SetModuleAccount(suite.ctx, mintAcc)

	// Set up TxConfig.
	encodingConfig := enciapp.MakeTestEncodingConfig()
	// We're using TestMsg encoding in some tests, so register it here.
	encodingConfig.Amino.RegisterConcrete(&testdata.TestMsg{}, "testdata.TestMsg", nil)
	testdata.RegisterInterfaces(encodingConfig.InterfaceRegistry)

}

// FundModuleAccount is a utility function that funds a module account by
// minting and sending the coins to the address. This should be used for testing
// purposes only!
func FundModuleAccount(bankKeeper bankkeeper.Keeper, faucetName string, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, faucetName, amounts); err != nil {
		return sdkerrors.Wrap(err, "Could not mint coins")
	}

	err := bankKeeper.SendCoinsFromModuleToModule(ctx, faucetName, recipientMod, amounts)
	if err != nil {
		return sdkerrors.Wrap(err, "Could not fund module account with error")
	}
	return nil
}

func TestAnteTestSuite(t *testing.T) {
	suite.Run(t, new(CharityTestSuite))
}

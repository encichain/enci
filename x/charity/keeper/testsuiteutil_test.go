package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	enciapp "github.com/user/encichain/app"
	charitytypes "github.com/user/encichain/x/charity/types"
)

type CharityTestSuite struct {
	suite.Suite

	app *enciapp.EnciApp
	ctx sdk.Context
}

// returns context and app with params set on account keeper
func CreateTestApp(isCheckTx bool, tempDir string) (*enciapp.EnciApp, sdk.Context) {
	app := enciapp.Setup(isCheckTx, tempDir)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.CharityKeeper.SetParams(ctx, charitytypes.DefaultParamsSet)
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	return app, ctx
}

// SetupTest setups a new test, with new app, context, and antehandler.
func (suite *CharityTestSuite) SetupTest(isCheckTx bool) {
	tempDir := suite.T().TempDir()
	suite.app, suite.ctx = CreateTestApp(isCheckTx, tempDir)
	suite.ctx = suite.ctx.WithBlockHeight(1)

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
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, faucetName, recipientMod, amounts)
}

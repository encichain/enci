package keeper

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	auth "github.com/cosmos/cosmos-sdk/x/auth"
	//bank "github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	params "github.com/cosmos/cosmos-sdk/x/params"
	custombank "github.com/encichain/enci/customcore/bank"
	coretypes "github.com/encichain/enci/types"

	//charity "github.com/encichain/enci/x/charity"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	simparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	//charitykeeper "github.com/encichain/enci/x/charity/keeper"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	charitytypes "github.com/encichain/enci/x/charity/types"
)

const faucetAccount = "faucet"

func MakeTestCodec(t *testing.T) codec.Codec {
	return MakeEncodingConfig(t).Marshaler
}

var ModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	custombank.AppModuleBasic{},
	distr.AppModuleBasic{},
	staking.AppModuleBasic{},
	params.AppModuleBasic{},
)

var (
	ValPubKeys = simapp.CreateTestPubKeys(5)

	PubKeys = []crypto.PubKey{
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
		secp256k1.GenPrivKey().PubKey(),
	}

	Addrs = []sdk.AccAddress{
		sdk.AccAddress(PubKeys[0].Address()),
		sdk.AccAddress(PubKeys[1].Address()),
		sdk.AccAddress(PubKeys[2].Address()),
	}

	ValAddrs = []sdk.ValAddress{
		sdk.ValAddress(PubKeys[0].Address()),
		sdk.ValAddress(PubKeys[1].Address()),
		sdk.ValAddress(PubKeys[2].Address()),
	}

	InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	InitCoins  = sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, InitTokens))
)

func MakeEncodingConfig(_ *testing.T) simparams.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	std.RegisterInterfaces(interfaceRegistry)
	std.RegisterLegacyAminoCodec(amino)

	ModuleBasics.RegisterLegacyAminoCodec(amino)
	ModuleBasics.RegisterInterfaces(interfaceRegistry)

	return simparams.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}

type TestApp struct {
	Ctx           sdk.Context
	Cdc           *codec.LegacyAmino
	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
	CharityKeeper Keeper
	DistrKeeper   distrkeeper.Keeper
	StakingKeeper stakingkeeper.Keeper
	ParamsKeeper  paramskeeper.Keeper
}

func CreateKeeperTestApp(t *testing.T) TestApp {
	keyAcc := sdk.NewKVStoreKey(authtypes.StoreKey)
	keyBank := sdk.NewKVStoreKey(banktypes.StoreKey)
	keyParams := sdk.NewKVStoreKey(paramstypes.StoreKey)
	tKeyParams := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	keyStaking := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	keyDistr := sdk.NewKVStoreKey(distrtypes.StoreKey)
	ParamKeyCharity := sdk.NewKVStoreKey(charitytypes.StoreKey)
	memkeyCharity := sdk.NewKVStoreKey(charitytypes.MemStoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ctx := sdk.NewContext(ms, tmproto.Header{Time: time.Now().UTC()}, false, log.NewNopLogger())
	encodingConfig := MakeEncodingConfig(t)
	appCodec, legacyAmino := encodingConfig.Marshaler, encodingConfig.Amino

	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBank, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tKeyParams, sdk.StoreTypeTransient, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(ParamKeyCharity, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyStaking, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyDistr, sdk.StoreTypeIAVL, db)

	require.NoError(t, ms.LoadLatestVersion())

	blackListAddrs := map[string]bool{
		authtypes.FeeCollectorName:        true,
		stakingtypes.NotBondedPoolName:    true,
		stakingtypes.BondedPoolName:       true,
		distrtypes.ModuleName:             true,
		faucetAccount:                     true,
		minttypes.ModuleName:              true,
		charitytypes.CharityCollectorName: true,
		charitytypes.BurnAccName:          false,
	}

	maccPerms := map[string][]string{
		faucetAccount:                     {authtypes.Minter, authtypes.Burner},
		minttypes.ModuleName:              {authtypes.Minter},
		authtypes.FeeCollectorName:        nil,
		stakingtypes.NotBondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:       {authtypes.Burner, authtypes.Staking},
		distrtypes.ModuleName:             nil,
		charitytypes.ModuleName:           {authtypes.Burner, authtypes.Minter},
		charitytypes.CharityCollectorName: nil,
		charitytypes.BurnAccName:          {authtypes.Burner},
	}

	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, keyParams, tKeyParams)
	accountKeeper := authkeeper.NewAccountKeeper(appCodec, keyAcc, paramsKeeper.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms)
	bankKeeper := bankkeeper.NewBaseKeeper(appCodec, keyBank, accountKeeper, paramsKeeper.Subspace(banktypes.ModuleName), blackListAddrs)

	totalSupply := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, InitTokens.MulRaw(int64(len(Addrs)*14))))
	err := bankKeeper.MintCoins(ctx, faucetAccount, totalSupply)
	require.NoError(t, err)

	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keyStaking,
		accountKeeper,
		bankKeeper,
		paramsKeeper.Subspace(stakingtypes.ModuleName),
	)

	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = coretypes.MicroTokenDenom
	stakingKeeper.SetParams(ctx, stakingParams)

	distrKeeper := distrkeeper.NewKeeper(
		appCodec,
		keyDistr, paramsKeeper.Subspace(distrtypes.ModuleName),
		accountKeeper, bankKeeper, stakingKeeper,
		authtypes.FeeCollectorName, blackListAddrs)

	distrKeeper.SetFeePool(ctx, distrtypes.InitialFeePool())
	distrParams := distrtypes.DefaultParams()
	distrParams.CommunityTax = sdk.NewDecWithPrec(2, 2)
	distrParams.BaseProposerReward = sdk.NewDecWithPrec(1, 2)
	distrParams.BonusProposerReward = sdk.NewDecWithPrec(4, 2)
	distrKeeper.SetParams(ctx, distrParams)
	stakingKeeper.SetHooks(stakingtypes.NewMultiStakingHooks(distrKeeper.Hooks()))

	feeCollectorAcc := authtypes.NewEmptyModuleAccount(authtypes.FeeCollectorName)
	mintAcc := authtypes.NewEmptyModuleAccount(minttypes.ModuleName, authtypes.Minter)
	notBondedPool := authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName, authtypes.Burner, authtypes.Staking)
	bondPool := authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName, authtypes.Burner, authtypes.Staking)
	distrAcc := authtypes.NewEmptyModuleAccount(distrtypes.ModuleName)
	charityAcc := authtypes.NewEmptyModuleAccount(charitytypes.ModuleName, authtypes.Burner, authtypes.Minter)
	charityCollectorAcc := authtypes.NewEmptyModuleAccount(charitytypes.CharityCollectorName)
	burnAcc := authtypes.NewEmptyModuleAccount(charitytypes.BurnAccName, authtypes.Burner)

	accountKeeper.SetModuleAccount(ctx, feeCollectorAcc)
	accountKeeper.SetModuleAccount(ctx, mintAcc)
	accountKeeper.SetModuleAccount(ctx, bondPool)
	accountKeeper.SetModuleAccount(ctx, notBondedPool)
	accountKeeper.SetModuleAccount(ctx, distrAcc)
	accountKeeper.SetModuleAccount(ctx, charityAcc)
	accountKeeper.SetModuleAccount(ctx, charityCollectorAcc)
	accountKeeper.SetModuleAccount(ctx, burnAcc)

	faucetAddr := accountKeeper.GetModuleAccount(ctx, faucetAccount)

	bal := bankKeeper.HasBalance(ctx, faucetAddr.GetAddress(), sdk.Coin{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(int64(1))})
	if bal {
		err := bankKeeper.MintCoins(ctx, faucetAccount, InitCoins)
		require.NoError(t, err)
	}
	// Test charity collector
	err = bankKeeper.SendCoinsFromModuleToModule(ctx, faucetAccount, charitytypes.CharityCollectorName, InitCoins)
	require.NoError(t, err)

	err = bankKeeper.SendCoinsFromModuleToModule(ctx, faucetAccount, stakingtypes.NotBondedPoolName, sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, InitTokens.MulRaw(int64(len(Addrs)+1)))))
	require.NoError(t, err)

	/*
		// Test charity burn
		err = bankKeeper.SendCoinsFromModuleToModule(ctx, faucetAccount, types.BurnAccName, InitCoins)
		require.NoError(t, err)
		err = bankKeeper.BurnCoins(ctx, charitytypes.BurnAccName, InitCoins)
		require.NoError(t, err)
	*/
	for _, addr := range Addrs {
		accountKeeper.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(addr))
		err := bankKeeper.SendCoinsFromModuleToAccount(ctx, faucetAccount, addr, InitCoins)
		require.NoError(t, err)
	}

	accountKeeper.SetModuleAccount(ctx, burnAcc)

	charityKeeper := NewKeeper(
		appCodec,
		ParamKeyCharity,
		memkeyCharity,
		bankKeeper,
		accountKeeper,
		paramsKeeper.Subspace(charitytypes.ModuleName),
	)

	charityKeeper.SetParams(ctx, charitytypes.DefaultParams())

	return TestApp{ctx, legacyAmino, accountKeeper, bankKeeper, *charityKeeper, distrKeeper, stakingKeeper, paramsKeeper}
}

//Taken from params proposal_handler.go. Attempts to update the params with the values in ParameterChangeProposal{}
func handleParameterChangeProposal(ctx sdk.Context, k paramskeeper.Keeper, p *proposal.ParameterChangeProposal) error {
	for _, c := range p.Changes {
		ss, ok := k.GetSubspace(c.Subspace)
		if !ok {
			return sdkerrors.Wrap(proposal.ErrUnknownSubspace, c.Subspace)
		}

		k.Logger(ctx).Info(
			fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
		)

		if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
			return sdkerrors.Wrapf(proposal.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
		}
	}

	return nil
}

// FundAccount is a utility function that funds an account by minting and
// sending the coins to the address. This should be used for testing purposes
// only!
//

func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, faucetAccount, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, faucetAccount, addr, amounts)
}

func CoreFundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

// FundModuleAccount is a utility function that funds a module account by
// minting and sending the coins to the address. This should be used for testing
// purposes only!
func FundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, faucetAccount, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, faucetAccount, recipientMod, amounts)
}

func CoreFundModuleAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToModule(ctx, minttypes.ModuleName, recipientMod, amounts)
}

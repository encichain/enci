package charity_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"
	charity "github.com/user/encichain/x/charity"
	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	coreapp "github.com/user/encichain/app"
)

// module test
func TestCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := coreapp.Setup(false)
	ctx := app.GetBaseApp().NewContext(false, tmproto.Header{})
	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)
	cAcc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.CharityCollectorName))
	bAcc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.BurnAccName))
	require.NotNil(t, cAcc)
	require.NotNil(t, bAcc)
}

func TestBurnEndblock(t *testing.T) {
	app := coreapp.Setup(false)
	ctx := app.GetBaseApp().NewContext(false, tmproto.Header{})
	// Burner account balance should be 0
	burnAddr := app.AccountKeeper.GetModuleAddress(types.BurnAccName)
	require.True(t, app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero())

	// Fund burner account
	testCoins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 1500))
	err := keeper.CoreFundModuleAccount(app.BankKeeper, ctx, types.BurnAccName, testCoins)
	require.NoError(t, err)
	require.True(t, app.BankKeeper.HasBalance(ctx, burnAddr, testCoins[0]))

	ctx = ctx.WithBlockHeight(int64(coretypes.BlocksPerPeriod - 1))
	// Call endblock
	charity.EndBlocker(ctx, app.CharityKeeper)
	require.True(t, app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero())
}

func TestUserSendBurnEndBlock(t *testing.T) {
	app := coreapp.Setup(false)
	ctx := app.GetBaseApp().NewContext(false, tmproto.Header{})
	// Burner account balance should be 0
	burnAddr := app.AccountKeeper.GetModuleAddress(types.BurnAccName)
	require.True(t, app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero())

	// Set up account and fund it
	_, _, addr1 := testdata.KeyTestPubAddr()
	testCoins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 1500))
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	require.NotNil(t, acc1)
	err := keeper.CoreFundAccount(app.BankKeeper, ctx, acc1.GetAddress(), testCoins)
	require.NoError(t, err)
	require.True(t, app.BankKeeper.HasBalance(ctx, acc1.GetAddress(), testCoins[0]))

	// Send coins from account to burn module
	err = app.BankKeeper.SendCoinsFromAccountToModule(ctx, acc1.GetAddress(), types.BurnAccName, testCoins)
	require.NoError(t, err)
	require.True(t, app.BankKeeper.HasBalance(ctx, burnAddr, testCoins[0]))

	// Set block height to end of period and call EndBlocker
	ctx = ctx.WithBlockHeight(int64(coretypes.BlocksPerPeriod) - 1)
	charity.EndBlocker(ctx, app.CharityKeeper)

	require.True(t, app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero())
}

func TestEndBlocker(t *testing.T) {
	app := keeper.CreateKeeperTestApp(t)
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")
	// Configure test charity accounts
	bech32addr1 := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	bech32addr2 := "enci1vvcw744ck9kzczrhf282lqmset47jnxe9090qt"
	addr1, err := sdk.AccAddressFromBech32(bech32addr1)
	require.NoError(t, err)
	addr2, err := sdk.AccAddressFromBech32(bech32addr2)
	require.NoError(t, err)
	acc1 := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr1)
	require.NotNil(t, acc1)
	acc2 := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr2)
	require.NotNil(t, acc2)
	taxAddr := app.AccountKeeper.GetModuleAddress(types.CharityCollectorName)

	// Set accounts to store
	app.AccountKeeper.SetAccount(app.Ctx, acc1)
	app.AccountKeeper.SetAccount(app.Ctx, acc2)
	// Create checksums and encode to strings
	checksum1 := keeper.CreateCharitySha256("Test Charity", bech32addr1)
	require.NotEqual(t, "", checksum1)

	checksum2 := keeper.CreateCharitySha256("Test Charity 2", bech32addr2)

	// Set params and target charities
	params := types.DefaultParams()
	params.TaxCaps = []types.TaxCap{{
		Denom: "uenci",
		Cap:   sdk.NewInt(int64(2000000)),
	}}
	params.Charities = []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr1,
			Checksum:   checksum1,
		},
		{CharityName: "Test Charity 2",
			AccAddress: bech32addr2,
			Checksum:   checksum2},
	}

	app.CharityKeeper.SetParams(app.Ctx, params)
	require.Equal(t, app.CharityKeeper.GetAllParams(app.Ctx), params)
	charities := app.CharityKeeper.GetCharity(app.Ctx)
	require.Equal(t, params.Charities, charities)

	// Make sure a charity account has 0 balance
	zeroBal1 := app.BankKeeper.GetAllBalances(app.Ctx, addr1).IsZero()
	zeroBal2 := app.BankKeeper.GetAllBalances(app.Ctx, addr2).IsZero()
	require.True(t, zeroBal1)
	require.True(t, zeroBal2)
	// Burner account balance should be 0
	burnAddr := app.AccountKeeper.GetModuleAddress(types.BurnAccName)
	require.True(t, app.BankKeeper.GetAllBalances(app.Ctx, burnAddr).IsZero())

	// Set test TaxProceeds to store
	proceeds := sdk.NewInt(int64(50000000))
	app.CharityKeeper.AddTaxProceeds(app.Ctx, sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: proceeds}})
	require.Equal(t, sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: proceeds}}, app.CharityKeeper.GetTaxProceeds(app.Ctx))

	// Make sure store taxcap is set to default taxcap
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, coretypes.MicroTokenDenom))

	// Set blockheight to end of period
	app.Ctx = app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerPeriod - 1))
	// Get balances and calculate burn amount and post deduction balance
	charityTaxBal := app.BankKeeper.GetAllBalances(app.Ctx, taxAddr)
	require.Equal(t, keeper.InitCoins, charityTaxBal)
	burnAmt := app.CharityKeeper.CalculateBurnAmount(app.Ctx, charityTaxBal)
	afterBurnBal := charityTaxBal.Sub(burnAmt)

	// Call EndBlocker
	charity.EndBlocker(app.Ctx, app.CharityKeeper)

	// Check if target charity accounts have received donation
	hasbal := app.BankKeeper.HasBalance(app.Ctx, addr1, sdk.NewCoin(coretypes.MicroTokenDenom, afterBurnBal[0].Amount.Quo(sdk.NewInt(int64(2)))))
	hasbal2 := app.BankKeeper.HasBalance(app.Ctx, addr2, sdk.NewCoin(coretypes.MicroTokenDenom, afterBurnBal[0].Amount.Quo(sdk.NewInt(int64(2)))))
	require.True(t, hasbal)
	require.True(t, hasbal2)

	// Verify burn amount has been deducted by ensuring charity recipient balance < pre-burn charity collector amount
	require.False(t, app.BankKeeper.HasBalance(app.Ctx, addr1, sdk.NewCoin(coretypes.MicroTokenDenom, keeper.InitTokens.Quo(sdk.NewInt(int64(2))))))
	require.False(t, app.BankKeeper.HasBalance(app.Ctx, addr2, sdk.NewCoin(coretypes.MicroTokenDenom, keeper.InitTokens.Quo(sdk.NewInt(int64(2))))))

	// Verify that burn module account is now zero in balance
	require.True(t, app.BankKeeper.GetAllBalances(app.Ctx, burnAddr).IsZero())

	// Check if Payouts have been created and set to store under *period*
	require.Equal(t, []types.Payout{
		{Recipientaddr: bech32addr1, Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: afterBurnBal[0].Amount.Quo(sdk.NewInt(int64(2)))}}},
		{Recipientaddr: bech32addr2, Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: afterBurnBal[0].Amount.Quo(sdk.NewInt(int64(2)))}}}},
		app.CharityKeeper.GetPayouts(app.Ctx, app.CharityKeeper.GetCurrentPeriod(app.Ctx)))

	// Check if taxproceeds have been stored under current *period*
	periodproceeds := sdk.NewCoins(sdk.Coin{Denom: coretypes.MicroTokenDenom, Amount: proceeds})
	require.True(t, app.CharityKeeper.GetPeriodTaxProceeds(app.Ctx, app.CharityKeeper.GetCurrentPeriod(app.Ctx))[0].Amount.GTE(periodproceeds[0].Amount))

	// Check if store taxcaps have been synced with param store taxcaps
	require.Equal(t, app.CharityKeeper.GetParamTaxCaps(app.Ctx), app.CharityKeeper.GetTaxCaps(app.Ctx))
	require.Equal(t, params.TaxCaps, app.CharityKeeper.GetParamTaxCaps(app.Ctx))

	// Check if store tax proceeds have been reset
	require.Equal(t, sdk.Coins{}.IsZero(), app.CharityKeeper.GetTaxProceeds(app.Ctx).IsZero())
}

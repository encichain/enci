package charity

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"

	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"
)

func TestEndBlocker(t *testing.T) {
	app := keeper.CreateTestApp(t)
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

	// Set accounts to store
	app.AccountKeeper.SetAccount(app.Ctx, acc1)
	app.AccountKeeper.SetAccount(app.Ctx, acc2)
	// Create checksums and encode to strings
	csb1 := sha256.Sum256([]byte("Test Charity" + bech32addr1))
	checksum1 := hex.EncodeToString(csb1[:])
	require.NotEqual(t, "", checksum1)

	csb2 := sha256.Sum256([]byte("Test Charity 2" + bech32addr2))
	checksum2 := hex.EncodeToString(csb2[:])

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
	hasbal := app.BankKeeper.HasBalance(app.Ctx, addr1, sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(1000))))
	hasbal2 := app.BankKeeper.HasBalance(app.Ctx, addr2, sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(1000))))
	require.False(t, hasbal)
	require.False(t, hasbal2)
	// Set test TaxProceeds to store
	proceeds := sdk.NewInt(int64(50000000))
	app.CharityKeeper.AddTaxProceeds(app.Ctx, sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: proceeds}})
	require.Equal(t, sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: proceeds}}, app.CharityKeeper.GetTaxProceeds(app.Ctx))

	// Make sure store taxcap is set to default taxcap
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, coretypes.MicroTokenDenom))

	// Set blockheight to end of period and call Charity EndBlocker
	app.Ctx = app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerPeriod - 1))
	require.Equal(t, keeper.InitCoins, app.BankKeeper.GetAllBalances(app.Ctx, app.AccountKeeper.GetModuleAddress(types.CharityCollectorName)))

	EndBlocker(app.Ctx, app.CharityKeeper)

	// Check if target charity accounts have received donation
	hasbal = app.BankKeeper.HasBalance(app.Ctx, addr1, sdk.NewCoin(coretypes.MicroTokenDenom, keeper.InitTokens.Quo(sdk.NewInt(int64(2)))))
	hasbal2 = app.BankKeeper.HasBalance(app.Ctx, addr2, sdk.NewCoin(coretypes.MicroTokenDenom, keeper.InitTokens.Quo(sdk.NewInt(int64(2)))))
	require.True(t, hasbal)
	require.True(t, hasbal2)

	// Check if Payouts have been created and set to store under *period*
	require.Equal(t, []types.Payout{
		{Recipientaddr: bech32addr1, Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: keeper.InitTokens.Quo(sdk.NewInt(int64(2)))}}},
		{Recipientaddr: bech32addr2, Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: keeper.InitTokens.Quo(sdk.NewInt(int64(2)))}}}},
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

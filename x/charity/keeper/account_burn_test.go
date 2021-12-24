package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/encichain/enci/types"
	"github.com/stretchr/testify/require"

	"github.com/encichain/enci/x/charity/types"
	//"go.uber.org/goleak"
)

func TestBurnCoins(t *testing.T) {
	app := CreateKeeperTestApp(t)
	// Get burner account address and ensure it has no balance
	burnAddr := app.AccountKeeper.GetModuleAddress(types.BurnAccName)
	isZeroBal := app.BankKeeper.GetAllBalances(app.Ctx, burnAddr).IsZero()
	require.True(t, isZeroBal)

	//Try burning with zero balance
	err := app.CharityKeeper.BurnCoinsFromBurner(app.Ctx)
	require.NoError(t, err)

	// Fund burner account
	coins := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(10000000))))
	err = FundModuleAccount(app.BankKeeper, app.Ctx, types.BurnAccName, coins)
	require.NoError(t, err)
	hasBal := app.BankKeeper.HasBalance(app.Ctx, burnAddr, sdk.NewCoin(coretypes.MicroTokenDenom, coins[0].Amount))
	require.True(t, hasBal)

	// Test BurnCoinsFromBurner
	err = app.CharityKeeper.BurnCoinsFromBurner(app.Ctx)
	require.NoError(t, err)
	isZeroBal = app.BankKeeper.GetAllBalances(app.Ctx, burnAddr).IsZero()
	require.True(t, isZeroBal)
}

func TestCalculateBurnAmount(t *testing.T) {
	app := CreateKeeperTestApp(t)
	ctx := app.Ctx
	app.CharityKeeper.SetParams(ctx, types.DefaultParams())
	require.Equal(t, types.DefaultBurnRate, app.CharityKeeper.GetBurnRate(ctx))

	for i := int64(1); i < 5; i++ {
		app.CharityKeeper.SetBurnRate(ctx, sdk.NewDecWithPrec(i, 2))
		burnRate := app.CharityKeeper.GetBurnRate(ctx)
		require.Equal(t, sdk.NewDecWithPrec(i, 2), burnRate)

		balance := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(100000)).Mul(sdk.NewInt(i))))
		expBurnAmt := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewDecFromInt(balance[0].Amount).Mul(burnRate).TruncateInt()))
		burnAmt := app.CharityKeeper.CalculateBurnAmount(ctx, balance)

		require.Equal(t, expBurnAmt, burnAmt)
	}

	//0 balance
	burnAmt := app.CharityKeeper.CalculateBurnAmount(ctx, sdk.NewCoins())
	require.Equal(t, sdk.Coins{}, burnAmt)
}

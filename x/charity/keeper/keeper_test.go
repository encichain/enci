package keeper

import (
	"testing"
	//"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
	"github.com/user/encichain/x/charity/types"

	coretypes "github.com/user/encichain/types"
)

func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	keeper := NewKeeper(
		codec.NewProtoCodec(registry),
		storeKey,
		memStoreKey,
		nil,
		nil,
		paramstypes.Subspace{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return keeper, ctx
}

func TestTaxRateLimits(t *testing.T) {
	app := CreateTestApp(t)

	for i := int64(0); i < 10; i++ {
		app.CharityKeeper.SetTaxRateLimits(app.Ctx, types.TaxRateLimits{
			RateMin: sdk.NewDecWithPrec(i, 3),
			RateMax: sdk.NewDecWithPrec(i, 2),
		},
		)
		require.Equal(t, types.TaxRateLimits{
			RateMin: sdk.NewDecWithPrec(i, 3),
			RateMax: sdk.NewDecWithPrec(i, 2),
		}, app.CharityKeeper.GetTaxRateLimits(app.Ctx))
	}
}

func TestTaxCap(t *testing.T) {
	app := CreateTestApp(t)

	for i := int64(0); i < 10; i++ {
		app.CharityKeeper.SetTaxCap(app.Ctx, coretypes.MicroTokenDenom, sdk.NewInt(i))
		require.Equal(t, sdk.NewInt(i), app.CharityKeeper.GetTaxCap(app.Ctx, coretypes.MicroTokenDenom))
	}

}

func TestIterateTaxCap(t *testing.T) {
	input := CreateTestApp(t)

	uenciCap := sdk.NewInt(1000000)
	input.CharityKeeper.SetTaxCap(input.Ctx, coretypes.MicroTokenDenom, uenciCap)
	require.Equal(t, input.CharityKeeper.GetTaxCap(input.Ctx, coretypes.MicroTokenDenom), uenciCap)

	input.CharityKeeper.IterateTaxCaps(input.Ctx, func(denom string, taxCap sdk.Int) bool {
		if denom == coretypes.MicroTokenDenom {
			require.Equal(t, uenciCap, taxCap)
		}
		return true
	})

}

func TestParams(t *testing.T) {
	input := CreateTestApp(t)

	defaultParams := types.DefaultParams()
	input.CharityKeeper.SetParams(input.Ctx, defaultParams)

	getParams := input.CharityKeeper.GetAllParams(input.Ctx)
	require.Equal(t, defaultParams, getParams)
}

func TestTaxProceeds(t *testing.T) {
	input := CreateTestApp(t)

	for i := int64(0); i < 10; i++ {
		proceeds := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(100+i)))
		for j := 0; j < 3; j++ {
			input.CharityKeeper.AddTaxProceeds(input.Ctx, proceeds)
		}

		require.Equal(t, proceeds.Add(proceeds...).Add(proceeds...), input.CharityKeeper.GetTaxProceeds(input.Ctx))
		require.False(t, input.CharityKeeper.GetTaxProceeds(input.Ctx).IsZero())
		input.CharityKeeper.SetTaxProceeds(input.Ctx, sdk.Coins{})
		require.True(t, input.CharityKeeper.GetTaxProceeds(input.Ctx).IsZero())
	}

	proceeds := sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(100)}}
	input.CharityKeeper.SetTaxProceeds(input.Ctx, proceeds)
	require.Equal(t, proceeds, input.CharityKeeper.GetTaxProceeds(input.Ctx))

	// Test AddTaxProceed single case
	input.CharityKeeper.SetTaxProceeds(input.Ctx, proceeds)
	input.CharityKeeper.AddTaxProceeds(input.Ctx, proceeds)
	require.Equal(t, proceeds.Add(proceeds...), input.CharityKeeper.GetTaxProceeds(input.Ctx))
	require.False(t, input.CharityKeeper.GetTaxProceeds(input.Ctx).IsZero())
}

func TestPeriodTaxProceeds(t *testing.T) {
	testApp := CreateTestApp(t)

	for i := int64(0); i < 10; i++ {
		// Set TaxProceed to store
		testApp.CharityKeeper.SetPeriodTaxProceeds(testApp.Ctx, i, sdk.Coins{
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*20000 + 100)},
		})

		// Try to get unset TaxProceed, should return sdk.Coins{}
		require.Equal(t, sdk.Coins{}, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i+1))

		require.NoError(t, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i).Validate())
		// Check if Get method retrieves valid set TaxProceed
		require.Equal(t, sdk.Coins{
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*20000 + 100)},
		}, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i))

		require.NotEqual(t, sdk.Coins{}, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i))
		require.False(t, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i).IsZero())
	}
	// Try clearing Tax Proceeds from store
	testApp.CharityKeeper.ClearPeriodTaxProceeds(testApp.Ctx)

	for i := int64(0); i < 10; i++ {
		require.Equal(t, sdk.Coins{}, testApp.CharityKeeper.GetPeriodTaxProceeds(testApp.Ctx, i))
	}

}

func TestPayouts(t *testing.T) {
	app := CreateTestApp(t)

	for i := int64(0); i < 10; i++ {
		payouts := []types.Payout{
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1), Recipientaddr: "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test"},
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1000), Recipientaddr: "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test2"},
		}

		app.CharityKeeper.SetPayouts(app.Ctx, i, payouts)
		require.Equal(t, payouts, app.CharityKeeper.GetPayouts(app.Ctx, i))
	}
	require.NotEqual(t, []types.Payout{}, app.CharityKeeper.GetPayouts(app.Ctx, 1))

	// Try clearing []Payout from store
	app.CharityKeeper.ClearPayouts(app.Ctx)

	for i := int64(0); i < 10; i++ {
		require.Equal(t, []types.Payout{}, app.CharityKeeper.GetPayouts(app.Ctx, i))
	}

}

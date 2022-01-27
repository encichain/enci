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
	"github.com/encichain/enci/x/charity/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	coretypes "github.com/encichain/enci/types"
)

func setupCharityKeeper(t testing.TB) (*Keeper, sdk.Context) {
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

	// Initialize params
	keeper.SetParams(ctx, types.DefaultParams())

	return keeper, ctx
}

func TestGetCurrentEpoch(t *testing.T) {
	app := CreateKeeperTestApp(t)
	for i := int64(0); i < 10; i++ {
		ctx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerEpoch) * i)
		epoch := app.CharityKeeper.GetCurrentEpoch(ctx)
		require.Equal(t, (ctx.BlockHeight() / int64(coretypes.BlocksPerEpoch)), epoch)
	}
}

func TestTaxRateLimits(t *testing.T) {
	app := CreateKeeperTestApp(t)

	for i := int64(0); i < 10; i++ {
		app.CharityKeeper.SetTaxRateLimits(app.Ctx, types.TaxRateLimits{
			RateMin:     sdk.NewDecWithPrec(i, 3),
			TaxRateMax:  sdk.NewDecWithPrec(i, 2),
			BurnRateMax: sdk.NewDecWithPrec(i, 2),
		},
		)
		require.Equal(t, types.TaxRateLimits{
			RateMin:     sdk.NewDecWithPrec(i, 3),
			TaxRateMax:  sdk.NewDecWithPrec(i, 2),
			BurnRateMax: sdk.NewDecWithPrec(i, 2),
		}, app.CharityKeeper.GetTaxRateLimits(app.Ctx))
	}
}

func TestTaxCap(t *testing.T) {
	app := CreateKeeperTestApp(t)

	for i := int64(0); i < 10; i++ {
		app.CharityKeeper.SetTaxCap(app.Ctx, coretypes.MicroTokenDenom, sdk.NewInt(i))
		require.Equal(t, sdk.NewInt(i), app.CharityKeeper.GetTaxCap(app.Ctx, coretypes.MicroTokenDenom))
	}
}

func TestIterateTaxCap(t *testing.T) {
	app := CreateKeeperTestApp(t)

	uenciCap := sdk.NewInt(1000000)
	app.CharityKeeper.SetTaxCap(app.Ctx, coretypes.MicroTokenDenom, uenciCap)
	require.Equal(t, app.CharityKeeper.GetTaxCap(app.Ctx, coretypes.MicroTokenDenom), uenciCap)

	app.CharityKeeper.IterateTaxCaps(app.Ctx, func(denom string, taxCap sdk.Int) bool {
		if denom == coretypes.MicroTokenDenom {
			require.Equal(t, uenciCap, taxCap)
		}
		return true
	})
}

func TestGetTaxCaps(t *testing.T) {
	app := CreateKeeperTestApp(t)
	app.CharityKeeper.SetTaxCap(app.Ctx, coretypes.MicroTokenDenom, types.DefaultCap)
	taxcaps := app.CharityKeeper.GetTaxCaps(app.Ctx)

	require.Equal(t, []types.TaxCap{{Denom: coretypes.MicroTokenDenom, Cap: types.DefaultCap}}, taxcaps)
}

func TestClearTaxCaps(t *testing.T) {
	app := CreateKeeperTestApp(t)
	defaultCap := sdk.NewInt(int64(2000000))
	testTaxCaps := []types.TaxCap{{Denom: "uenci", Cap: defaultCap}, {Denom: "menci", Cap: defaultCap}, {Denom: "enci", Cap: defaultCap}}
	// Set taxcaps to store
	for _, taxcap := range testTaxCaps {
		app.CharityKeeper.SetTaxCap(app.Ctx, taxcap.Denom, taxcap.Cap)
	}

	require.Equal(t, defaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "uenci"))
	require.Equal(t, defaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "menci"))
	require.Equal(t, defaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "enci"))
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "Nonexistentdenom"))

	//Clear taxcaps from store
	app.CharityKeeper.ClearTaxCaps(app.Ctx)
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "uenci"))
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "menci"))
	require.Equal(t, types.DefaultCap, app.CharityKeeper.GetTaxCap(app.Ctx, "enci"))
}

func TestParams(t *testing.T) {
	app := CreateKeeperTestApp(t)

	defaultParams := types.DefaultParams()
	app.CharityKeeper.SetParams(app.Ctx, defaultParams)

	getParams := app.CharityKeeper.GetAllParams(app.Ctx)
	require.Equal(t, defaultParams, getParams)
}

func TestTaxProceeds(t *testing.T) {
	app := CreateKeeperTestApp(t)

	for i := int64(0); i < 10; i++ {
		proceeds := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(100+i)))
		for j := 0; j < 3; j++ {
			app.CharityKeeper.RecordTaxProceeds(app.Ctx, proceeds)
		}

		require.Equal(t, proceeds.Add(proceeds...).Add(proceeds...), app.CharityKeeper.GetTaxProceeds(app.Ctx))
		require.False(t, app.CharityKeeper.GetTaxProceeds(app.Ctx).IsZero())
		app.CharityKeeper.SetTaxProceeds(app.Ctx, sdk.Coins{})
		require.True(t, app.CharityKeeper.GetTaxProceeds(app.Ctx).IsZero())
	}

	proceeds := sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(100)}}
	app.CharityKeeper.SetTaxProceeds(app.Ctx, proceeds)
	require.Equal(t, proceeds, app.CharityKeeper.GetTaxProceeds(app.Ctx))

	// Test RecordTaxProceeds single case
	app.CharityKeeper.SetTaxProceeds(app.Ctx, proceeds)
	app.CharityKeeper.RecordTaxProceeds(app.Ctx, proceeds)
	require.Equal(t, proceeds.Add(proceeds...), app.CharityKeeper.GetTaxProceeds(app.Ctx))
	require.False(t, app.CharityKeeper.GetTaxProceeds(app.Ctx).IsZero())
}

func TestEpochTaxProceeds(t *testing.T) {
	testApp := CreateKeeperTestApp(t)

	for i := int64(0); i < 10; i++ {
		// Set TaxProceed to store
		testApp.CharityKeeper.SetEpochTaxProceeds(testApp.Ctx, i, sdk.Coins{
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*20000 + 100)},
		})

		// Try to get unset TaxProceed, should return sdk.Coins{}
		require.Equal(t, sdk.Coins{}, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i+1))
		require.NoError(t, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i).Validate())

		// Check if Get method retrieves valid set TaxProceed
		require.Equal(t, sdk.Coins{
			{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*20000 + 100)},
		}, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i))

		require.NotEqual(t, sdk.Coins{}, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i))
		require.False(t, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i).IsZero())
	}
	// Try clearing Tax Proceeds from store
	testApp.CharityKeeper.ClearEpochTaxProceeds(testApp.Ctx)

	for i := int64(0); i < 10; i++ {
		require.Equal(t, sdk.Coins{}, testApp.CharityKeeper.GetEpochTaxProceeds(testApp.Ctx, i))
	}
}

func TestPayouts(t *testing.T) {
	app := CreateKeeperTestApp(t)
	addr1 := "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test"
	addr2 := "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test2"
	for i := int64(0); i < 10; i++ {
		payouts := []types.Payout{
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1)}}, Recipientaddr: addr1},
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1000)}}, Recipientaddr: addr2},
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

func TestGetCollectionEpochs(t *testing.T) {
	app := CreateKeeperTestApp(t)
	addr1 := "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test"
	addr2 := "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test2"
	for i := int64(0); i < 10; i++ {
		payouts := []types.Payout{
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1)}}, Recipientaddr: addr1},
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1000)}}, Recipientaddr: addr2},
		}
		app.CharityKeeper.SetPayouts(app.Ctx, i, payouts)

		taxproceeds := sdk.Coins{sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64((i+1)*1000)))}
		app.CharityKeeper.SetEpochTaxProceeds(app.Ctx, i, taxproceeds)
	}
	// Create expected []types.CollectionEpoch{} value
	expectedval := []types.CollectionEpoch{}
	for i := int64(0); i < 10; i++ {
		payouts := []types.Payout{
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1)}}, Recipientaddr: addr1},
			{Coins: sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(i*sdk.DefaultPowerReduction.Int64() + 1000)}}, Recipientaddr: addr2},
		}
		taxproceeds := sdk.Coins{sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64((i+1)*1000)))}
		expectedval = append(expectedval, types.CollectionEpoch{Epoch: uint64(i), TaxCollected: taxproceeds, Payouts: payouts})
	}
	ctx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerEpoch * 10))
	require.Equal(t, expectedval, app.CharityKeeper.GetCollectionEpochs(ctx))
}

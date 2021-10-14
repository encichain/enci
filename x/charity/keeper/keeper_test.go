package keeper

import (
	"testing"

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
	input := CreateTestApp(t)

	// See that we can get and set tax rate
	for i := int64(0); i < 10; i++ {
		input.CharityKeeper.SetTaxRateLimits(input.Ctx, types.TaxRateLimits{
			RateMin: sdk.NewDecWithPrec(i, 3),
			RateMax: sdk.NewDecWithPrec(i, 2),
		},
		)
		require.Equal(t, types.TaxRateLimits{
			RateMin: sdk.NewDecWithPrec(i, 3),
			RateMax: sdk.NewDecWithPrec(i, 2),
		}, input.CharityKeeper.GetTaxRateLimits(input.Ctx))
	}
}

func TestIterateTaxCap(t *testing.T) {
	input := CreateTestApp(t)

	uenciCap := sdk.NewInt(1000000)
	input.CharityKeeper.SetTaxCap(input.Ctx, coretypes.MicroTokenDenom, uenciCap)

	input.CharityKeeper.IterateTaxCaps(input.Ctx, func(denom string, taxCap sdk.Int) bool {
		if denom == coretypes.MicroTokenDenom {
			require.Equal(t, uenciCap, taxCap)
		}
		return true
	})

}

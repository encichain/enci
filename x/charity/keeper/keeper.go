package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/types"

	// this line is used by starport scaffolding # ibc/keeper/import

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		paramStore    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
	paramStore paramstypes.Subspace,
) *Keeper {

	// set KeyTable if it has not already been set
	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	}

	// Check if charity tax collector address is set. Panic if nil
	if collectaddr := accountKeeper.GetModuleAddress(types.CharityCollectorName); collectaddr == nil {
		panic(fmt.Sprintf("%s module account not set", types.CharityCollectorName))
	}

	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		paramStore:    paramStore,
	}
}

// Logger returns a module specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetTaxRateLimits gets the tax rate limits
func (k Keeper) GetTaxRateLimits(ctx sdk.Context) types.TaxRateLimits {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.KeyTaxRateLimits)
	if b == nil {
		return types.DefaultTaxRateLimits
	}

	taxlimits := types.TaxRateLimits{}
	k.cdc.MustUnmarshal(b, &taxlimits)

	return taxlimits
}

// SetTaxRate sets the TaxRate in the store
func (k Keeper) SetTaxRateLimits(ctx sdk.Context, taxratelimits types.TaxRateLimits) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&types.TaxRateLimits{
		RateMin: taxratelimits.RateMin,
		RateMax: taxratelimits.RateMax,
	})

	// Set the store
	store.Set(types.KeyTaxRateLimits, b)
}

// GetCurrentPeriod calculates the current CollectionPeriod period by dividing current Block height by a Block week.
func GetCurrentPeriod(ctx sdk.Context) int64 {
	return (ctx.BlockHeight() / int64(coretypes.BlocksPerWeek))
}

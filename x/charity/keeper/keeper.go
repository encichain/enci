package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity/types"

	// this line is used by starport scaffolding # ibc/keeper/import

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		BankKeeper    types.BankKeeper
		AccountKeeper types.AccountKeeper
		paramStore    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	BankKeeper types.BankKeeper,
	AccountKeeper types.AccountKeeper,
	paramStore paramstypes.Subspace,
) *Keeper {

	// set KeyTable if it has not already been set
	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	}

	// Check if charity tax collector address is set. Panic if nil
	if collectAddr := AccountKeeper.GetModuleAddress(types.CharityCollectorName); collectAddr == nil {
		panic(fmt.Sprintf("%s module account not set", types.CharityCollectorName))
	}

	// Check if burner account address is set. Panic if nil
	if burnerAddr := AccountKeeper.GetModuleAddress(types.BurnAccName); burnerAddr == nil {
		panic(fmt.Sprintf("%s module account not set", types.BurnAccName))
	}

	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		BankKeeper:    BankKeeper,
		AccountKeeper: AccountKeeper,
		paramStore:    paramStore,
	}
}

// Logger returns a module specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetCurrentEpoch calculates the current CollectionEpoch epoch by dividing current Block height by a Block week.
func (k Keeper) GetCurrentEpoch(ctx sdk.Context) int64 {
	return (ctx.BlockHeight() / int64(coretypes.BlocksPerEpoch))
}

// GetTaxRateLimits gets the tax rate limits
func (k Keeper) GetTaxRateLimits(ctx sdk.Context) types.TaxRateLimits {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TaxRateLimitsKey)
	if bz == nil {
		return types.DefaultTaxRateLimits
	}

	taxlimits := types.TaxRateLimits{}
	k.cdc.MustUnmarshal(bz, &taxlimits)

	return taxlimits
}

// SetTaxRate sets the TaxRate in the store
func (k Keeper) SetTaxRateLimits(ctx sdk.Context, taxratelimits types.TaxRateLimits) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.TaxRateLimits{
		RateMin:     taxratelimits.RateMin,
		TaxRateMax:  taxratelimits.TaxRateMax,
		BurnRateMax: taxratelimits.BurnRateMax,
	})

	// Set the store
	store.Set(types.TaxRateLimitsKey, bz)
}

// IterateTaxCaps iterates over all the stored TaxCap and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateTaxCaps(ctx sdk.Context, cb func(denom string, taxcap sdk.Int) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TaxCapKeyPref)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		denom := string(iterator.Key()[len(types.TaxCapKeyPref):])
		ip := sdk.IntProto{}
		k.cdc.MustUnmarshal(iterator.Value(), &ip)

		if cb(denom, ip.Int) {
			break
		}
	}
}

// ClearTaxCaps iterates over all stored TaxCap and deletes the key from store
func (k Keeper) ClearTaxCaps(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TaxCapKeyPref)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}

// GetTaxCap fetches a TaxCap Cap from the store stored by *denom*
func (k Keeper) GetTaxCap(ctx sdk.Context, denom string) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTaxCapKey(denom))

	// Return default cap is no tax cap has been set. Default cap assumes microtoken denom.
	if bz == nil {
		return types.DefaultCap
	}

	ip := sdk.IntProto{}
	k.cdc.MustUnmarshal(bz, &ip)
	return ip.Int
}

// SetTaxCap sets a TaxCap to the store
func (k Keeper) SetTaxCap(ctx sdk.Context, denom string, taxcap sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.IntProto{Int: taxcap})

	store.Set(types.GetTaxCapKey(denom), bz)
}

// GetTaxCaps returns all TaxCap
func (k Keeper) GetTaxCaps(ctx sdk.Context) []types.TaxCap {
	var taxCaps []types.TaxCap

	k.IterateTaxCaps(ctx, func(denom string, taxcap sdk.Int) bool {
		taxCaps = append(taxCaps, types.TaxCap{
			Denom: denom,
			Cap:   taxcap,
		})
		return false
	})

	return taxCaps
}

// RecordTaxProceeds adds collected tax to the TaxProceeds record for the current *Epoch*
func (k Keeper) RecordTaxProceeds(ctx sdk.Context, proceeds sdk.Coins) {
	if proceeds.IsZero() {
		return
	}
	taxproceeds := k.GetTaxProceeds(ctx)
	k.SetTaxProceeds(ctx, taxproceeds.Add(proceeds...))
}

// GetTaxProceeds fetches the current tax proceeds collected in the current *Epoch* before the end of said *Epoch*
func (k Keeper) GetTaxProceeds(ctx sdk.Context) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.TaxProceedsKey)
	csp := types.TaxProceeds{}
	if bz == nil {
		csp.TaxProceeds = sdk.Coins{}
	} else {
		k.cdc.MustUnmarshal(bz, &csp)
	}
	return csp.TaxProceeds
}

// SetTaxProceeds sets the tax proceeds to the store
func (k Keeper) SetTaxProceeds(ctx sdk.Context, proceeds sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.TaxProceeds{TaxProceeds: proceeds})

	store.Set(types.TaxProceedsKey, bz)
}

// GetEpochTaxProceeds fetches the tax proceeds collected a specified epoch
func (k Keeper) GetEpochTaxProceeds(ctx sdk.Context, epoch int64) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetEpochTaxProceedsKey(epoch))

	csp := types.TaxProceeds{}
	if bz == nil {
		return sdk.Coins{}
	}
	k.cdc.MustUnmarshal(bz, &csp)
	return csp.TaxProceeds
}

// SetEpochTaxProceeds sets the tax proceeds collected during a epoch to the store
func (k Keeper) SetEpochTaxProceeds(ctx sdk.Context, epoch int64, proceeds sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.TaxProceeds{TaxProceeds: proceeds})

	store.Set(types.GetEpochTaxProceedsKey(epoch), bz)
}

// GetPayouts fetches []Payout for a specified *epoch* from the store
func (k Keeper) GetPayouts(ctx sdk.Context, epoch int64) []types.Payout {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPayoutsKey(epoch))
	pos := types.Payouts{}

	if bz == nil {
		return []types.Payout{}
	}

	k.cdc.MustUnmarshal(bz, &pos)
	return pos.Payouts
}

// SetPayouts sets the []Payout to the store, stored under *epoch*.
// Payout is used for query purposes only, hence the lack of need for storing individual Payout objects.
func (k Keeper) SetPayouts(ctx sdk.Context, epoch int64, payouts []types.Payout) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.Payouts{Payouts: payouts})

	store.Set(types.GetPayoutsKey(epoch), bz)
}

// GetCollectionEpochs creates and returns a slice of all existing CollectionEpoch
func (k Keeper) GetCollectionEpochs(ctx sdk.Context) []types.CollectionEpoch {
	collectionEpochs := []types.CollectionEpoch{}

	// Iterate through existing *epoch*s and create CollectionEpoch per epoch
	for p := int64(0); p < k.GetCurrentEpoch(ctx); p++ {
		taxProceeds := k.GetEpochTaxProceeds(ctx, p)
		payouts := k.GetPayouts(ctx, p)

		// do not include CollectionEpochs that have empty tax proceeds && empty payouts
		if taxProceeds.IsZero() && (len(payouts) == 0) {
			continue
		}
		collectionEpoch := types.CollectionEpoch{
			Epoch:        uint64(p),
			TaxCollected: taxProceeds,
			Payouts:      payouts,
		}
		collectionEpochs = append(collectionEpochs, collectionEpoch)
	}

	return collectionEpochs
}

// ClearEpochTaxProceeds deletes all Tax Proceeds stored by *epoch* from the store
// NOTE: For testing purposes only
func (k Keeper) ClearEpochTaxProceeds(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.EpochTaxProceedsKeyPref)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// ClearPayouts deletes all []Payout stored by *epoch* from the store
// NOTE: For testing purposes only
func (k Keeper) ClearPayouts(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.PayoutsKeyPref)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

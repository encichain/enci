package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/charity/x/charity/types"

	// this line is used by starport scaffolding # ibc/keeper/import

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
		paramStore    paramstypes.Subspace
	}
)

func NewKeeper(
	cdc codec.Marshaler,
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

// GetTaxRate gets the tax rate
func (k Keeper) GetTaxRate(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.TaxRateKey)
	if b == nil {
		return types.DefaultTaxRate
	}
	decpro := sdk.DecProto{}
	k.cdc.MustUnmarshalBinaryBare(b, &decpro)

	return decpro.Dec
}

// SetTaxRate sets the TaxRate in the store
func (k Keeper) SetTaxRate(ctx sdk.Context, taxRate sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&sdk.DecProto{Dec: taxRate})

	// Set the store
	store.Set(types.TaxRateKey, b)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

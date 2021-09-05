package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/charity/x/charity/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute

		bankKeeper    types.BankKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	bankKeeper types.BankKeeper,
	accountKeeper types.AccountKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
	}
}

// Check if charity tax collector account is set. Panic if nil

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

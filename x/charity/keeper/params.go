package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
)

// GetAllParams returns total param set
func (k Keeper) GetAllParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return
}

// GetCharity returns all Charities of params.
func (k Keeper) GetCharity(ctx sdk.Context) (charities []types.Charity) {
	k.paramStore.Get(ctx, types.KeyCharities, &charities)
	return
}

// GetTaxRate returns the current charity tax rate
func (k Keeper) GetTaxRate(ctx sdk.Context) (taxrate sdk.Dec) {
	k.paramStore.Get(ctx, types.KeyTaxRate, &taxrate)
	return
}

// GetParamTaxCaps returns the []TaxCap from the paramstore
func (k Keeper) GetParamTaxCaps(ctx sdk.Context) (taxcaps []types.TaxCap) {
	k.paramStore.Get(ctx, types.KeyTaxCaps, &taxcaps)
	return
}

// SetParams sets all params of charity module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

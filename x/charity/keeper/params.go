package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/charity/types"
)

// GetAllParams returns total param set
func (k Keeper) GetAllParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return
}

// GetCharity returns all Charities of params.
func (k Keeper) GetCharity(ctx sdk.Context) (charities []types.Charity) {
	k.paramStore.Get(ctx, types.ParamKeyCharities, &charities)
	return
}

// GetTaxRate returns the current charity tax rate
func (k Keeper) GetTaxRate(ctx sdk.Context) (taxrate sdk.Dec) {
	k.paramStore.Get(ctx, types.ParamKeyTaxRate, &taxrate)
	return
}

// SetTaxRate sets the specified TaxRate to the param store
// Note: For testing purposes only
func (k Keeper) SetTaxRate(ctx sdk.Context, taxRate sdk.Dec) error {
	// Marshal the Dec. x/Params uses amino codec
	taxJSON, err := taxRate.MarshalJSON()
	if err != nil {
		return err
	}
	return k.paramStore.Update(ctx, types.ParamKeyTaxRate, taxJSON)
}

// GetParamTaxCaps returns the []TaxCap from the paramstore
func (k Keeper) GetParamTaxCaps(ctx sdk.Context) (taxcaps []types.TaxCap) {
	k.paramStore.Get(ctx, types.ParamKeyTaxCaps, &taxcaps)
	return
}

// GetBurnRate returns the current charity burn rate
func (k Keeper) GetBurnRate(ctx sdk.Context) (burnRate sdk.Dec) {
	k.paramStore.Get(ctx, types.ParamKeyBurnRate, &burnRate)
	return
}

// SetBurnRate sets the specified BurnRate to the param store..
// Note: For testing purposes only
func (k Keeper) SetBurnRate(ctx sdk.Context, burnRate sdk.Dec) error {
	// Marshal the Dec. x/Params uses amino codec
	rateJSON, err := burnRate.MarshalJSON()
	if err != nil {
		return nil
	}
	return k.paramStore.Update(ctx, types.ParamKeyBurnRate, rateJSON)
}

// SetParams sets all params of charity module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

// SyncTaxCaps syncs the store []Taxcap with the paramstore []Taxcap.
// To be called at end of period
func (k Keeper) SyncTaxCaps(ctx sdk.Context) {
	// Clear store Taxcaps
	k.ClearTaxCaps(ctx)
	taxcaps := k.GetParamTaxCaps(ctx)

	// Set taxcaps to store
	for _, taxcap := range taxcaps {
		k.SetTaxCap(ctx, taxcap.Denom, taxcap.Cap)
	}
}

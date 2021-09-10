package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/charity/x/charity/types"
)

// GetAllParams returns total param set
func (k Keeper) GetAllParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return params
}

// CharityOne returns CharityOne of params.
func (k Keeper) CharityOne(ctx sdk.Context) (charity types.Charity) {
	k.paramStore.Get(ctx, types.KeyCharityOne, &charity)
	return charity
}

// CharityOne returns CharityOne of params.
func (k Keeper) CharityTwo(ctx sdk.Context) (charity types.Charity) {
	k.paramStore.Get(ctx, types.KeyCharityTwo, &charity)
	return charity
}

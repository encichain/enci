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

// GetCharityOne returns CharityOne of params.
func (k Keeper) GetCharityOne(ctx sdk.Context) (charity types.Charity) {
	k.paramStore.Get(ctx, types.KeyCharityOne, &charity)
	return
}

// GetCharityOne returns CharityOne of params.
func (k Keeper) GetCharityTwo(ctx sdk.Context) (charity types.Charity) {
	k.paramStore.Get(ctx, types.KeyCharityTwo, &charity)
	return
}

// SetParams sets all params of charity module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

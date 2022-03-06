package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
)

func (k Keeper) GetVoteFrequency(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyVoteFrequency, &res)
	return
}

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return params
}

//func (k Keeper) GetPrevotePeriod(ctx sdk.Context) (res uint64) {}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

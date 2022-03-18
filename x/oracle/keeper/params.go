package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
)

// GetParams returns the total set of oracle parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramStore.GetParamSet(ctx, &params)
	return params
}

// GetVoteFrequency returns the VoteFrequency from params
func (k Keeper) GetVoteFrequency(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyVoteFrequency, &res)
	return
}

//GetVotePeriod returns the VotePeriod from params
func (k Keeper) GetVotePeriod(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyVotePeriod, &res)
	return
}

//GetVoteThreshold returns the minimum VoteThreshold from params
func (k Keeper) GetVoteThreshold(ctx sdk.Context) (res sdk.Dec) {
	k.paramStore.Get(ctx, types.KeyVoteThreshold, &res)
	return
}

// GetPrevotePeriod returns the PrevotePeriod from params
func (k Keeper) GetPrevotePeriod(ctx sdk.Context) (res uint64) {
	k.paramStore.Get(ctx, types.KeyPrevotePeriod, &res)
	return
}

// SetParams sets the total set of oracle parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramStore.SetParamSet(ctx, &params)
}

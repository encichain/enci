package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/encichain/enci/x/oracle/types"
)

// GetVoterDelegate returns a delegate address if it exists, given a *Validator* operator address
// Return error if no delegate mapped
func (k Keeper) GetVoterDelegate(ctx sdk.Context, val sdk.ValAddress) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValDelKey(val))

	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrNoVoteDelegate, val.String())
	}

	return sdk.AccAddress(bz), nil
}

// GetVoterDelegator returns a validator address if it exists, given a *delegate* address.
// Return error if no validator delegator address mapped
func (k Keeper) GetVoterDelegator(ctx sdk.Context, del sdk.AccAddress) (sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetDelValKey(del))

	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrNoVoteDelegator, del.String())
	}

	return sdk.AccAddress(bz), nil
}

// SetVoterDelegation stores a voter delegation in both directions to the store for easy lookup
func (k Keeper) SetVoterDelegation(ctx sdk.Context, del sdk.AccAddress, val sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDelValKey(del), val.Bytes())
	store.Set(types.GetValDelKey(val), del.Bytes())
}

// IterateVoterDelegations iterates through all stored voter delegations and performs callback function
// Stops iteration when no more valid values. Uses only *Validator* address key to reduce redundancy
func (k Keeper) IterateVoterDelegations(ctx sdk.Context, cb func(val sdk.ValAddress, del sdk.AccAddress) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValDelKey)

	for ; iterator.Valid(); iterator.Next() {
		validator := sdk.ValAddress(iterator.Key()[2:])
		delegate := sdk.AccAddress(iterator.Value())

		if cb(validator, delegate) {
			break
		}
	}
}

// GetAllVoterDelegations returns a slice of all stored VoterDelegation
func (k Keeper) GetAllVoterDelegations(ctx sdk.Context) []types.VoterDelegation {
	var delegations []types.VoterDelegation
	k.IterateVoterDelegations(ctx, func(val sdk.ValAddress, del sdk.AccAddress) bool {
		delegation := types.VoterDelegation{
			DelegateAddress:  del.String(),
			ValidatorAddress: val.String(),
		}
		delegations = append(delegations, delegation)

		return false
	})
	return delegations
}

// ClearInactiveDelegations deletes all Voter delegations with Unbonded or unknown Validator status from the store
func (k Keeper) ClearInactiveDelegations(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	k.IterateVoterDelegations(ctx, func(val sdk.ValAddress, del sdk.AccAddress) bool {
		validator, found := k.StakingKeeper.GetValidator(ctx, val)
		if !found {
			store.Delete(types.GetValDelKey(val))
			store.Delete(types.GetDelValKey(del))
		}

		if validator.Status < stakingtypes.BondStatus(2) {
			store.Delete(types.GetValDelKey(val))
			store.Delete(types.GetDelValKey(del))
		}
		return false
	})
}

package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/encichain/enci/x/oracle/types"
)

type (
	// Keeper of the oracle store
	Keeper struct {
		cdc           codec.Codec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		StakingKeeper types.StakingKeeper
		paramStore    paramstypes.Subspace
	}
)

// NewKeeper instatiates the oracle keeper
func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey, stakingKeeper types.StakingKeeper, paramStore paramstypes.Subspace) *Keeper {

	// set KeyTable if it has not already been set
	if !paramStore.HasKeyTable() {
		paramStore = paramStore.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		StakingKeeper: stakingKeeper,
		paramStore:    paramStore,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetVoteRound returns a VoteRound stored by *claimType*
func (k Keeper) GetVoteRound(ctx sdk.Context, claimType string) types.VoteRound {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVoteRoundKey(claimType))
	res := types.VoteRound{}

	if bz == nil {
		return res
	}

	k.cdc.MustUnmarshal(bz, &res)

	return res
}

// SetVoteRound sets a VoteRound to the store - stored by *claimType*
func (k Keeper) SetVoteRound(ctx sdk.Context, voteRound types.VoteRound) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.VoteRound{
		ClaimType: voteRound.ClaimType,
		Votes:     voteRound.Votes,
	})
	store.Set(types.GetVoteRoundKey(voteRound.ClaimType), bz)
}

// IterateVoteRounds iterates through all stored VoteRounds and performs callback function
// Stops iteration when no more valid
func (k Keeper) IterateVoteRounds(ctx sdk.Context, cb func(voteRound types.VoteRound) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.VoteRoundKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		voteRound := types.VoteRound{}
		k.cdc.MustUnmarshal(iterator.Value(), &voteRound)

		if cb(voteRound) {
			break
		}
	}
}

// GetAllVoteRounds returns all VoteRounds for all claim types
func (k Keeper) GetAllVoteRounds(ctx sdk.Context) []types.VoteRound {
	var voteRounds []types.VoteRound

	k.IterateVoteRounds(ctx, func(voteRound types.VoteRound) bool {
		voteRounds = append(voteRounds, types.VoteRound{
			ClaimType: voteRound.ClaimType,
			Votes:     voteRound.Votes,
		})
		return false
	})
	return voteRounds
}

// IteratePrevoteRounds iterates through all stored PrevoteRounds and performs callback function
// Stops iteration when no more valid
func (k Keeper) IteratePrevoteRounds(ctx sdk.Context, cb func(prevoteRound types.PrevoteRound) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrevoteRoundKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		//claimType := string(iterator.Key()[len(types.PrevoteRoundKey):])
		prevoteRound := types.PrevoteRound{}
		k.cdc.MustUnmarshal(iterator.Value(), &prevoteRound)

		if cb(prevoteRound) {
			break
		}
	}
}

// GetAllPrevoteRounds returns all PrevoteRounds
func (k Keeper) GetAllPrevoteRounds(ctx sdk.Context) []types.PrevoteRound {
	var prevoteRounds []types.PrevoteRound

	k.IteratePrevoteRounds(ctx, func(prevoteRound types.PrevoteRound) bool {
		prevoteRounds = append(prevoteRounds, types.PrevoteRound{
			ClaimType: prevoteRound.ClaimType,
			Prevotes:  prevoteRound.Prevotes,
		})
		return false
	})
	return prevoteRounds
}

// GetPrevoteRound returns a PrevoteRound stored by *claimType*
func (k Keeper) GetPrevoteRound(ctx sdk.Context, claimType string) types.PrevoteRound {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPrevoteRoundKey(claimType))
	res := types.PrevoteRound{}

	if bz == nil {
		return res
	}
	k.cdc.MustUnmarshal(bz, &res)

	return res
}

// SetPrevoteRound sets a Prevote round to the store by *claimType*
func (k Keeper) SetPrevoteRound(ctx sdk.Context, prevoteRound types.PrevoteRound) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.PrevoteRound{
		ClaimType: prevoteRound.ClaimType,
		Prevotes:  prevoteRound.Prevotes,
	})
	store.Set(types.GetPrevoteRoundKey(prevoteRound.ClaimType), bz)
}

// GetVote returns a Vote from the store, by *Validator* address
func (k Keeper) GetVote(ctx sdk.Context, val sdk.ValAddress) types.Vote {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVoteKey(val))
	vote := types.Vote{}

	if bz == nil {
		return vote
	}
	k.cdc.MustUnmarshal(bz, &vote)

	return vote
}

// SetVote sets a Vote to the store, by *Validator* adderss
func (k Keeper) SetVote(ctx sdk.Context, val sdk.ValAddress, vote types.Vote) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.Vote{
		Claim:     vote.Claim,
		Validator: vote.Validator,
	})
	store.Set(types.GetVoteKey(val), bz)
}

// IterateVotes iterates through all stored Vote and performs callback function
// Stops iteration when no more valid
func (k Keeper) IterateVotes(ctx sdk.Context, cb func(vote types.Vote) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.VoteKey)

	for ; iterator.Valid(); iterator.Next() {
		vote := types.Vote{}
		k.cdc.MustUnmarshal(iterator.Value(), &vote)

		if cb(vote) {
			break
		}
	}
}

// GetAllVotes returns a slice of all stored votes
func (k Keeper) GetAllVotes(ctx sdk.Context) []types.Vote {
	votes := []types.Vote{}

	k.IterateVotes(ctx, func(vote types.Vote) bool {
		votes = append(votes, types.Vote{
			Claim:     vote.Claim,
			Validator: vote.Validator,
		})
		return false
	})
	return votes
}

// GetPrevote returns a Prevote from the store, by *Validator* address
func (k Keeper) GetPrevote(ctx sdk.Context, val sdk.ValAddress) types.Prevote {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetPrevoteKey(val))
	prevote := types.Prevote{}

	if bz == nil {
		return prevote
	}
	k.cdc.MustUnmarshal(bz, &prevote)

	return prevote
}

// SetPrevote sets a Prevote to the store, by *Validator* address
func (k Keeper) SetPrevote(ctx sdk.Context, val sdk.ValAddress, prevote types.Prevote) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.Prevote{
		Hash:        prevote.Hash,
		Validator:   prevote.Validator,
		SubmitBlock: prevote.SubmitBlock,
	})
	store.Set(types.GetPrevoteKey(val), bz)
}

// IteratePrevotes iterates through all stored prevote and performs callback function
// Stops iteration when no more valid
func (k Keeper) IteratePrevotes(ctx sdk.Context, cb func(prevote types.Prevote) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PrevoteKey)

	for ; iterator.Valid(); iterator.Next() {
		prevote := types.Prevote{}
		k.cdc.MustUnmarshal(iterator.Value(), &prevote)

		if cb(prevote) {
			break
		}
	}
}

// GetAllPrevotes returns a slice of all stored prevotes
func (k Keeper) GetAllPrevotes(ctx sdk.Context) []types.Prevote {
	prevotes := []types.Prevote{}
	k.IteratePrevotes(ctx, func(prevote types.Prevote) bool {
		prevotes = append(prevotes, types.Prevote{
			Hash:        prevote.Hash,
			Validator:   prevote.Validator,
			SubmitBlock: prevote.SubmitBlock,
		})
		return false
	})
	return prevotes
}

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

// ClearInactiveDelegations deletes all Voter delegations with Validator Unbonded or unknown status from the store
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

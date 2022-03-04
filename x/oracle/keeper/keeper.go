package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
		ClaimType:      voteRound.ClaimType,
		Votes:          voteRound.Votes,
		AggregatePower: voteRound.AggregatePower,
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
			ClaimType:      voteRound.ClaimType,
			Votes:          voteRound.Votes,
			AggregatePower: voteRound.AggregatePower,
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

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
)

// TODO: Check for existing claim
//AddVoteToRound appends a Vote to a VoteRound
func (k Keeper) AppendVoteToRound(ctx sdk.Context, vote types.Vote, claimType string) {
	round := k.GetVoteRound(ctx, claimType)
	round.Votes = append(round.Votes, vote)
	round.AggregatePower += vote.VotePower

	k.SetVoteRound(ctx, round)
}

// GetVoteRound returns a VoteRound stored by *claimType*
func (k Keeper) GetVoteRound(ctx sdk.Context, claimType string) types.VoteRound {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVoteRoundKey(claimType))
	round := types.VoteRound{}

	if bz == nil {
		return round
	}

	k.cdc.MustUnmarshal(bz, &round)

	return round
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

// ClearVoteRounds deletes all VoteRounds from the key store
func (k Keeper) ClearVoteRounds(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.VoteRoundKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// GetAllVoteRounds returns all VoteRounds for all claim types
func (k Keeper) GetAllVoteRounds(ctx sdk.Context) []types.VoteRound {
	voteRounds := []types.VoteRound{}

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
	prevoteRounds := []types.PrevoteRound{}

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

// ClearPrevoteRounds deletes all Prevotes from the key store
func (k Keeper) ClearPrevoteRounds(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PrevoteRoundKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

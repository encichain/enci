package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// IsVotePeriod checks if current block is part of a VotePeriod...
// 0 modulus values are included in the check
// Ex: let VoteFrequency = 100, VotePeriod = 3, and PrevotePeriod = 3
// excluding genesis, first Prevote period will begin at block height 99 (calculated as 100) and end at 101
// for a total of three blocks. VotePeriod would begin at 102 and end at 104
func (k Keeper) IsVotePeriod(ctx sdk.Context) bool {
	params := k.GetParams(ctx)

	if i := uint64(ctx.BlockHeight()+1) % params.VoteFrequency; i < params.VotePeriod+params.PrevotePeriod {
		return i >= params.PrevotePeriod
	}
	return false
}

// IsPrevotePeriod check if current block is part of a PrevotePeriod
func (k Keeper) IsPrevotePeriod(ctx sdk.Context) bool {
	params := k.GetParams(ctx)

	if i := uint64(ctx.BlockHeight()+1) % params.VoteFrequency; i < params.PrevotePeriod {
		return true
	}
	return false
}

// IsVotePeriodEnd checks if it is the last block of a VotePeriod
func (k Keeper) IsVotePeriodEnd(ctx sdk.Context) bool {
	params := k.GetParams(ctx)
	if i := uint64(ctx.BlockHeight()+1) % params.VoteFrequency; i%(params.PrevotePeriod+params.VotePeriod-1) == 0 {
		return true
	}
	return false
}

// PreviousVotePeriod returns the height of the start of the previous prevotePeriod
func (k Keeper) PreviousPrevotePeriod(ctx sdk.Context) uint64 {
	params := k.GetParams(ctx)
	return (uint64(ctx.BlockHeight()+1) / params.VoteFrequency) * params.VoteFrequency
}

// GetVote returns a Vote from the store, by *claim type* | *Validator* address
func (k Keeper) GetVote(ctx sdk.Context, val sdk.ValAddress, claimType string) types.Vote {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetVoteKey(val, claimType))
	vote := types.Vote{}

	if bz == nil {
		return vote
	}
	k.cdc.MustUnmarshal(bz, &vote)

	return vote
}

// SetVote sets a Vote to the store, by *claim type* | *Validator* address
func (k Keeper) SetVote(ctx sdk.Context, val sdk.ValAddress, vote types.Vote, claimType string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.Vote{
		Claim:     vote.Claim,
		Validator: vote.Validator,
		VotePower: vote.VotePower,
	})
	store.Set(types.GetVoteKey(val, claimType), bz)
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

// GetVotesByClaim returns a slice of all stored votes for a specific claim type
func (k Keeper) GetVotesByClaimType(ctx sdk.Context, claimType string) (votes []types.Vote) {
	store := ctx.KVStore(k.storeKey)
	voteClaimKey := append(types.VoteKey, types.ClaimLengthPrefix([]byte(claimType))...)
	iter := sdk.KVStorePrefixIterator(store, voteClaimKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		vote := types.Vote{}
		k.cdc.MustUnmarshal(iter.Value(), &vote)
		votes = append(votes, vote)
	}
	return
}

// DeleteVote deletes a vote for a specified claim type from a validator
func (k Keeper) DeleteVote(ctx sdk.Context, val sdk.ValAddress, claimType string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetVoteKey(val, claimType))
}

// DeleteAllVotes deletes all votes from the store for all claim types
func (k Keeper) DeleteAllVotes(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.VoteKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// GetPrevote returns a Prevote from the store, by *claim type* | *Validator* address
func (k Keeper) GetPrevote(ctx sdk.Context, val sdk.ValAddress, claimType string) (types.Prevote, error) {
	store := ctx.KVStore(k.storeKey)
	prevote := types.Prevote{}
	bz := store.Get(types.GetPrevoteKey(val, claimType))
	if bz == nil {
		return prevote, sdkerrors.Wrap(types.ErrNoPrevote, val.String())
	}
	k.cdc.MustUnmarshal(bz, &prevote)

	return prevote, nil
}

// SetPrevote sets a Prevote to the store, by *claim type* | *Validator* address
func (k Keeper) SetPrevote(ctx sdk.Context, prevote types.Prevote, claimType string) error {
	store := ctx.KVStore(k.storeKey)
	valAddr, err := sdk.ValAddressFromBech32(prevote.Validator)
	if err != nil {
		return err
	}
	bz := k.cdc.MustMarshal(&types.Prevote{
		Hash:        prevote.Hash,
		Validator:   prevote.Validator,
		SubmitBlock: prevote.SubmitBlock,
	})
	store.Set(types.GetPrevoteKey(valAddr, claimType), bz)
	return nil
}

// DeletePrevote deletes a prevote for a specified claim type from a validator
func (k Keeper) DeletePrevote(ctx sdk.Context, val sdk.ValAddress, claimType string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetPrevoteKey(val, claimType))
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

// GetPrevotesByClaimType returns a slice of all stored prevotes for a claim type
func (k Keeper) GetPrevotesByClaimType(ctx sdk.Context, claimType string) (prevotes []types.Prevote) {
	store := ctx.KVStore(k.storeKey)
	prevoteClaimKey := append(types.PrevoteKey, types.ClaimLengthPrefix([]byte(claimType))...)
	iter := sdk.KVStorePrefixIterator(store, prevoteClaimKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		prevote := types.Prevote{}
		k.cdc.MustUnmarshal(iter.Value(), &prevote)
		prevotes = append(prevotes, prevote)
	}
	return
}

// DeleteAllPrevotes deletes all prevotes from the store for all claim types
func (k Keeper) DeleteAllPrevotes(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.PrevoteKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// RegisterClaimType registers claim type to the store
func (k Keeper) RegisterClaimType(ctx sdk.Context, claimType string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.ClaimType{ClaimType: claimType})
	store.Set(types.GetClaimTypeKey(claimType), bz)
}

// GetAllClaimTypes returns a slice of all registered claim types in string form
func (k Keeper) GetAllClaimTypes(ctx sdk.Context) (claimTypes []string) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ClaimTypeKey)

	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		claimType := types.ClaimType{}
		k.cdc.MustUnmarshal(iter.Value(), &claimType)
		claimTypes = append(claimTypes, claimType.ClaimType)
	}
	return
}

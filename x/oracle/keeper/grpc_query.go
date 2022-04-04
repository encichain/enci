package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// NewQuerier returns a x/oracle implementation of a QueryServer
func NewQuerier(keeper Keeper) types.QueryServer {
	return &Querier{Keeper: keeper}
}

// Params returns the x/oracle params set
func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// VoteRounds returns a slice of VoteRound for every claim type, containing all current votes
func (q Querier) VoteRounds(c context.Context, req *types.QueryVoteRoundsRequest) (*types.QueryVoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	claimTypes := q.GetAllClaimTypes(ctx)
	voteRounds := []types.VoteRound{}
	for _, claimType := range claimTypes {
		votes := q.GetVotesByClaimType(ctx, claimType)
		if len(votes) == 0 {
			continue
		}
		voteRounds = append(voteRounds, types.NewVoteRound(claimType, votes))
	}
	return &types.QueryVoteRoundsResponse{VoteRounds: voteRounds}, nil
}

// PrevoteRounds returns a PrevoteRound for every claim type, containing all current prevotes
func (q Querier) PrevoteRounds(c context.Context, req *types.QueryPrevoteRoundsRequest) (*types.QueryPrevoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	claimTypes := q.GetAllClaimTypes(ctx)
	prevoteRounds := []types.PrevoteRound{}
	for _, claimType := range claimTypes {
		prevotes := q.GetPrevotesByClaimType(ctx, claimType)
		if len(prevotes) == 0 {
			continue
		}
		prevoteRounds = append(prevoteRounds, types.NewPrevoteRound(claimType, prevotes))
	}
	return &types.QueryPrevoteRoundsResponse{PrevoteRounds: prevoteRounds}, nil
}

// VoterDelegations returns a slice of all active VoterDelegation
func (q Querier) VoterDelegations(c context.Context, req *types.QueryVoterDelegationsRequest) (*types.QueryVoterDelegationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	delegations := q.GetAllVoterDelegations(ctx)
	return &types.QueryVoterDelegationsResponse{VoterDelegations: delegations}, nil
}

// QueryDelegateAddress returns a delegate address connected to validator address if delegation exists
func (q Querier) DelegateAddress(c context.Context, req *types.QueryDelegateAddressRequest) (*types.QueryDelegateAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, err := sdk.ValAddressFromBech32(req.Validator)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	delegateAddr, err := q.GetVoterDelegate(ctx, val)
	if delegateAddr == nil || err != nil {
		return nil, status.Errorf(
			codes.NotFound, "there is no delegator for address %s", req.Validator,
		)
	}

	return &types.QueryDelegateAddressResponse{
		Delegate: delegateAddr.String(),
	}, nil
}

// DelegatorAddress returns a validator address connected to a delegate address, if delegation exists
func (q Querier) DelegatorAddress(c context.Context, req *types.QueryDelegatorAddressRequest) (*types.QueryDelegatorAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	del, err := sdk.AccAddressFromBech32(req.Delegate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	validatorAddr, err := q.GetVoterDelegator(ctx, del)
	if validatorAddr == nil || err != nil {
		return nil, status.Errorf(
			codes.NotFound, "delegator address for delegate %s", req.Delegate,
		)
	}

	return &types.QueryDelegatorAddressResponse{
		Validator: validatorAddr.String(),
	}, nil
}

// NextVotePeriod returns the block height of the beginning of the next VotePeriod
// Note that genesis vote period is skipped
func (q Querier) NextVotePeriod(c context.Context, req *types.QueryNextVotePeriodRequest) (*types.QueryNextVotePeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	height := uint64(ctx.BlockHeight())
	params := q.GetParams(ctx)

	nextPeriod := (height/params.VoteFrequency+1)*params.VoteFrequency + params.PrevotePeriod - 1

	return &types.QueryNextVotePeriodResponse{Block: nextPeriod}, nil
}

// NextPrevote returns the block height of the beginning of the next PrevotePeriod
// Note that genesis prevote period is skipped
func (q Querier) NextPrevote(c context.Context, req *types.QueryNextPrevoteRequest) (*types.QueryNextPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	height := uint64(ctx.BlockHeight())
	params := q.GetParams(ctx)

	nextPeriod := ((height/params.VoteFrequency)+1)*params.VoteFrequency - 1

	return &types.QueryNextPrevoteResponse{Block: nextPeriod}, nil
}

// ClaimTypes returns all registered claim types
func (q Querier) ClaimTypes(c context.Context, req *types.QueryClaimTypesRequest) (*types.QueryClaimTypesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	claimTypes := q.GetAllClaimTypes(ctx)
	return &types.QueryClaimTypesResponse{ClaimTypes: claimTypes}, nil
}

// IsVotePeriod returns if current block is part of a vote period
func (q Querier) IsVotePeriod(c context.Context, req *types.QueryIsVotePeriodRequest) (*types.QueryIsVotePeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryIsVotePeriodResponse{IsVotePeriod: q.Keeper.IsVotePeriod(ctx)}, nil
}

// IsPrevotePeriod returns if current block is part of a prevote period
func (q Querier) IsPrevotePeriod(c context.Context, req *types.QueryIsPrevotePeriodRequest) (*types.QueryIsPrevotePeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryIsPrevotePeriodResponse{IsPrevotePeriod: q.Keeper.IsPrevotePeriod(ctx)}, nil
}

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

// VoteRounds returns a slice of VoteRound for every claim type
func (q Querier) VoteRounds(c context.Context, req *types.QueryVoteRoundsRequest) (*types.QueryVoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	rounds := q.GetAllVoteRounds(ctx)
	return &types.QueryVoteRoundsResponse{VoteRounds: rounds}, nil
}

// PrevoteRounds returns a slice of PrevoteRound for every claim type
func (q Querier) PrevoteRounds(c context.Context, req *types.QueryPrevoteRoundsRequest) (*types.QueryPrevoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	rounds := q.GetAllPrevoteRounds(ctx)
	return &types.QueryPrevoteRoundsResponse{PrevoteRounds: rounds}, nil
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
func (q Querier) NextVotePeriod(c context.Context, req *types.QueryNextVotePeriodRequest) (*types.QueryNextVotePeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	height := uint64(ctx.BlockHeight())
	params := q.GetParams(ctx)

	nextPeriod := (height/params.VoteFrequency+1)*params.VoteFrequency + params.PrevotePeriod - 1

	return &types.QueryNextVotePeriodResponse{Block: nextPeriod}, nil
}

// NextPrevote returns the block height of the beginning of the next PrevotePeriod
func (q Querier) NextPrevote(c context.Context, req *types.QueryNextPrevoteRequest) (*types.QueryNextPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	height := uint64(ctx.BlockHeight())
	params := q.GetParams(ctx)

	nextPeriod := (height/params.VoteFrequency+1)*params.VoteFrequency - 1

	return &types.QueryNextPrevoteResponse{Block: nextPeriod}, nil
}

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

// Params implements the Query/Params gRPC method
func (q Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// VoteRounds implements the Query VoteRounds gRPC method
func (q Querier) VoteRounds(c context.Context, req *types.QueryVoteRoundsRequest) (*types.QueryVoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	rounds := q.GetAllVoteRounds(ctx)
	return &types.QueryVoteRoundsResponse{VoteRounds: rounds}, nil
}

// PrevoteRounds implements the Query PrevoteRounds gRPC method
func (q Querier) PrevoteRounds(c context.Context, req *types.QueryPrevoteRoundsRequest) (*types.QueryPrevoteRoundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	rounds := q.GetAllPrevoteRounds(ctx)
	return &types.QueryPrevoteRoundsResponse{PrevoteRounds: rounds}, nil
}

// VoterDelegations implements the Query Voter Delegations gRPC method
func (q Querier) VoterDelegations(c context.Context, req *types.QueryVoterDelegationsRequest) (*types.QueryVoterDelegationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	delegations := q.GetAllVoterDelegations(ctx)
	return &types.QueryVoterDelegationsResponse{VoterDelegations: delegations}, nil
}

// QueryDelegateAddress implements Query Delegate Address gRPC method
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

// DelegatorAddress implements Query Validator Address gRPC method
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

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/charity/x/charity/types"
)

// Querier used as alias to Keeper to avoid duplicate methods.
type Quierer struct {
	Keeper
}

//NewQuerier returns QueryServer for the Keeper
func NewQuerier(k Keeper) types.QueryServer {
	return &Quierer{Keeper: k}
}

var _ types.QueryServer = Quierer{}

// TaxRate returns the set tax rate
func (q Quierer) TaxRate(context context.Context, req *types.QueryTaxRateRequest) (*types.QueryTaxRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxRateResponse{TaxRate: q.GetTaxRate(ctx)}, nil
}

// AllParams returns all params from param store
func (q Quierer) Params(context context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryParamsResponse{Params: q.GetAllParams(ctx)}, nil
}

// CharityOne returns the set charity one
func (q Quierer) CharityOne(context context.Context, req *types.QueryCharityOneRequest) (*types.QueryCharityOneResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryCharityOneResponse{CharityOne: q.GetCharityOne(ctx)}, nil

}

// CharityOne returns the set charity two
func (q Quierer) CharityTwo(context context.Context, req *types.QueryCharityTwoRequest) (*types.QueryCharityTwoResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryCharityTwoResponse{CharityTwo: q.GetCharityTwo(ctx)}, nil
}

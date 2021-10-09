package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
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

// Charities returns the set charities or []Charity{} if empty
func (q Quierer) Charities(context context.Context, req *types.QueryCharityOneRequest) (*types.QueryCharityOneResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryCharityResponse{CharityOne: q.GetCharityOne(ctx)}, nil

}

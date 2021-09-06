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

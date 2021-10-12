package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
)

// Querier used as alias to Keeper to avoid duplicate methods.
type Querier struct {
	Keeper
}

//NewQuerier returns QueryServer for the Keeper
func NewQuerier(k Keeper) types.QueryServer {
	return &Querier{Keeper: k}
}

var _ types.QueryServer = Querier{}

// TaxRate returns the set tax rate
func (q Querier) TaxRate(context context.Context, req *types.QueryTaxRateRequest) (*types.QueryTaxRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxRateResponse{TaxRate: q.GetTaxRate(ctx)}, nil
}

// AllParams returns all params from param store
func (q Querier) Params(context context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryParamsResponse{Params: q.GetAllParams(ctx)}, nil
}

// Charities returns the set charities or []Charity{} if empty
func (q Querier) Charities(context context.Context, req *types.QueryCharitiesRequest) (*types.QueryCharitiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryCharitiesResponse{Charity: q.GetCharity(ctx)}, nil

}

// TaxCap returns a tax cap based on *denom*
func (q Querier) TaxCap(context context.Context, req *types.QueryTaxCapRequest) (*types.QueryTaxCapResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}
	err := sdk.ValidateDenom(req.Denom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "denom is not valid")
	}

	return &types.QueryTaxCapResponse{Cap: q.GetTaxCap(ctx, req.Denom)}, nil
}

// TaxCaps returns all tax caps
func (q Querier) TaxCaps(context context.Context, req *types.QueryTaxCapsRequest) (*types.QueryTaxCapsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	var taxcaps []types.TaxCap

	// Iterate tax caps and append each to var taxcaps
	q.IterateTaxCaps(ctx, func(denom string, taxcap sdk.Int) bool {
		taxcaps = append(taxcaps, types.TaxCap{
			Denom: denom,
			Cap:   taxcap,
		})
		return false
	})
	return &types.QueryTaxCapsResponse{TaxCaps: taxcaps}, nil
}

// TaxRateLimits returns the limits to the taxrate
func (q Querier) TaxRateLimits(context context.Context, req *types.QueryTaxRateLimitsRequest) (*types.QueryTaxRateLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxRateLimitsResponse{TaxRateLimits: q.GetTaxRateLimits(ctx)}, nil
}

// TaxProceeds returns the current tax proceeds collected for the current period
func (q Querier) TaxProceeds(context context.Context, req *types.QueryTaxProceedsRequest) (*types.QueryTaxProceedsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxProceedsResponse{TaxProceeds: q.GetTaxProceeds(ctx)}, nil
}

// CollectionPeriods returns all CollectionPeriod
func (q Querier) CollectionPeriods(context context.Context, req *types.QueryAllCollectionPeriodsRequest) (*types.QueryAllCollectionPeriodsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	var collectionPeriods []types.CollectionPeriod

	// Iterate through existing *period*s and create CollectionPeriod per period
	for p := int64(0); p < q.GetCurrentPeriod(ctx); p++ {
		collectionPeriod := types.CollectionPeriod{
			Period:       uint64(p),
			TaxCollected: q.GetPeriodTaxProceeds(ctx, p),
			Payouts:      q.GetPayouts(ctx, p),
		}
		collectionPeriods = append(collectionPeriods, collectionPeriod)
	}

	return &types.QueryAllCollectionPeriodsResponse{CollectionPeriods: collectionPeriods}, nil
}

// CollectionPeriod returns a CollectionPeriod based on *period*
func (q Querier) CollectionPeriod(context context.Context, req *types.QueryCollectionPeriodRequest) (*types.QueryCollectionPeriodResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.Period >= uint64(q.GetCurrentPeriod(ctx)) {
		return nil, status.Error(codes.InvalidArgument, "period must be valid")
	}

	collectionPeriod := types.CollectionPeriod{
		Period:       req.Period,
		TaxCollected: q.GetPeriodTaxProceeds(ctx, int64(req.Period)),
		Payouts:      q.GetPayouts(ctx, int64(req.Period)),
	}

	return &types.QueryCollectionPeriodResponse{CollectionPeriod: collectionPeriod}, nil
}

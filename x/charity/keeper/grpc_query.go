package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/charity/types"
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

// BurnRate returns the charity burn rate
func (q Querier) BurnRate(context context.Context, req *types.QueryBurnRateRequest) (*types.QueryBurnRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryBurnRateResponse{BurnRate: q.GetBurnRate(ctx)}, nil
}

// TaxRateLimits returns the limits to the taxrate
func (q Querier) TaxRateLimits(context context.Context, req *types.QueryTaxRateLimitsRequest) (*types.QueryTaxRateLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxRateLimitsResponse{TaxRateLimits: q.GetTaxRateLimits(ctx)}, nil
}

// TaxProceeds returns the current tax proceeds collected for the current epoch
func (q Querier) TaxProceeds(context context.Context, req *types.QueryTaxProceedsRequest) (*types.QueryTaxProceedsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	return &types.QueryTaxProceedsResponse{TaxProceeds: q.GetTaxProceeds(ctx)}, nil
}

// CollectionEpochs returns all CollectionEpoch excluding empty CollectionEpochs
func (q Querier) CollectionEpochs(context context.Context, req *types.QueryAllCollectionEpochsRequest) (*types.QueryAllCollectionEpochsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	collectionEpochs := q.GetCollectionEpochs(ctx)

	return &types.QueryAllCollectionEpochsResponse{CollectionEpochs: collectionEpochs}, nil
}

// CollectionEpoch returns a CollectionEpoch based on *epoch*
func (q Querier) CollectionEpoch(context context.Context, req *types.QueryCollectionEpochRequest) (*types.QueryCollectionEpochResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	if req.Epoch >= uint64(q.GetCurrentEpoch(ctx)) {
		return nil, status.Error(codes.InvalidArgument, "epoch must be valid")
	}

	collectionEpoch := types.CollectionEpoch{
		Epoch:        req.Epoch,
		TaxCollected: q.GetEpochTaxProceeds(ctx, int64(req.Epoch)),
		Payouts:      q.GetPayouts(ctx, int64(req.Epoch)),
	}

	return &types.QueryCollectionEpochResponse{CollectionEpoch: collectionEpoch}, nil
}

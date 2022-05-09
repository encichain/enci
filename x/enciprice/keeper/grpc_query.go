package keeper

import (
	"context"

	"github.com/encichain/enci/x/enciprice/types"
)

var _ types.QueryServer = Querier{}

// Use Querier alias to prevent duplicate Keeper method
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) types.QueryServer {
	return &Querier{Keeper: k}
}

// Params returns the current params of the x/enciprice module
// placeholder
func (q Querier) Params(context context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	// ctx := sdk.UnwrapSDKContext(context)
	return nil, nil
}

// EnciUsd returns current ENCI USD exchange rate stored on the blockchain
// placeholder
func (q Querier) EnciUsd(context context.Context, req *types.QueryEnciUsdRequest) (*types.QueryEnciUsdResponse, error) {
	//ctx := sdk.UnwrapSDKContext(context)
	return nil, nil
}

// MissCounter returns the miss counter for a specific validator address
// placeholder
func (q Querier) MissCounter(context context.Context, req *types.QueryMissCounterRequest) (*types.QueryMissCounterResponse, error) {
	//ctx := sdk.UnwrapSDKContext(context)
	return nil, nil
}

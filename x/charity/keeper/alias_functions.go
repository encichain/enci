package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k Keeper) GetDenomMetaData(ctx sdk.Context, denom string) banktypes.Metadata {
	return k.bankKeeper.GetDenomMetaData(ctx, denom)
}

func (k Keeper) SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata) {
	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)
}

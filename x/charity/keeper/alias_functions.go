package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/user/encichain/x/charity/types"
)

func (k Keeper) GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool) {
	return k.bankKeeper.GetDenomMetaData(ctx, denom)
}

func (k Keeper) SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata) {
	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)
}

func (k Keeper) GetCharityCollectorAcc(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.CharityCollectorName)
}

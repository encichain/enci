package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/user/encichain/x/charity/types"
)

func (k Keeper) GetCharityCollectorAcc(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.AccountKeeper.GetModuleAccount(ctx, types.CharityCollectorName)
}

func (k Keeper) GetBurnAcc(ctx sdk.Context) authtypes.AccountI {
	return k.AccountKeeper.GetModuleAccount(ctx, types.BurnAccName)
}

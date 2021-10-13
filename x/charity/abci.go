package charity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/user/encichain/x/charity/keeper"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	return
}

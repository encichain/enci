package charity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	charitytypes "github.com/user/charity/types"
	"github.com/user/charity/x/charity/keeper"
	"github.com/user/charity/x/charity/types"

	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init

	// TODO: Cosmos SDK v0.43.0 introduces a bool return value to the GetDenomMetaData function. Refactor to check for non-existent denom meta data.
	// Set the Denom meta data
	empty := banktypes.Metadata{}
	meta := k.GetDenomMetaData(ctx, charitytypes.MicroTokenDenom)

	if meta.Base == empty.Base {
		k.SetDenomMetaData(ctx, charitytypes.TokenMetaData)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}

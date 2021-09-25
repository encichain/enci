package charity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	charitymaintypes "github.com/user/charity/types"
	"github.com/user/charity/x/charity/keeper"
	"github.com/user/charity/x/charity/types"
	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init

	// TODO: Cosmos SDK v0.43.0 introduces a bool return value to the GetDenomMetaData function. Refactor to check for non-existent denom meta data.
	// Set the Denom meta data
	_, set := k.GetDenomMetaData(ctx, charitymaintypes.MicroTokenDenom)

	if !set {
		k.SetDenomMetaData(ctx, charitymaintypes.TokenMetaData)
	}
	k.SetTaxRate(ctx, genState.TaxRate)
	k.SetParams(ctx, genState.Params)

	//Ensure charity collector module account is set
	if k.GetCharityCollectorAcc(ctx) == nil {
		panic(fmt.Sprintf("Module account not set: %s", types.CharityCollectorName))
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	taxrate := k.GetTaxRate(ctx)
	params := k.GetAllParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return types.NewGenesis(taxrate, params)
}

package charity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity/keeper"
	"github.com/encichain/enci/x/charity/types"
	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init

	// Set the Denom meta data
	_, set := k.BankKeeper.GetDenomMetaData(ctx, coretypes.MicroTokenDenom)
	if !set {
		k.BankKeeper.SetDenomMetaData(ctx, coretypes.TokenMetaData)
	}

	k.SetParams(ctx, genState.Params)
	k.SetTaxRateLimits(ctx, genState.TaxRateLimits)

	// Set tax caps
	for _, cap := range genState.TaxCaps {
		k.SetTaxCap(ctx, cap.Denom, cap.Cap)
	}

	// Set current epoch tax proceeds
	k.SetTaxProceeds(ctx, genState.TaxProceeds)

	// Set CollectionEpoch data
	for _, epoch := range genState.CollectionEpochs {
		k.SetEpochTaxProceeds(ctx, int64(epoch.Epoch), epoch.TaxCollected)
		k.SetPayouts(ctx, int64(epoch.Epoch), epoch.Payouts)
	}

	//Ensure charity collector module account is set
	cAcc := k.GetCharityCollectorAcc(ctx)
	if cAcc == nil {
		panic(fmt.Sprintf("Module account not set: %s", types.CharityCollectorName))
	}

	bAcc := k.GetBurnAcc(ctx)
	if bAcc == nil {
		panic(fmt.Sprintf("Module account not set: %s", types.BurnAccName))
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	taxRateLimits := k.GetTaxRateLimits(ctx)
	params := k.GetAllParams(ctx)
	taxCaps := k.GetTaxCaps(ctx)
	taxProceeds := k.GetTaxProceeds(ctx)
	collectionEpochs := k.GetCollectionEpochs(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return types.NewGenesisState(params, taxRateLimits, taxCaps, taxProceeds, collectionEpochs)
}

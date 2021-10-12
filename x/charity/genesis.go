package charity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	charitymaintypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"
	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init

	// Set the Denom meta data
	_, set := k.GetDenomMetaData(ctx, charitymaintypes.MicroTokenDenom)
	if !set {
		k.SetDenomMetaData(ctx, charitymaintypes.TokenMetaData)
	}

	k.SetParams(ctx, genState.Params)
	k.SetTaxRateLimits(ctx, genState.TaxRateLimits)

	// Set tax caps
	for _, cap := range genState.TaxCaps {
		k.SetTaxCap(ctx, cap.Denom, cap.Cap)
	}

	// Set current period tax proceeds
	k.SetTaxProceeds(ctx, genState.TaxProceeds)

	// Set CollectionPeriod data
	for _, period := range genState.CollectionPeriods {
		k.SetPeriodTaxProceeds(ctx, int64(period.Period), period.TaxCollected)
		k.SetPayouts(ctx, int64(period.Period), period.Payouts)
	}

	//Ensure charity collector module account is set
	if k.GetCharityCollectorAcc(ctx) == nil {
		panic(fmt.Sprintf("Module account not set: %s", types.CharityCollectorName))
	}

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	taxRateLimits := k.GetTaxRateLimits(ctx)
	params := k.GetAllParams(ctx)
	taxCaps := k.GetTaxCaps(ctx)
	taxProceeds := k.GetTaxProceeds(ctx)
	collectionPeriods := k.GetCollectionPeriods(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return types.NewGenesisState(params, taxRateLimits, taxCaps, taxProceeds, collectionPeriods)
}

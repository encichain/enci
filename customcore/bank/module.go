package bank

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	customcli "github.com/encichain/enci/customcore/bank/client/cli"
	customsim "github.com/encichain/enci/customcore/bank/simulation"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModule           = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

// AppModuleBasic defines the basic application module used by the distribution module.
type AppModuleBasic struct {
	bank.AppModuleBasic
}

// GetTxCmd returns the root tx command for the bank module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return customcli.NewTxCmd()
}

// AppModule implements an application module for the bank module.
type AppModule struct {
	bank.AppModule
	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, keeper keeper.Keeper, accountKeeper types.AccountKeeper) AppModule {
	return AppModule{
		AppModule:     bank.NewAppModule(cdc, keeper, accountKeeper),
		keeper:        keeper,
		accountKeeper: accountKeeper,
	}
}

// WeightedOperations return random bank module operation.
// A custom bank WeightedOperations is implemented to prevent simulation failure caused by added charity Tax fees
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return customsim.WeightedOperations(
		simState.AppParams, simState.Cdc, am.accountKeeper, am.keeper,
	)
}

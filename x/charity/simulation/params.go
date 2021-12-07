package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/user/encichain/x/charity/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamKeyCharities),
			func(r *rand.Rand) string {
				bz, _ := json.Marshal(GenCharities(r))
				return string(bz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamKeyTaxCaps),
			func(r *rand.Rand) string {
				bz, _ := json.Marshal(GenTaxCaps(r))
				return string(bz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamKeyTaxRate),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenTaxRate(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.ParamKeyBurnRate),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenBurnRate(r))
			},
		),
	}
}

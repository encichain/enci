package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this line is used by starport scaffolding # genesis/types/import
// this line is used by starport scaffolding # ibc/genesistype/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		Params:            DefaultParams(),
		TaxRateLimits:     DefaultTaxRateLimits,
		TaxCaps:           DefaultTaxCaps,
		TaxProceeds:       DefaultTaxProceeds,
		CollectionPeriods: []CollectionPeriod{},
	}
}

// NewGenesis returns a new genesisState object. NOTE: For use with ExportGenesis
func NewGenesisState(params Params, taxratelimits TaxRateLimits, taxcaps []TaxCap, taxproceeds sdk.Coins, collection_periods []CollectionPeriod) *GenesisState {
	return &GenesisState{
		Params:            params,
		TaxRateLimits:     taxratelimits,
		TaxCaps:           taxcaps,
		TaxProceeds:       taxproceeds,
		CollectionPeriods: collection_periods,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	if gs.Params.TaxRate.LT(gs.TaxRateLimits.RateMin) {
		return fmt.Errorf("TaxRate must be greater than RateMin(%s)", gs.TaxRateLimits.RateMin)
	}

	if gs.Params.TaxRate.GT(gs.TaxRateLimits.RateMax) {
		return fmt.Errorf("TaxRate must be less than RateMax(%s)", gs.TaxRateLimits.RateMin)
	}

	return nil
	// this line is used by starport scaffolding # genesis/types/validate
}

// GetGenesisStateFromAppState returns x/charity GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

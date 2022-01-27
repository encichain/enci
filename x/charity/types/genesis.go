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
		Params:           DefaultParams(),
		TaxRateLimits:    DefaultTaxRateLimits,
		TaxCaps:          DefaultTaxCaps,
		TaxProceeds:      DefaultTaxProceeds,
		CollectionEpochs: []CollectionEpoch{},
	}
}

// NewGenesis returns a new genesisState object. NOTE: For use with ExportGenesis
func NewGenesisState(params Params, taxratelimits TaxRateLimits, taxcaps []TaxCap, taxproceeds sdk.Coins, collection_epochs []CollectionEpoch) *GenesisState {
	return &GenesisState{
		Params:           params,
		TaxRateLimits:    taxratelimits,
		TaxCaps:          taxcaps,
		TaxProceeds:      taxproceeds,
		CollectionEpochs: collection_epochs,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	//TaxRateLimits cannot be nil
	if gs.TaxRateLimits.RateMin.IsNil() || gs.TaxRateLimits.TaxRateMax.IsNil() || gs.TaxRateLimits.BurnRateMax.IsNil() {
		return fmt.Errorf("rateMin(%s), taxRateMax(%s), burnRateMax(%s) should not be nil", gs.TaxRateLimits.RateMin, gs.TaxRateLimits.TaxRateMax, &gs.TaxRateLimits.BurnRateMax)
	}
	//RateMin cannot be negative && RateMax cannot be greater than DefaultTaxRateMax && BurnRateMax cannot be greater than DefaultBurnRateMax
	if gs.TaxRateLimits.RateMin.IsNegative() || gs.TaxRateLimits.TaxRateMax.GT(DefaultTaxRateMax) || gs.TaxRateLimits.BurnRateMax.GT(DefaultBurnRateMax) {
		return fmt.Errorf("rateMin(%s) must be positive, taxRateMax(%s) must be less than %s, and burnRateMax less than %s",
			gs.TaxRateLimits.RateMin, gs.TaxRateLimits.TaxRateMax, DefaultTaxRateMax, DefaultBurnRateMax)
	}

	//TaxRate must be within range of TaxRateLimits
	if gs.Params.TaxRate.LT(gs.TaxRateLimits.RateMin) || gs.Params.TaxRate.GT(gs.TaxRateLimits.TaxRateMax) {
		return fmt.Errorf("taxRate must be greater than RateMin(%s) and less than RateMax(%s)", gs.TaxRateLimits.RateMin, &gs.TaxRateLimits.TaxRateMax)
	}

	//BurnRate must be within range of TaxRateLimits
	if gs.Params.BurnRate.LT(gs.TaxRateLimits.RateMin) || gs.Params.BurnRate.GT(gs.TaxRateLimits.BurnRateMax) {
		return fmt.Errorf("burnRate must be greater than RateMin(%s) and less than RateMax(%s)", gs.TaxRateLimits.RateMin, &gs.TaxRateLimits.BurnRateMax)
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

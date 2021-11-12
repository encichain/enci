package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	coretypes "github.com/user/encichain/types"
)

//Parameter keys
var (
	ParamKeyCharities = []byte("Charities")
	ParamKeyTaxCaps   = []byte("Taxcaps")
	ParamKeyTaxRate   = []byte("TaxRate")
)

// Default values
var (
	DefaultTaxRate = sdk.NewDecWithPrec(5, 3)   // 0.005 || 0.5%
	DefaultCap     = sdk.NewInt(int64(1000000)) // 1000000 utoken or 1 token
	DefaultTaxCaps = []TaxCap{{
		Denom: coretypes.MicroTokenDenom,
		Cap:   DefaultCap,
	}}
	DefaultCharity = Charity{
		CharityName: "",
		AccAddress:  "",
		Checksum:    "",
	}
	DefaultCharities     = []Charity{DefaultCharity}
	DefaultRateMin       = sdk.NewDecWithPrec(1, 3) // 0.001 || 0.1%
	DefaultRateMax       = sdk.NewDecWithPrec(1, 2) // 0.01 || 1%
	DefaultTaxRateLimits = TaxRateLimits{RateMin: DefaultRateMin, RateMax: DefaultRateMax}
	DefaultCoinProceed   = sdk.Coin{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(100)}
	DefaultTaxProceeds   = sdk.Coins{}
	DefaultParamsSet     = Params{
		Charities: DefaultCharities,
		TaxCaps:   DefaultTaxCaps,
		TaxRate:   DefaultTaxRate,
	}
)

var _ paramstypes.ParamSet = &Params{}

// DefaultParams creates default empty param charity sets
func DefaultParams() Params {
	return DefaultParamsSet
}

// ParamKeyTable returns the param key table for the charity module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements paramstypes.ParamSet interface. Returns ParamSetPairs (key/value pairs)
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(ParamKeyCharities, &p.Charities, validateCharities),
		paramstypes.NewParamSetPair(ParamKeyTaxRate, &p.TaxRate, validateTaxRate),
		paramstypes.NewParamSetPair(ParamKeyTaxCaps, &p.TaxCaps, validateTaxCaps),
	}
}

// Validate performs basic validation on charity parameters.
func (p Params) Validate() error {
	// Validate charities
	for _, charity := range p.Charities {
		addrlength := len([]rune(charity.AccAddress))
		if charity.AccAddress != "" {
			if addrlength < 39 {
				return fmt.Errorf("invalid address length")
			}
		}

		hashlength := len([]rune(charity.Checksum))
		if charity.Checksum != "" && hashlength != 64 {
			return fmt.Errorf("invalid sha256 length")
		}
	}

	// validate taxrate
	if p.TaxRate.IsNegative() {
		return fmt.Errorf("tax Rate must be positive")
	}

	// Validate taxcaps
	for _, taxcap := range p.TaxCaps {

		err := sdk.ValidateDenom(taxcap.Denom)

		if err != nil {
			return fmt.Errorf("taxCap Denom must be valid")
		}

		if taxcap.Cap.IsNegative() || taxcap.Cap.IsZero() || taxcap.Cap.IsNil() {
			return fmt.Errorf("taxCap Cap is invalid: Must not be negative, 0, nor nil")
		}
	}

	return nil
}

// validateCharity performs basic validation on charity parameter objects
func validateCharities(i interface{}) error {
	// Type check
	v, ok := i.([]Charity)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T. Expected []Charity", i)
	}

	// Iterate charities
	for _, charity := range v {
		addrlength := len([]rune(charity.AccAddress))
		if charity.AccAddress != "" {
			if addrlength < 39 {
				return fmt.Errorf("invalid address length")
			}
		}

		hashlength := len([]rune(charity.Checksum))
		if charity.Checksum != "" && hashlength != 64 {
			return fmt.Errorf("invalid sha256 hash length")
		}
	}
	return nil
}

// validateTaxRate performs basic validation on TaxRate
func validateTaxRate(i interface{}) error {
	// Type check
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T. Expected sdk.Dec", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("tax Rate must be positive")
	}

	return nil
}

// validateTaxCap performs basic validation on TaxCap
func validateTaxCaps(i interface{}) error {
	// Type check
	v, ok := i.([]TaxCap)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T. Expected sdk.Int", i)
	}
	// Iterate tax caps
	for _, taxcap := range v {

		err := sdk.ValidateDenom(taxcap.Denom)

		if err != nil {
			return fmt.Errorf("taxCap Denom must be valid")
		}

		if taxcap.Cap.IsNegative() || taxcap.Cap.IsZero() || taxcap.Cap.IsNil() {
			return fmt.Errorf("taxCap Cap is invalid: Must not be negative, 0, nor nil")
		}
	}
	return nil
}

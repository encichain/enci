package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

//Parameter keys
var (
	KeyCharity = []byte("Charity")
	KeyTaxCap  = []byte("Taxcap")
	KeyTaxRate = []byte("TaxRate")
)

// Default values
var (
	DefaultTaxRate   = sdk.NewDecWithPrec(1, 1)   // 0.1 || 10%
	DefaultTaxCap    = sdk.NewInt(int64(1000000)) // 1000000 utoken or 1 token
	DefaultCharity   = Charity{CharityName: "", AccAddress: "", Checksum: ""}
	DefaultParamsSet = Params{
		Charity: DefaultCharity,
		TaxCap:  DefaultTaxCap,
		TaxRate: DefaultTaxRate,
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
		paramstypes.NewParamSetPair(KeyCharity, &p.Charity, validateCharity),
		paramstypes.NewParamSetPair(KeyTaxRate, &p.TaxRate, validateTaxRate),
		paramstypes.NewParamSetPair(KeyTaxCap, &p.TaxCap, validateTaxCap),
	}
}

// Validate performs basic validation on charity parameters.
func (p Params) Validate() error {
	addrlength := len([]rune(p.Charity.AccAddress))

	if p.Charity.AccAddress != "" {
		if addrlength < 39 {
			return fmt.Errorf("invalid address length")
		}
	}

	hashlength := len([]rune(p.Charity.Checksum))
	if p.Charity.Checksum != "" && hashlength != 64 {
		return fmt.Errorf("invalid sha256 hash length")
	}

	return nil
}

// validateCharity performs basic validation on charity parameter objects
func validateCharity(i interface{}) error {
	// Type check
	v, ok := i.(Charity)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T. Expected Charity object", i)
	}
	addrlength := len([]rune(v.AccAddress))
	if v.AccAddress != "" {
		if addrlength < 39 {
			return fmt.Errorf("invalid address length")
		}
	}

	hashlength := len([]rune(v.Checksum))
	if v.Checksum != "" && hashlength != 64 {
		return fmt.Errorf("invalid sha256 hash length")
	}
	return nil
}

// validateTaxRate performs basic validation on TaxRate
func validateTaxRate(i interface{}) error {
	// Type check
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("Invalid parameter type: %T. Expected sdk.Dec", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("Tax Rate must be positive")
	}

	return nil
}

// validateTaxCap performs basic validation on TaxCap
func validateTaxCap(i interface{}) error {
	// Type check
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("Invalid parameter type: %T. Expected sdk.Int", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("Tax cap must be positive")
	}
	return nil
}

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

//Parameter keys
var (
	KeyCharityOne = []byte("charityOne")
	KeyCharityTwo = []byte("charityTwo")
)

// Default values
var (
	DefaultTaxRate    = sdk.NewDecWithPrec(1, 1) // 0.1 || 10%
	DefaultCharityOne = Charity{AccAddress: ""}
	DefaultCharityTwo = Charity{AccAddress: ""}
	DefaultParamsSet  = Params{
		CharityOne: DefaultCharityOne,
		CharityTwo: DefaultCharityTwo,
	}
)

var _ paramstypes.ParamSet = &Params{}

// DefaultParams creates default empty param charity sets
func DefaultParams() Params {
	return Params{
		CharityOne: DefaultCharityOne,
		CharityTwo: DefaultCharityTwo,
	}
}

// ParamKeyTable returns the param key table for the charity module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements paramstypes.ParamSet interface. Returns ParamSetPairs (key/value pairs)
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyCharityOne, &p.CharityOne, validateCharity),
		paramstypes.NewParamSetPair(KeyCharityTwo, &p.CharityTwo, validateCharity),
	}
}

// Validate performs basic validation on charity parameters.
func (p Params) Validate() error {
	addrlength := len([]rune(p.CharityOne.AccAddress))

	if p.CharityOne.AccAddress != "" {
		if addrlength < 39 {
			return fmt.Errorf("invalid address length")
		}
	}

	addrlength = len([]rune(p.CharityTwo.AccAddress))
	if p.CharityTwo.AccAddress != "" {
		if addrlength < 39 {
			return fmt.Errorf("invalid address length")
		}
	}
	return nil
}

// validateCharity performs basic validation on charity parameter objects
func validateCharity(i interface{}) error {
	v, ok := i.(Charity)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	addrlength := len([]rune(v.AccAddress))
	if v.AccAddress != "" {
		if addrlength < 39 {
			return fmt.Errorf("invalid address length")
		}
	}
	return nil
}

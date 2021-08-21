package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultTaxRate = sdk.NewDecWithPrec(1, 1) // 0.1 || 10%
)

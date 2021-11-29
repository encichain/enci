package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidate(t *testing.T) {

	testCases := []struct {
		name     string
		genState GenesisState
		expErr   bool
	}{
		{
			"Valid genesis",
			*DefaultGenesis(),
			false,
		},
		// Empty params results in nil pointer dereference due to TaxRate.IsNeg() check in Validate()
		{
			"Empty genesis",
			GenesisState{},
			true,
		},
		{
			"invalid params",
			GenesisState{
				Params{
					Charities: DefaultCharities,
					TaxRate:   DefaultTaxRate.Neg(),
					TaxCaps:   DefaultTaxCaps,
				},
				DefaultTaxRateLimits,
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
		{
			"negative rateMin",
			GenesisState{
				DefaultParams(),
				TaxRateLimits{RateMin: sdk.NewDecWithPrec(1, 4).Neg(), TaxRateMax: DefaultTaxRateMax, BurnRateMax: DefaultBurnRateMax},
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
		{
			"too high TaxRateMax",
			GenesisState{
				DefaultParams(),
				TaxRateLimits{RateMin: DefaultRateMin, TaxRateMax: sdk.NewDecWithPrec(6, 2), BurnRateMax: DefaultBurnRateMax},
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
		{
			"too high BurnRateMax",
			GenesisState{
				DefaultParams(),
				TaxRateLimits{RateMin: DefaultRateMin, TaxRateMax: DefaultTaxRateMax, BurnRateMax: sdk.NewDecWithPrec(60, 2)},
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
		{
			"taxRate lower than RateMin",
			GenesisState{
				Params{
					Charities: DefaultCharities,
					TaxRate:   DefaultRateMin.Sub(sdk.NewDecWithPrec(1, 4)),
					TaxCaps:   DefaultTaxCaps,
				},
				DefaultTaxRateLimits,
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
		{
			"taxRate higher than RateMin",
			GenesisState{
				Params{
					Charities: DefaultCharities,
					TaxRate:   DefaultTaxRateMax.Add(sdk.NewDecWithPrec(1, 4)),
					TaxCaps:   DefaultTaxCaps,
				},
				DefaultTaxRateLimits,
				DefaultTaxCaps,
				sdk.Coins{},
				[]CollectionPeriod{},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			err := tc.genState.Validate()

			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

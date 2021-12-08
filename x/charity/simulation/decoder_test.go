package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	app "github.com/user/encichain/app"
	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/simulation"
	"github.com/user/encichain/x/charity/types"
)

func TestDecodeDistributionStore(t *testing.T) {
	cdc := app.MakeTestEncodingConfig().Codec
	dec := simulation.NewDecodeStore(cdc)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 705000))

	taxRateLimits := types.TaxRateLimits{RateMin: sdk.ZeroDec(), TaxRateMax: sdk.NewDecWithPrec(123, 3), BurnRateMax: sdk.NewDecWithPrec(123, 3)}
	taxCap := sdk.IntProto{Int: sdk.NewInt(12345678)}
	taxProceeds := types.TaxProceeds{TaxProceeds: sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 9876543))}
	periodTaxProceeds := types.TaxProceeds{TaxProceeds: sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 234567))}
	periodPayouts := types.Payouts{Payouts: []types.Payout{
		{Recipientaddr: "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55", Coins: coins},
		{Recipientaddr: "TEST", Coins: coins.Add(coins...)},
	},
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.TaxRateLimitsKey, Value: cdc.MustMarshal(&taxRateLimits)},
			{Key: types.TaxCapKeyPref, Value: cdc.MustMarshal(&taxCap)},
			{Key: types.TaxProceedsKey, Value: cdc.MustMarshal(&taxProceeds)},
			{Key: types.PeriodTaxProceedsKeyPref, Value: cdc.MustMarshal(&periodTaxProceeds)},
			{Key: types.PayoutsKeyPref, Value: cdc.MustMarshal(&periodPayouts)},
			{Key: []byte{0x15}, Value: []byte{0x15}},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"TaxRateLimits", fmt.Sprintf("%v\n%v", taxRateLimits, taxRateLimits)},
		{"TaxCap", fmt.Sprintf("%v\n%v", taxCap, taxCap)},
		{"TaxProceeds", fmt.Sprintf("%v\n%v", taxProceeds, taxProceeds)},
		{"PeriodTaxProceeds", fmt.Sprintf("%v\n%v", periodTaxProceeds, periodTaxProceeds)},
		{"Payouts", fmt.Sprintf("%v\n%v", periodPayouts, periodPayouts)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}

package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/encichain/enci/x/charity/types"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding charity type.
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.TaxRateLimitsKey):
			var taxRateLimitsA, taxRateLimitsB types.TaxRateLimits
			cdc.MustUnmarshal(kvA.Value, &taxRateLimitsA)
			cdc.MustUnmarshal(kvB.Value, &taxRateLimitsB)
			return fmt.Sprintf("%v\n%v", taxRateLimitsA, taxRateLimitsB)

		case bytes.Equal(kvA.Key[:1], types.TaxCapKeyPref):
			var taxCapA, taxCapB sdk.IntProto
			cdc.MustUnmarshal(kvA.Value, &taxCapA)
			cdc.MustUnmarshal(kvB.Value, &taxCapB)
			return fmt.Sprintf("%v\n%v", taxCapA, taxCapB)

		case bytes.Equal(kvA.Key[:1], types.TaxProceedsKey):
			var taxProceedsA, taxProceedsB types.TaxProceeds
			cdc.MustUnmarshal(kvA.Value, &taxProceedsA)
			cdc.MustUnmarshal(kvB.Value, &taxProceedsB)
			return fmt.Sprintf("%v\n%v", taxProceedsA, taxProceedsB)

		case bytes.Equal(kvA.Key[:1], types.EpochTaxProceedsKeyPref):
			var epochTaxProceedsA, epochTaxProceedsB types.TaxProceeds
			cdc.MustUnmarshal(kvA.Value, &epochTaxProceedsA)
			cdc.MustUnmarshal(kvB.Value, &epochTaxProceedsB)
			return fmt.Sprintf("%v\n%v", epochTaxProceedsA, epochTaxProceedsB)

		case bytes.Equal(kvA.Key[:1], types.PayoutsKeyPref):
			var epochPayoutsA, epochPayoutsB types.Payouts
			cdc.MustUnmarshal(kvA.Value, &epochPayoutsA)
			cdc.MustUnmarshal(kvB.Value, &epochPayoutsB)
			return fmt.Sprintf("%v\n%v", epochPayoutsA, epochPayoutsB)

		default:
			panic(fmt.Sprintf("invalid charity key prefix %X", kvA.Key[:1]))
		}
	}
}

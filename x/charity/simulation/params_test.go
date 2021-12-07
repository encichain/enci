package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/user/encichain/x/charity/simulation"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	expected := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"charity/Charities", "Charities", "", "charity"},
		{"charity/Taxcaps", "Taxcaps", "[{\"denom\":\"uenci\",\"Cap\":\"7779410\"}]", "charity"},
		{"charity/TaxRate", "TaxRate", "\"0.008000000000000000\"", "charity"},
		{"charity/BurnRate", "BurnRate", "\"0.370000000000000000\"", "charity"},
	}

	paramChanges := simulation.ParamChanges(r)

	require.Len(t, paramChanges, 4)

	for i, p := range paramChanges {
		if p.Key() == "Charities" {
			continue
		}
		require.Equal(t, expected[i].composedKey, p.ComposedKey())
		require.Equal(t, expected[i].key, p.Key())
		require.Equal(t, expected[i].simValue, p.SimValue()(r))
		require.Equal(t, expected[i].subspace, p.Subspace())
	}
}

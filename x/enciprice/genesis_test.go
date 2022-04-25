package enciprice_test

import (
	"testing"

	keepertest "github.com/encichain/enci/testutil/keeper"
	"github.com/encichain/enci/testutil/nullify"
	"github.com/encichain/enci/x/enciprice"
	"github.com/encichain/enci/x/enciprice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EncipriceKeeper(t)
	enciprice.InitGenesis(ctx, *k, genesisState)
	got := enciprice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}

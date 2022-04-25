package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/encichain/enci/testutil/keeper"
	"github.com/encichain/enci/x/enciprice/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EncipriceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

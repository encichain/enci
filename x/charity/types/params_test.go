package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	params := DefaultParams()
	require.NoError(t, params.Validate())

	// Set account address to non-empty string of invalid length. Expect: error
	params.CharityOne.AccAddress = "invalidlengthtest"
	require.Error(t, params.Validate())

	// Set account address of both charity objects to correct format. Expect: No error
	params = DefaultParams()
	params.CharityOne.AccAddress = "cosmos1065sa0t8v3dwjxhkchttmzf489she4yjfeq7nx"
	params.CharityTwo.AccAddress = "cosmos1wlmzwvk8pxertaef4cjfqux7e7esdcmu973twc"
	require.NoError(t, params.Validate())

	// Change one account address to invalid length
	params.CharityOne.AccAddress = "1065sa0t8v3dwjxhkchttmzf489she4yjfeq7n"
	require.Error(t, params.Validate())

	// Settings param set pairs should not result in nil
	require.NotNil(t, params.ParamSetPairs())
	params = DefaultParams()
	require.NotNil(t, params.ParamSetPairs())

}

package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	p1 := DefaultParams()
	require.NoError(t, p1.Validate())

	// Set account address to non-empty string of invalid length. Expect: error
	p1.Charity.AccAddress = "invalidlengthtest"
	require.Error(t, p1.Validate())

	// Set account address of charity objects to correct format. Expect: No error
	p2 := DefaultParams()
	p2.Charity.AccAddress = "cosmos1065sa0t8v3dwjxhkchttmzf489she4yjfeq7nx"
	require.NoError(t, p2.Validate())

	// Change account address to invalid length
	p2.Charity.AccAddress = "1065sa0t8v3dwjxhkchttmzf489she4yjfeq7n"
	require.Error(t, p2.Validate())

	// Set checksum to correct length of 256 bits - 64 characters
	p3 := DefaultParams()
	p3.Charity.Checksum = "D821F6C986794B80524D765385A106C9D68F79FA6EAD18FF2F79F58B45DAA848"
	require.NoError(t, p3.Validate())

	// Set checksum to incorrect length of 63 characters
	p3.Charity.Checksum = "D821F6C986794B80524D765385A106C9D68F79FA6EAD18FF2F79F58B45DAA84"
	require.Error(t, p3.Validate())

	// Settings param set pairs should not result in nil
	p4 := DefaultParams()
	require.NotNil(t, p4.ParamSetPairs())
	require.NoError(t, p4.Validate())

}

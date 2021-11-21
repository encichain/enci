package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"
)

func TestParams(t *testing.T) {
	defaultParamsSet := DefaultParams()
	p1 := defaultParamsSet
	require.NoError(t, p1.Validate())

	// Set account address to non-empty string of invalid length. Expect: error
	p1.Charities[0].AccAddress = "invalidlengthtest"
	require.Error(t, p1.Validate())

	// Change account address to invalid length
	p1.Charities[0].AccAddress = "1065sa0t8v3dwjxhkchttmzf489she4yjfeq7n"
	require.Error(t, p1.Validate())

	// Set account address of charity objects to correct format. Expect: No error
	p1.Charities[0].AccAddress = "cosmos1065sa0t8v3dwjxhkchttmzf489she4yjfeq7nx"
	require.NoError(t, p1.Validate())

	// Set checksum to incorrect length of 63 characters
	p1.Charities[0].Checksum = "D821F6C986794B80524D765385A106C9D68F79FA6EAD18FF2F79F58B45DAA84"
	err := p1.Validate()
	require.Error(t, err)

	// Set checksum to correct length of 256 bits - 64 characters
	p1.Charities[0].AccAddress = ""
	p1.Charities[0].Checksum = "D821F6C986794B80524D765385A106C9D68F79FA6EAD18FF2F79F58B45DAA848"
	require.NoError(t, p1.Validate())

	// negative taxrate
	p1.TaxRate = DefaultTaxRate.Neg()
	require.Error(t, p1.Validate())

	// zero cap
	p1.TaxCaps = []TaxCap{{Denom: coretypes.MicroTokenDenom, Cap: sdk.NewInt(int64(0))}}
	require.Error(t, p1.Validate())

	//negative cap
	p1.TaxCaps = []TaxCap{{Denom: coretypes.MicroTokenDenom, Cap: DefaultCap.Neg()}}
	require.Error(t, p1.Validate())

	// Setting param set pairs should not result in nil
	p2 := DefaultParams()
	require.Equal(t, defaultParamsSet, p2)
	require.NotNil(t, p2.ParamSetPairs())
	require.NoError(t, p2.Validate())

}

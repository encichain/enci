package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestClaimLengthPrefix(t *testing.T) {
	valAddr := sdk.ValAddress("test________________")
	claimType := "test"
	key := types.GetVoteKey(valAddr, claimType)
	typeBz := key[2 : len(claimType)+2]
	require.Equal(t, claimType, string(typeBz))
}

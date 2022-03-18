package types_test

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestVoteHash(t *testing.T) {
	claim := types.NewTestClaim(0, "test", "test")
	claimHash := claim.Hash()
	valAddr := sdk.ValAddress([]byte("encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee"))

	voteHash := types.CreateVoteHash("1", claimHash.String(), valAddr)
	hexStr := hex.EncodeToString(voteHash)
	require.Equal(t, hexStr, voteHash.String())

	voteHashFromStr, err := types.HexStringToVoteHash(hexStr)
	require.NoError(t, err)
	require.Equal(t, voteHash, voteHashFromStr)

	require.True(t, types.VoteHash([]byte{}).Empty())
}

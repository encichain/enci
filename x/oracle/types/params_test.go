package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	params := types.DefaultParams()
	// Valid
	err := params.Validate()
	require.NoError(t, err)

	// Zero VotePeriod
	params.VotePeriod = 0
	err = params.Validate()
	require.Error(t, err)
	//Reset
	params.VotePeriod = 3

	// zero prevotePeriod
	params.PrevotePeriod = 0
	err = params.Validate()
	require.Error(t, err)
	params.PrevotePeriod = 3

	// Zero VoteFrequency
	params.VoteFrequency = 0
	err = params.Validate()
	require.Error(t, err)
	params.VoteFrequency = 100

	// Less than 33% VoteThreshold
	params.VoteThreshold = sdk.NewDecWithPrec(30, 2)
	params.Validate()
	require.Error(t, err)
	// Greater than 100%
	params.VoteThreshold = sdk.NewDecWithPrec(101, 2)
	params.Validate()
	require.Error(t, err)
}

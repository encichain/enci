package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	coreapp "github.com/encichain/enci/app"
	oracle "github.com/encichain/enci/x/oracle"
	"github.com/encichain/enci/x/oracle/types"
)

func newTestVote(validatorAddress string) types.Vote {
	valAddr, _ := sdk.ValAddressFromBech32(validatorAddress)
	testVote, _ := types.NewVote(types.NewTestClaim(3, "test", "test"), valAddr, 100)
	return testVote
}

func newTestPrevote(validatorAddress string) types.Prevote {
	valAddr, _ := sdk.ValAddressFromBech32(validatorAddress)
	claim := types.NewTestClaim(0, "test", "test")
	voteHash := types.CreateVoteHash("0", claim.Hash().String(), valAddr)
	testPrevote := types.NewPrevote(voteHash, valAddr, 0)
	return testPrevote
}

func TestExportImportGenesis(t *testing.T) {
	params := types.DefaultParams()
	params.VoteFrequency = 20000
	valAddr := "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee"
	voterDelegations := []types.VoterDelegation{
		{
			DelegateAddress:  "enci16ruw3nnsrt963y47y8m8h0g6p4pkyudvm5j3fc",
			ValidatorAddress: valAddr,
		},
	}
	claimTypes := []types.ClaimType{
		{ClaimType: "test"},
	}
	// Create initial genesis state
	initialGenState := types.NewGenesisState(params, voterDelegations, []types.VoteRound{types.NewVoteRound("test", []types.Vote{newTestVote(valAddr)})},
		[]types.PrevoteRound{}, claimTypes)

	// create new app and context and import gen state
	app := coreapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})
	oracle.InitGenesis(ctx, app.OracleKeeper, *initialGenState)

	// Test export genesis state and ensure it equals inital gen state
	exportedGenState := oracle.ExportGenesis(ctx, app.OracleKeeper)

	require.Equal(t, initialGenState, exportedGenState)

}

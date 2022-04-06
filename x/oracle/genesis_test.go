package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	coreapp "github.com/encichain/enci/app"
	oracle "github.com/encichain/enci/x/oracle"
	"github.com/encichain/enci/x/oracle/types"
)

func TestExportImportGenesis(t *testing.T) {
	params := types.DefaultParams()
	params.VoteFrequency = 20000
	voterDelegations := []types.VoterDelegation{
		{
			DelegateAddress:  "enci16ruw3nnsrt963y47y8m8h0g6p4pkyudvm5j3fc",
			ValidatorAddress: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee",
		},
	}
	claimTypes := []types.ClaimType{
		{ClaimType: "test"},
	}
	// Create initial genesis state
	initialGenState := types.NewGenesisState(params, voterDelegations, []types.VoteRound{}, []types.PrevoteRound{}, claimTypes)

	// create new app and context and import gen state
	app := coreapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})
	oracle.InitGenesis(ctx, app.OracleKeeper, *initialGenState)

	// Test export genesis state and ensure it equals inital gen state
	exportedGenState := oracle.ExportGenesis(ctx, app.OracleKeeper)

	require.Equal(t, initialGenState, exportedGenState)

}

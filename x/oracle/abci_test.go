package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	coreapp "github.com/encichain/enci/app"
	oracle "github.com/encichain/enci/x/oracle"
	"github.com/encichain/enci/x/oracle/types"
)

func TestEndBlocker(t *testing.T) {
	app := coreapp.Setup(false)
	// Not prevote/vote period start
	ctx := app.NewContext(false, tmproto.Header{Height: 10}).
		WithEventManager(sdk.NewEventManager())

	params := types.DefaultParams()
	params.OracleEnabled = true
	app.OracleKeeper.SetParams(ctx, params)

	oracle.EndBlocker(ctx, app.OracleKeeper)
	events := ctx.EventManager().Events()
	require.Len(t, events, 0)

	// beginning of prevote period
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency) - 1)
	isPrevoteBegin := app.OracleKeeper.IsPrevotePeriodBegin(ctx, params)
	require.True(t, isPrevoteBegin)

	oracle.EndBlocker(ctx, app.OracleKeeper)
	events = ctx.EventManager().Events()
	require.Len(t, events, 1)
	// beginning of vote period
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency+params.PrevotePeriod-1) - 1)
	isVoteBegin := app.OracleKeeper.IsVotePeriodBegin(ctx, params)
	require.True(t, isVoteBegin)

	oracle.EndBlocker(ctx, app.OracleKeeper)
	events = ctx.EventManager().Events()
	require.Len(t, events, 2)
}

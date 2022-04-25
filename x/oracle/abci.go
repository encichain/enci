package oracle

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/keeper"
	"github.com/encichain/enci/x/oracle/types"
)

// EndBlocker is called at the end of every block
// Check if last block before start of prevote/vote period, and emit event if true
// handling of votes to be done in external modules
// Possible refactor: Emit event for each block of a voting/prevoting period
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	params := k.GetParams(ctx)

	if params.OracleEnabled {
		if prevoteBegin := k.IsPrevotePeriodBegin(ctx, params); prevoteBegin {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypePrevoteBegin,
				),
			)
		} else if voteBegin := k.IsVotePeriodBegin(ctx, params); voteBegin {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeVoteBegin,
				),
			)
		}
	}
}

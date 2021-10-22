package charity

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if coretypes.IsLastBlockPeriod(ctx) {
		period := k.GetCurrentPeriod(ctx)
		charities := k.GetCharity(ctx)
		// Reset tax proceeds
		defer k.SetTaxProceeds(ctx, sdk.Coins{})

		// Disburse donations according to CharityTaxCollector balance
		payouts := k.DisburseDonations(ctx, charities)
		// Set payouts to store under current *period*
		k.SetPayouts(ctx, period, payouts)
		// Set period tax proceeds to store
		k.SetPeriodTaxProceeds(ctx, period, k.GetTaxProceeds(ctx))

	}
}

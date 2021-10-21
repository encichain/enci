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
		payouts := []types.Payout{}
		charities := k.GetCharity(ctx)

		// Get the donation split
		finalsplit := k.CalculateSplit(ctx, charities)

		// Perform charity payouts
		for _, charity := range charities {
			err := k.DonateCharity(ctx, finalsplit, charity)
			if err != nil {
				continue
			}
			payout := types.Payout{Recipientaddr: charity.AccAddress, Coins: finalsplit}
			payouts = append(payouts, payout)
		}
		// Set payouts to store under current *period*
		k.SetPayouts(ctx, period, payouts)
		// Set period tax proceeds to store
		k.SetPeriodTaxProceeds(ctx, period, k.GetTaxProceeds(ctx))
		// Reset tax proceeds
		k.SetTaxProceeds(ctx, sdk.Coins{})
	}
}

package charity

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity/keeper"
	"github.com/encichain/enci/x/charity/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	if coretypes.IsLastBlockPeriod(ctx) {
		period := k.GetCurrentPeriod(ctx)
		charities := k.GetCharity(ctx)
		charityBal := k.BankKeeper.SpendableCoins(ctx, k.AccountKeeper.GetModuleAddress(types.CharityCollectorName))
		burnAmt := sdk.Coins{}

		// Calculate burn amount from CharityTaxCollector and send to burner account
		if !charityBal.IsZero() {
			burnAmt = k.CalculateBurnAmount(ctx, charityBal)
			if !burnAmt.IsZero() {
				err := k.BankKeeper.SendCoinsFromModuleToModule(ctx, types.CharityCollectorName, types.BurnAccName, burnAmt)
				if err != nil {
					panic(fmt.Sprintf("could not send coins from CharityTaxCollector to burn account: %s", err))
				}
			}
		}
		// Burn all balances from burn module account
		err := k.BurnCoinsFromBurner(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to burn coins from burner account: %s", err))
		}

		// Disburse donations according to CharityTaxCollector balance
		payouts, errs := k.DisburseDonations(ctx, charities)

		// Set payouts to store under current *period*, if not empty
		if len(payouts) > 0 {
			k.SetPayouts(ctx, period, payouts)
		}
		// Set period tax proceeds to store, if not zero
		taxProceeds := k.GetTaxProceeds(ctx)
		if !taxProceeds.IsZero() {
			k.SetPeriodTaxProceeds(ctx, period, k.GetTaxProceeds(ctx))
		}
		// Sync taxcaps
		k.SyncTaxCaps(ctx)

		// Reset tax proceeds
		k.SetTaxProceeds(ctx, sdk.Coins{})

		ctx.EventManager().EmitTypedEvent(
			&types.EventPayout{
				Period:      uint64(period),
				Payouts:     payouts,
				BurnedCoins: burnAmt,
			})
		if len(errs) > 0 {
			ctx.EventManager().EmitTypedEvent(
				&types.EventFailedPayouts{
					Period: uint64(period),
					Errors: errs,
				})
		}
	}
}

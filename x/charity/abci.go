package charity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {

	if coretypes.IsLastBlockPeriod(ctx) {
		period := k.GetCurrentPeriod(ctx)
		payouts := []types.Payout{}
		charities := k.GetCharity(ctx)
		taxaddr := k.AccountKeeper.GetModuleAddress(types.CharityCollectorName)
		// Get balance of tax collector
		balance := k.BankKeeper.SpendableCoins(ctx, taxaddr)
		coins := []sdk.Coin{}

		if balance.IsZero() {
			return
		}

		for _, coin := range balance {
			split := sdk.NewInt(int64(len(charities)))
			sc := sdk.Coin{
				Denom:  coin.Denom,
				Amount: coin.Amount.Quo(split),
			}
			coins = append(coins, sc)
		}
		finalsplit := sdk.NewCoins(coins...)

		// Perform charity payouts
		for _, charity := range charities {
			err := k.DonateCharity(ctx, finalsplit, charity)
			if err != nil {
				continue
			}
			payout := types.Payout{}
		}
	}
	return
}

package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
)

// BurnCoinsFromBurner burns all coins from the Burner account
func (k Keeper) BurnCoinsFromBurner(ctx sdk.Context) error {
	burnerAddr := k.AccountKeeper.GetModuleAddress(types.BurnAccName)
	if burnerAddr == nil {
		panic(fmt.Sprintf("burn failed. %s account address not set", types.BurnAccName))
	}
	// Get all balances
	bals := k.BankKeeper.GetAllBalances(ctx, burnerAddr)
	// Burn coins if balance is not zero
	if !bals.IsZero() {
		err := k.BankKeeper.BurnCoins(ctx, types.BurnAccName, bals)
		if err != nil {
			return err
		}
	}
	return nil
}

// CalculateBurnAmount calculates the coins to be sent to burn account based on current app burn rate
func (k Keeper) CalculateBurnAmount(ctx sdk.Context, balance sdk.Coins) sdk.Coins {
	burnRate := k.GetBurnRate(ctx)
	coins := sdk.Coins{}
	if burnRate.IsZero() {
		return coins
	}

	if !balance.IsZero() {
		for _, coin := range balance {
			bc := sdk.NewCoin(
				coin.Denom,
				sdk.NewDecFromInt(coin.Amount).Mul(burnRate).TruncateInt(),
			)
			coins = append(coins, bc)
		}
	}
	return sdk.NewCoins(coins...)
}

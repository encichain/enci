package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
)

// DonateCharity sends proceeds to the specified charity
func (k Keeper) DonateCharity(ctx sdk.Context, proceeds sdk.Coins, charity types.Charity) error {
	err := k.IsValidCharity(ctx, charity)
	if err != nil {
		return err
	}
	addr, err := sdk.AccAddressFromBech32(charity.AccAddress)
	if err != nil {
		return err
	}
	// Try to send coins from tax collector module account to charity address
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.CharityCollectorName, addr, proceeds)
	if err != nil {
		return err
	}
	return nil
}

// IsValidCharity performs Validation on Charity object
func (k Keeper) IsValidCharity(ctx sdk.Context, charity types.Charity) error {
	// Check if Checksum is valid
	csb := sha256.Sum256([]byte(charity.CharityName + charity.AccAddress))
	checksum := hex.EncodeToString(csb[:])

	if checksum != charity.Checksum {
		return fmt.Errorf("checksum is invalid")
	}

	// Check account address
	// TODO: Use accountKeeper.HasAccount method when implemented in cosmos-sdk.
	addr, err := sdk.AccAddressFromBech32(charity.AccAddress)
	if err != nil {
		return fmt.Errorf("invalid address")
	}
	acc := k.AccountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		return fmt.Errorf("account does not exist for the provided address")
	}

	return nil
}

// CalculateSplit returns the sdk.Coins proceed donation split based on spendable balance of Charity Tax Collector account and number of charities
func (k Keeper) CalculateSplit(ctx sdk.Context, charities []types.Charity) sdk.Coins {
	taxaddr := k.AccountKeeper.GetModuleAddress(types.CharityCollectorName)
	if taxaddr == nil {
		return sdk.Coins{}
	}
	balance := k.BankKeeper.SpendableCoins(ctx, taxaddr)
	if balance.IsZero() {
		return sdk.Coins{}
	}
	coins := []sdk.Coin{}

	// Calculate donation split
	for _, coin := range balance {
		split := sdk.NewInt(int64(len(charities)))
		sc := sdk.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Quo(split),
		}
		coins = append(coins, sc)
	}
	return sdk.NewCoins(coins...)
}

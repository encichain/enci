package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/encichain/enci/x/charity/types"
)

// DisburseDonations sends funds from CharityTaxCollector to all specified charities and returns []Payout and []string representation of errors..
// Should not be called except during end of epoch.
func (k Keeper) DisburseDonations(ctx sdk.Context, charities []types.Charity) (payouts []types.Payout, errs []string) {
	payouts = []types.Payout{}
	errs = []string{}
	// Get the donation split
	finalSplit := k.CalculateSplit(ctx, charities)

	// Perform charity payouts
	for _, charity := range charities {
		err := k.DonateCharity(ctx, finalSplit, charity)
		if err != nil {
			errmsg := sdkerrors.Wrapf(err, "Payout failed for charity: %s, with error", charity.CharityName)
			errs = append(errs, errmsg.Error())
			continue
		}
		payout := types.Payout{Recipientaddr: charity.AccAddress, Coins: finalSplit}
		payouts = append(payouts, payout)
	}
	return
}

// DonateCharity sends proceeds to the specified charity
func (k Keeper) DonateCharity(ctx sdk.Context, proceeds sdk.Coins, charity types.Charity) error {
	// Validate Charity
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
	checksum := CreateCharitySha256(charity.CharityName, charity.AccAddress)
	if checksum != charity.Checksum {
		return fmt.Errorf("checksum is invalid")
	}

	// Check account address
	// TODO: Use AccountKeeper.HasAccount method when implemented in cosmos-sdk.
	addr, err := sdk.AccAddressFromBech32(charity.AccAddress)
	if err != nil {
		return fmt.Errorf("invalid address")
	}
	hasAcc := k.AccountKeeper.HasAccount(ctx, addr)
	if !hasAcc {
		return fmt.Errorf("account does not exist for the provided address")
	}

	return nil
}

// CalculateSplit returns the sdk.Coins proceed donation split based on spendable balance of Charity Tax Collector account and number of charities
func (k Keeper) CalculateSplit(ctx sdk.Context, charities []types.Charity) sdk.Coins {
	taxAddr := k.AccountKeeper.GetModuleAddress(types.CharityCollectorName)
	if taxAddr == nil {
		return sdk.Coins{}
	}
	balance := k.BankKeeper.SpendableCoins(ctx, taxAddr)
	if balance.IsZero() {
		return sdk.Coins{}
	}
	coins := []sdk.Coin{}
	split := sdk.NewInt(int64(len(charities)))
	// Calculate donation split
	for _, coin := range balance {
		sc := sdk.Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Quo(split),
		}
		coins = append(coins, sc)
	}
	return sdk.NewCoins(coins...)
}

// CreateCharitySha256 returns the hexadecimal encoding of sha256 checksum of JSON(charity name + charity accAddress(Bech32))
func CreateCharitySha256(charityName string, accAddr string) string {
	bz, _ := json.Marshal(charityName + accAddr)

	csb := sha256.Sum256(bz)
	checksum := hex.EncodeToString(csb[:])
	return checksum
}

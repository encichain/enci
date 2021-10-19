package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/user/encichain/x/charity/types"
)

// IsValidCharity performs checks on Charity object
func (k Keeper) IsValidCharity(ctx sdk.Context, charity types.Charity) (bool, error) {
	// Check if Checksum is valid
	csb := sha256.Sum256([]byte(charity.CharityName + charity.AccAddress))
	checksum := hex.EncodeToString(csb[:])

	if checksum != charity.Checksum {
		return false, fmt.Errorf("checksum is invalid")
	}

	// Check account address
	// TODO: Use accountKeeper.HasAccount method when implemented in cosmos-sdk.
	addr, err := sdk.AccAddressFromBech32(charity.AccAddress)
	if err != nil {
		return false, fmt.Errorf("invalid address")
	}
	acc := k.accountKeeper.GetAccount(ctx, addr)
	if acc == nil {
		return false, fmt.Errorf("Aacount does not exist for the provided address")
	}

	return true, nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper defines the interface for the expected staking keeper
type StakingKeeper interface {
	Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool) // get validator by operator address; nil when validator not found
	TotalBondedTokens(sdk.Context) sdk.Int                                                            // total bonded tokens within the validator set
}

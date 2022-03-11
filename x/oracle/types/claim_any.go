package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/encichain/enci/x/oracle/exported"
)

// GetClaim extracts a claim from a Vote
func (v Vote) GetClaim() (exported.Claim, error) {
	claim, ok := v.Claim.GetCachedValue().(exported.Claim)
	if !ok {
		return nil, sdkerrors.Wrap(ErrInvalidClaim, "could not get claim")
	}
	return claim, nil
}

// Vote receiver UnpackInterfaces unpacks Any Claim cached value
func (v Vote) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var c exported.Claim
	return unpacker.UnpackAny(v.Claim, &c)
}

// VoteRound receiver UnpackInterfaces unpacks Any Claim cached value for all Votes in VoteRound
func (voteRound VoteRound) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, v := range voteRound.Votes {
		err := v.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

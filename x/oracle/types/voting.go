package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/exported"
)

// NewPrevote creates a prevote object
func NewPrevote(hash VoteHash, validator sdk.ValAddress, submitBlock uint64) Prevote {
	return Prevote{
		Hash:        hash.String(),
		Validator:   validator.String(),
		SubmitBlock: submitBlock,
	}
}

// NewVote creates a Vote object
func NewVote(claim exported.Claim, validator sdk.ValAddress, votePower uint64) (Vote, error) {
	claimAny, err := codectypes.NewAnyWithValue(claim)
	if err != nil {
		return Vote{}, err
	}
	return Vote{
		Claim:     claimAny,
		Validator: validator.String(),
		VotePower: votePower,
	}, nil
}

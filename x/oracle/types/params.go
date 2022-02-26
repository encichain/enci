package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	coretypes "github.com/encichain/enci/types"
)

// Parameter keys
var (
	KeyVoteFrequency = []byte("VoteFrequency")
	KeyClaimParams   = []byte("ClaimParams")
)

// Default params for testing
var (
	TestClaimType               = "test"
	TestPrevoteClaimType        = "prevoteTest"
	TestVotePeriod       uint64 = 3
)

// Default parameter values
var (
	DefaultVoteFrequency = coretypes.BlocksPerDay
	DefaultVoteThreshold = sdk.NewDecWithPrec(50, 2) // 0.50 -> 50%
	DefaultVotePeriod    = uint64(4)
	DefaultPrevotePeriod = uint64(4)
	DefaultClaimParams   = map[string](ClaimParams){}
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable for oracle module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {
	return Params{
		VoteFrequency: DefaultVoteFrequency,
		ClaimParams:   DefaultClaimParams,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVoteFrequency, &p.VoteFrequency, validateVoteFrequency),
		paramtypes.NewParamSetPair(KeyClaimParams, &p.ClaimParams, validateClaimParams),
	}
}

// ValidateBasic performs basic validation on oracle parameters.
func (p Params) Validate() error {
	if p.VoteFrequency < 1 {
		return fmt.Errorf("invalid vote frequency: %d", p.VoteFrequency)
	}

	// Validate claim params
	for _, param := range p.ClaimParams {
		if param.VotePeriod <= 0 {
			return fmt.Errorf("vote period must be greater than 0: %d", param.VotePeriod)
		}
		if param.PrevotePeriod <= 0 {
			return fmt.Errorf("prevote period must be greater than 0: %d", param.PrevotePeriod)
		}
		if param.VoteThreshold.LTE(sdk.NewDecWithPrec(33, 2)) {
			return fmt.Errorf("oracle parameter VoteTheshold must be greater than 33 percent")
		}
		if param.VoteThreshold.GT(sdk.OneDec()) {
			return fmt.Errorf("vote threshold too large: %s", param.VoteThreshold)
		}
	}
	return nil
}

func validateClaimParams(i interface{}) error {
	claimParams, ok := i.(map[string](ClaimParams))
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, param := range claimParams {
		if param.VotePeriod <= 0 {
			return fmt.Errorf("vote period must be greater than 0: %d", param.VotePeriod)
		}
		if param.PrevotePeriod <= 0 {
			return fmt.Errorf("prevote period must be greater than 0: %d", param.PrevotePeriod)
		}
		if param.VoteThreshold.LTE(sdk.NewDecWithPrec(33, 2)) {
			return fmt.Errorf("oracle parameter VoteTheshold must be greater than 33 percent")
		}
		if param.VoteThreshold.GT(sdk.OneDec()) {
			return fmt.Errorf("vote threshold too large: %s", param.VoteThreshold)
		}
	}
	return nil
}

func validateVoteFrequency(i interface{}) error {
	v, ok := i.(uint64)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 1 {
		return fmt.Errorf("invalid vote frequency: %d", v)
	}

	return nil
}

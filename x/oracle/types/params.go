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
	KeyPrevotePeriod = []byte("PrevotePeriod")
	KeyVotePeriod    = []byte("VotePeriod")
	KeyVoteThreshold = []byte("VoteThreshold")
	KeyOracleEnabled = []byte("OracleEnabled")
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
	DefaultVotePeriod    = uint64(3)
	DefaultPrevotePeriod = uint64(3)
	DefaultOracleEnabled = false
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable for oracle module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {
	return Params{
		PrevotePeriod: DefaultPrevotePeriod,
		VotePeriod:    DefaultVotePeriod,
		VoteThreshold: DefaultVoteThreshold,
		VoteFrequency: DefaultVoteFrequency,
		OracleEnabled: DefaultOracleEnabled,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVoteFrequency, &p.VoteFrequency, validateVoteFrequency),
		paramtypes.NewParamSetPair(KeyPrevotePeriod, &p.PrevotePeriod, validatePrevotePeriod),
		paramtypes.NewParamSetPair(KeyVotePeriod, &p.VotePeriod, validateVotePeriod),
		paramtypes.NewParamSetPair(KeyVoteThreshold, &p.VoteThreshold, validateVoteThreshold),
		paramtypes.NewParamSetPair(KeyOracleEnabled, &p.OracleEnabled, validateOracleEnabled),
	}
}

// ValidateBasic performs basic validation on oracle parameters.
func (p Params) Validate() error {
	if p.VoteFrequency < 1 {
		return fmt.Errorf("invalid vote frequency: %d", p.VoteFrequency)
	}
	if p.VotePeriod <= 0 {
		return fmt.Errorf("vote period must be greater than 0: %d", p.VotePeriod)
	}
	if p.PrevotePeriod <= 0 {
		return fmt.Errorf("prevote period must be greater than 0: %d", p.PrevotePeriod)
	}
	if p.VoteThreshold.LTE(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("oracle parameter VoteTheshold must be greater than 33 percent")
	}
	if p.VoteThreshold.GT(sdk.OneDec()) {
		return fmt.Errorf("vote threshold too large: %s", p.VoteThreshold)
	}
	return nil
}

func validatePrevotePeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 1 {
		return fmt.Errorf("invalid Prevote period: %d", v)
	}
	return nil
}

func validateVotePeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 1 {
		return fmt.Errorf("invalid Vote period: %d", v)
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

func validateVoteThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LTE(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("vote threshold should be greater than 33%%: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("vote threshold must be below 100%%: %s", v)
	}
	return nil
}

func validateOracleEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

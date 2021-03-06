package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params,
	delegations []VoterDelegation,
	voteRounds []VoteRound,
	prevoteRounds []PrevoteRound,
	claimTypes []ClaimType,
) *GenesisState {

	return &GenesisState{
		Params:           params,
		VoterDelegations: delegations,
		Votes:            voteRounds,
		Prevotes:         prevoteRounds,
		ClaimTypes:       claimTypes,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		VoterDelegations: []VoterDelegation{},
		Votes:            []VoteRound{},
		Prevotes:         []PrevoteRound{},
		ClaimTypes:       []ClaimType{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	err := gs.Params.Validate()
	if err != nil {
		return err
	}
	// Validate stored claims in VoteRound
	for _, voteRound := range gs.Votes {
		for _, vote := range voteRound.Votes {
			_, err := vote.GetClaim()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetGenesisStateFromAppState returns x/oracle GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs GenesisState) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, round := range gs.Votes {
		err := round.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

package oracle

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/keeper"
	"github.com/encichain/enci/x/oracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}
	// Set params
	k.SetParams(ctx, genState.Params)

	// Set voter delegations to store
	for _, d := range genState.VoterDelegations {
		valAddr, err := sdk.ValAddressFromBech32(d.ValidatorAddress)
		if err != nil {
			panic(err)
		}
		delAddr, err := sdk.AccAddressFromBech32(d.DelegateAddress)
		if err != nil {
			panic(err)
		}
		k.SetVoterDelegation(ctx, delAddr, valAddr)
	}
	// Set Prevote round and prevotes
	for _, pRound := range genState.Prevotes {
		k.SetPrevoteRound(ctx, pRound)
		for _, prevote := range pRound.Prevotes {
			err := k.SetPrevote(ctx, prevote, pRound.ClaimType)
			if err != nil {
				panic(err)
			}
		}
	}
	// Set Vote round and votes
	for _, vRound := range genState.Votes {
		k.SetVoteRound(ctx, vRound)
		for _, vote := range vRound.Votes {
			valAddr, err := sdk.ValAddressFromBech32(vote.Validator)
			if err != nil {
				panic(err)
			}
			k.SetVote(ctx, valAddr, vote, vRound.ClaimType)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	delegations := k.GetAllVoterDelegations(ctx)
	voteRounds := k.GetAllVoteRounds(ctx)
	prevoteRounds := k.GetAllPrevoteRounds(ctx)

	return types.NewGenesisState(
		params,
		delegations,
		voteRounds,
		prevoteRounds,
	)
}

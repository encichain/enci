package types_test

import (
	fmt "fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	gs := types.DefaultGenesis()
	require.NotNil(t, gs.Votes)
	require.Len(t, gs.Votes, 0)

	require.NotNil(t, gs.Prevotes)
	require.Len(t, gs.Prevotes, 0)

	require.NotNil(t, gs.Params)
	require.Equal(t, gs.Params, types.DefaultParams())

	require.NotNil(t, gs.VoterDelegations)
	require.Len(t, gs.VoterDelegations, 0)
}

func TestNewGenesisState(t *testing.T) {
	var (
		voteRounds    []types.VoteRound
		delegations   []types.VoterDelegation
		prevoteRounds []types.PrevoteRound
	)
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"can proto marshal",
			func() {
				claim := types.TestClaim{}
				claimHash := claim.Hash()
				claimAny, _ := codectypes.NewAnyWithValue(&claim)
				opAddrStr := "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee"
				voteRounds = []types.VoteRound{{
					ClaimType: "test",
					Votes: []types.Vote{{
						Claim:     claimAny,
						Validator: opAddrStr,
						VotePower: 100,
					},
					},
					AggregatePower: 100,
				}}
				opAddr, _ := sdk.ValAddressFromBech32(opAddrStr)
				prevoteRounds = []types.PrevoteRound{{
					ClaimType: "test",
					Prevotes: []types.Prevote{{
						Hash:        types.CreateVoteHash("123", claimHash.String(), opAddr).String(),
						Validator:   opAddrStr,
						SubmitBlock: 101,
					}},
				}}
				delegations = []types.VoterDelegation{{
					DelegateAddress:  "enci1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qc84hfe8",
					ValidatorAddress: opAddrStr,
				}}
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {
			tc.malleate()

			if tc.expPass {
				require.NotPanics(t, func() {
					types.NewGenesisState(
						types.DefaultParams(),
						delegations,
						voteRounds,
						prevoteRounds,
					)
				})
			} else {
				require.Panics(t, func() {
					types.NewGenesisState(
						types.DefaultParams(),
						delegations,
						voteRounds,
						prevoteRounds,
					)
				})
			}
		})
	}
}

func TestGenesisStateValidate(t *testing.T) {
	var (
		genesisState  *types.GenesisState
		voteRounds    []types.VoteRound
		delegations   []types.VoterDelegation
		prevoteRounds []types.PrevoteRound
	)
	//params := types.DefaultParams()

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"valid",
			func() {
				genesisState = types.DefaultGenesis()
			},
			true,
		},
		{
			"invalid",
			func() {
				invParams := types.Params{
					PrevotePeriod: 0,
					VotePeriod:    1,
					VoteThreshold: types.DefaultVoteThreshold,
					VoteFrequency: 0,
				}
				genesisState = types.NewGenesisState(
					invParams, delegations, voteRounds, prevoteRounds,
				)
			},
			false,
		},
		{
			"expected claim",
			func() {
				genesisState = &types.GenesisState{
					Params:           types.DefaultParams(),
					VoterDelegations: delegations,
					Votes: []types.VoteRound{{
						Votes: []types.Vote{{
							Claim:     &codectypes.Any{},
							Validator: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee",
							VotePower: 100,
						}},
					}},
					Prevotes: prevoteRounds,
				}
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {
			tc.malleate()

			if tc.expPass {
				require.NoError(t, genesisState.Validate())
			} else {
				require.Error(t, genesisState.Validate())
			}
		})
	}
}

func TestUnpackInterfaces(t *testing.T) {
	claim := types.TestClaim{}
	claimAny, _ := codectypes.NewAnyWithValue(&claim)
	gs := types.GenesisState{
		Params:           types.DefaultParams(),
		VoterDelegations: []types.VoterDelegation{},
		Votes: []types.VoteRound{{
			ClaimType: "test",
			Votes: []types.Vote{{
				Claim:     claimAny,
				Validator: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee",
				VotePower: 100,
			}},
			AggregatePower: 100,
		}},
	}

	testCases := []struct {
		msg      string
		unpacker codectypes.AnyUnpacker
		expPass  bool
	}{
		{
			"success",
			codectypes.NewInterfaceRegistry(),
			true,
		},
		{
			"error",
			codec.NewLegacyAmino(),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.msg), func(t *testing.T) {

			if tc.expPass {
				require.NoError(t, gs.UnpackInterfaces(tc.unpacker))
				_, err := gs.Votes[0].Votes[0].GetClaim()
				require.NoError(t, err)

			} else {
				require.Error(t, gs.UnpackInterfaces(tc.unpacker))
			}
		})
	}
}

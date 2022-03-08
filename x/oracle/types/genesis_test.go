package types_test

import (
	fmt "fmt"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/encichain/enci/x/oracle/exported"
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
				claimAny, _ := codectypes.NewAnyWithValue(&types.TestClaim{})
				voteRounds = []types.VoteRound{
					{ClaimType: "test",
						Votes: []types.Vote{
							{
								Claim:     claimAny,
								Validator: "testvalidator",
								VotePower: 100,
							},
						},
						AggregatePower: 100,
					}}
				rounds = []types.Round{}
				pending = map[string][]uint64{
					"test": {1},
				}
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
						rounds,
						claims,
						pending,
						delegations,
						prevotes,
						finalizedRounds,
					)
				})
			} else {
				require.Panics(t, func() {
					types.NewGenesisState(
						types.DefaultParams(),
						rounds,
						claims,
						pending,
						delegations,
						prevotes,
						finalizedRounds,
					)
				})
			}
		})
	}
}

func TestGenesisStateValidate(t *testing.T) {
	var (
		genesisState    *types.GenesisState
		testClaim       []exported.Claim
		pending         map[string]([]uint64)
		delegations     []types.MsgDelegate
		prevotes        [][]byte
		finalizedRounds map[string](uint64)
	)
	round := []types.Round{}
	params := types.DefaultParams()

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"valid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i + 1),
						Content:     "test",
						ClaimType:   "test",
					}
				}
				genesisState = types.NewGenesisState(
					params, round, testClaim, pending, delegations, prevotes, finalizedRounds,
				)
			},
			true,
		},
		{
			"invalid",
			func() {
				testClaim = make([]exported.Claim, 100)
				for i := 0; i < 100; i++ {
					testClaim[i] = &types.TestClaim{
						BlockHeight: int64(i),
						Content:     "test",
						ClaimType:   "test",
					}
				}
				genesisState = types.NewGenesisState(
					params, round, testClaim, pending, delegations, prevotes, finalizedRounds,
				)
			},
			false,
		},
		{
			"expected claim",
			func() {
				genesisState = &types.GenesisState{
					Claims: []*codectypes.Any{{}},
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

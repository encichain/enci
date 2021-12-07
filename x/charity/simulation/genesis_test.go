package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity/simulation"
	"github.com/user/encichain/x/charity/types"
)

func TestGenCharities(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	charities := simulation.GenCharities(r)
	require.NotEmpty(t, charities)
}

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abonormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)
	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: 1000000,
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var charityGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &charityGenesis)

	require.Len(t, charityGenesis.Params.Charities, 4)
	require.Equal(t, sdk.NewDecWithPrec(3, 2), charityGenesis.Params.TaxRate)
	require.Equal(t, []types.TaxCap{{Denom: coretypes.MicroTokenDenom, Cap: sdk.NewInt(int64(3361739))}}, charityGenesis.Params.TaxCaps)
	require.Equal(t, sdk.NewDecWithPrec(36, 2), charityGenesis.Params.BurnRate)
}

// TestAbRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
func TestAbRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	// all these tests will panic
	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			}, "assignment to entry in nil map"},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}

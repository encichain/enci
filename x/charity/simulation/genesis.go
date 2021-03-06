package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity/keeper"
	"github.com/encichain/enci/x/charity/types"
)

// Simulation parameter constants
const (
	charitiesKey = "charities"
	taxCapsKey   = "tax_caps"
	taxRateKey   = "tax_rate"
	burnRateKey  = "burn_rate"
)

// GenCharities randomized Charities
func GenCharities(r *rand.Rand) []types.Charity {
	charities := []types.Charity{}
	addrs := createTestAddresses(r)
	amt := r.Intn(len(addrs)) + 1
	for i := 0; i < amt; i++ {
		cName := randomString(r, r.Intn(10))
		addr := addrs[i]

		ch := types.Charity{
			CharityName: cName,
			AccAddress:  addr,
			Checksum:    keeper.CreateCharitySha256(cName, addr),
		}

		charities = append(charities, ch)
	}
	return charities
}

//GenTaxCaps randomized TaxCaps
func GenTaxCaps(r *rand.Rand) []types.TaxCap {
	return []types.TaxCap{
		{Denom: coretypes.MicroTokenDenom, Cap: sdk.NewInt(int64(r.Int63() % 10000000))},
		{Denom: sdk.DefaultBondDenom, Cap: sdk.NewInt(int64(r.Intn(2)))},
	}
}

//GenTaxRate randomized TaxRate
func GenTaxRate(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(40)+1), 3)
}

//GenBurnRate randomized BurnRate
func GenBurnRate(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(49)+1), 2)
}

// RandomizedGenState generates a random GenesisState for charity module
func RandomizedGenState(simState *module.SimulationState) {

	var charities []types.Charity
	simState.AppParams.GetOrGenerate(
		simState.Cdc, charitiesKey, &charities, simState.Rand,
		func(r *rand.Rand) { charities = GenCharities(r) },
	)

	var taxRate sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, taxRateKey, &taxRate, simState.Rand,
		func(r *rand.Rand) { taxRate = GenTaxRate(r) },
	)

	var taxCaps []types.TaxCap
	simState.AppParams.GetOrGenerate(
		simState.Cdc, taxCapsKey, &taxCaps, simState.Rand,
		func(r *rand.Rand) { taxCaps = GenTaxCaps(r) },
	)

	var burnRate sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, burnRateKey, &burnRate, simState.Rand,
		func(r *rand.Rand) { burnRate = GenBurnRate(r) },
	)

	charityGenesis := types.NewGenesisState(
		types.Params{
			Charities: charities,
			TaxRate:   taxRate,
			TaxCaps:   taxCaps,
			BurnRate:  burnRate,
		},
		types.DefaultTaxRateLimits,
		taxCaps,
		sdk.Coins{},
		[]types.CollectionEpoch{},
	)

	bz, err := json.MarshalIndent(&charityGenesis, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Selected randomly generated charity parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(charityGenesis)
}

package testoracle

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	enciapp "github.com/encichain/enci/app"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/oracle/keeper"
	"github.com/encichain/enci/x/oracle/types"
)

// AddClaimType Registers claimType as an orcale params
func AddClaimType(ctx sdk.Context, k keeper.Keeper, claimType string) {
	params := types.DefaultParams()
	params.ClaimParams = map[string](types.ClaimParams){
		claimType: {
			ClaimType: claimType,
		},
	}
	k.SetParams(ctx, params)
}

// CreateTestInput Returns a simapp with custom OracleKeeper
// to avoid messing with the hooks.
func CreateTestInput() (*enciapp.EnciApp, sdk.Context) {
	return enciapp.CreateTestInput()
}

// CreateValidators intializes validators
func CreateValidators(t *testing.T, ctx sdk.Context, app *enciapp.EnciApp, powers []int64) ([]sdk.AccAddress, []sdk.ValAddress, []stakingtypes.ValidatorI) {
	addrs := enciapp.AddTestAddrsIncremental(app, ctx, 5, sdk.NewInt(30000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	pks := simapp.CreateTestPubKeys(5)

	stakingHelper := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	stakingHelper.Denom = coretypes.MicroTokenDenom

	appCodec := app.AppCodec()

	app.StakingKeeper = stakingkeeper.NewKeeper(
		appCodec,
		app.GetKey(stakingtypes.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)

	var vals []stakingtypes.ValidatorI
	for i, power := range powers {
		stakingHelper.CreateValidatorWithValPower(valAddrs[i], pks[i], power, true)
		val := app.StakingKeeper.Validator(ctx, valAddrs[i])
		vals = append(vals, val)
	}

	_ = staking.EndBlocker(ctx, app.StakingKeeper)
	return addrs, valAddrs, vals
}

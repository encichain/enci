package keeper_test

import "github.com/cosmos/cosmos-sdk/x/staking/types"

func (suite *KeeperTestSuite) TestGetSetDelegation() {
	ctx, k := suite.ctx, suite.app.OracleKeeper
	require := suite.Require()

	accAddrs := suite.addrs
	for i, val := range suite.validators {
		k.SetVoterDelegation(ctx, accAddrs[i], val)
		delegate, err := k.GetVoterDelegate(ctx, val)
		require.NoError(err)
		require.Equal(accAddrs[i], delegate)

		delegator, err := k.GetVoterDelegator(ctx, delegate)
		require.NoError(err)
		require.Equal(val, delegator)
	}

	delegations := k.GetAllVoterDelegations(ctx)
	require.Len(delegations, len(suite.validators))
}

func (suite *KeeperTestSuite) TestClearInactiveDelegations() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()
	accAddrs := suite.addrs
	validators := suite.validators

	for i, val := range validators {
		k.SetVoterDelegation(ctx, accAddrs[i], val)
		delegate, err := k.GetVoterDelegate(ctx, val)
		require.NoError(err)
		require.Equal(accAddrs[i], delegate)
	}
	delegations := k.GetAllVoterDelegations(ctx)
	require.Len(delegations, len(suite.validators))

	// alter validator bond status
	val, found := k.StakingKeeper.GetValidator(ctx, validators[0])
	require.True(found)
	val2, found := k.StakingKeeper.GetValidator(ctx, validators[1])
	require.True(found)

	val.Status = types.BondStatus(1)
	// Unbonding, should not be removed
	val2.Status = types.BondStatus(2)

	// Clear inactive delegations
	k.ClearInactiveDelegations(ctx)

	delegations = k.GetAllVoterDelegations(ctx)
	require.Len(delegations, len(suite.validators)-1)
}

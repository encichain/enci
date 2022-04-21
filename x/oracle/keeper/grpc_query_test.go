package keeper_test

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/keeper"
	"github.com/encichain/enci/x/oracle/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func (suite *KeeperTestSuite) TestQueryParam() {
	var (
		req *types.QueryParamsRequest
	)

	testCases := []struct {
		msg       string
		malleate  func()
		expPass   bool
		posttests func(res *types.QueryParamsResponse)
	}{
		{
			"success",
			func() {
				req = &types.QueryParamsRequest{}
			},
			true,
			func(res *types.QueryParamsResponse) {
				suite.Require().NotNil(res)
				suite.Require().Equal(res.Params, types.DefaultParams())
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest()

			tc.malleate()
			ctx := sdk.WrapSDKContext(suite.ctx)

			res, err := suite.queryClient.Params(ctx, req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(res)
			}

			tc.posttests(res)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryVoteRounds() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient
	require := suite.Require()
	// Register claim type
	app.OracleKeeper.RegisterClaimType(ctx, "test")
	// No Votes or Voteround set
	res, err := queryClient.VoteRounds(sdk.WrapSDKContext(ctx), &types.QueryVoteRoundsRequest{})
	require.NoError(err)
	require.NotNil(res)

	// Set Vote
	claim := types.TestClaim{}
	claimAny, err := codectypes.NewAnyWithValue(&claim)
	require.NoError(err)
	vote := types.Vote{
		Claim:     claimAny,
		Validator: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee",
		VotePower: 100,
	}
	valAddr, _ := sdk.ValAddressFromBech32(vote.Validator)
	app.OracleKeeper.SetVote(ctx, valAddr, vote, "test")
	res, err = queryClient.VoteRounds(sdk.WrapSDKContext(ctx), &types.QueryVoteRoundsRequest{})

	require.NoError(err)
	require.Len(res.VoteRounds, 1)
}

func (suite *KeeperTestSuite) TestQueryPrevoteRounds() {
	k, ctx, queryClient := suite.app.OracleKeeper, suite.ctx, suite.queryClient
	require := suite.Require()
	// Register claim type
	k.RegisterClaimType(ctx, "test")

	// No prevoteRound nor prevote set
	res, err := queryClient.PrevoteRounds(sdk.WrapSDKContext(ctx), &types.QueryPrevoteRoundsRequest{})
	require.NoError(err)
	require.NotNil(res)

	// Set prevote
	claim := types.TestClaim{ClaimType: "test"}
	valAddr := suite.validators[0]
	voteHash := types.CreateVoteHash("1", claim.Hash().String(), valAddr)
	prevote := types.NewPrevote(voteHash, valAddr, 1)
	k.SetPrevote(ctx, prevote, claim.Type())

	// Get PrevoteRound
	res, _ = queryClient.PrevoteRounds(sdk.WrapSDKContext(ctx), &types.QueryPrevoteRoundsRequest{})
	require.Len(res.PrevoteRounds, 1)
}

func (suite *KeeperTestSuite) TestQueryVoterDelegations() {
	k, ctx, queryClient, require := suite.app.OracleKeeper, suite.ctx, suite.queryClient, suite.Require()
	// no active delegations
	res, _ := queryClient.VoterDelegations(sdk.WrapSDKContext(ctx), &types.QueryVoterDelegationsRequest{})
	require.Len(res.VoterDelegations, 0)

	for i, valAddr := range suite.validators {
		k.SetVoterDelegation(ctx, suite.addrs[i], valAddr)
	}

	delegations := k.GetAllVoterDelegations(ctx)
	require.Len(delegations, len(suite.validators))

	res, _ = queryClient.VoterDelegations(sdk.WrapSDKContext(ctx), &types.QueryVoterDelegationsRequest{})
	require.Len(res.VoterDelegations, len(delegations))
	//============================
	// Test Query Delegate address
	for i, valAddr := range suite.validators {
		res, err := queryClient.DelegateAddress(
			sdk.WrapSDKContext(ctx),
			&types.QueryDelegateAddressRequest{Validator: valAddr.String()})
		require.NoError(err)
		require.Equal(res.Delegate, suite.addrs[i].String())
	}
	// empty request
	_, err := queryClient.DelegateAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegateAddressRequest{})
	require.Error(err)
	// invalid address
	_, err = queryClient.DelegateAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegateAddressRequest{Validator: ""})
	require.Error(err)
	// no delegate found
	_, err = queryClient.DelegateAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegateAddressRequest{Validator: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee"})
	require.Error(err)

	//=================================
	// Test Query Delegator address
	for i, accAddr := range suite.addrs {
		res, err := queryClient.DelegatorAddress(
			sdk.WrapSDKContext(ctx),
			&types.QueryDelegatorAddressRequest{Delegate: accAddr.String()})
		require.NoError(err)
		require.Equal(res.Validator, suite.validators[i].String())
	}
	// Empty request
	_, err = queryClient.DelegatorAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegatorAddressRequest{})
	require.Error(err)
	//invalid address
	_, err = queryClient.DelegatorAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegatorAddressRequest{Delegate: ""})
	require.Error(err)
	// no delegator found
	_, err = queryClient.DelegatorAddress(
		sdk.WrapSDKContext(ctx),
		&types.QueryDelegatorAddressRequest{Delegate: "enci16ruw3nnsrt963y47y8m8h0g6p4pkyudvm5j3fc"})
	require.Error(err)
}

func (suite *KeeperTestSuite) TestQueryNextPeriod() {
	k, ctx, queryClient, require := suite.app.OracleKeeper, suite.ctx, suite.queryClient, suite.Require()
	k.SetParams(ctx, types.DefaultParams())
	params := k.GetParams(ctx)
	// Set blockheight to 0
	ctx = ctx.WithBlockHeight(0)

	// Check next prevote, should not be current prevote period
	prevoteRes, err := queryClient.NextPrevote(sdk.WrapSDKContext(ctx), &types.QueryNextPrevoteRequest{})
	require.NoError(err)
	require.Equal(params.VoteFrequency-1, prevoteRes.Block)

	// Check next vote period
	voteRes, _ := queryClient.NextVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryNextVotePeriodRequest{})
	require.Equal(int(params.VoteFrequency+params.PrevotePeriod-1), int(voteRes.Block))

	// We use querier instance as queryClient blockheight is not updated with new context
	querier := keeper.NewQuerier(k)
	// Set blockheight to above first prevote and vote period (17286)
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency + params.PrevotePeriod + params.VotePeriod))

	prevoteRes, _ = querier.NextPrevote(sdk.WrapSDKContext(ctx), &types.QueryNextPrevoteRequest{})
	require.Equal(int(params.VoteFrequency*2-1), int(prevoteRes.Block))

	voteRes, _ = querier.NextVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryNextVotePeriodRequest{})
	require.Equal(params.VoteFrequency*2+params.PrevotePeriod-1, voteRes.Block)

	//=========================
	// Test with new queryClient
	ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{Height: int64(params.VoteFrequency)})
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)
	queryClient = types.NewQueryClient(queryHelper)

	voteRes, err = queryClient.NextVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryNextVotePeriodRequest{})
	require.NoError(err)
	require.Equal(int(params.VoteFrequency*1+params.PrevotePeriod-1), int(voteRes.Block))

}

func (suite *KeeperTestSuite) TestQueryClaimTypes() {
	k, ctx, queryClient := suite.app.OracleKeeper, suite.ctx, suite.queryClient
	require := suite.Require()

	// set claim types
	for i := 0; i < 5; i++ {
		k.RegisterClaimType(ctx, "test"+strconv.Itoa(i))
	}

	res, err := queryClient.ClaimTypes(sdk.WrapSDKContext(ctx), &types.QueryClaimTypesRequest{})
	require.NoError(err)
	require.Len(res.ClaimTypes, 5)
}

func (suite *KeeperTestSuite) TestQueryIsVotePeriod() {
	k, ctx, queryClient := suite.app.OracleKeeper, suite.ctx, suite.queryClient
	require := suite.Require()
	querier := keeper.NewQuerier(k)
	params := k.GetParams(ctx)

	// queryClient ctx blockheight == 1 - PrevotePeriod = True
	prevoteRes, _ := queryClient.IsPrevotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsPrevotePeriodRequest{})
	require.True(prevoteRes.IsPrevotePeriod)

	voteRes, _ := queryClient.IsVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsVotePeriodRequest{})
	require.False(voteRes.IsVotePeriod)

	// genesis vote period
	ctx = ctx.WithBlockHeight(3)
	prevoteRes, _ = querier.IsPrevotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsPrevotePeriodRequest{})
	require.False(prevoteRes.IsPrevotePeriod)

	voteRes, _ = querier.IsVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsVotePeriodRequest{})
	require.True(voteRes.IsVotePeriod)

	// second prevote period after genesis
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency * 2))
	prevoteRes, _ = querier.IsPrevotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsPrevotePeriodRequest{})
	require.True(prevoteRes.IsPrevotePeriod)

	voteRes, _ = querier.IsVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsVotePeriodRequest{})
	require.False(voteRes.IsVotePeriod)

	// second Vote Period
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency*2 + params.PrevotePeriod))
	prevoteRes, _ = querier.IsPrevotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsPrevotePeriodRequest{})
	require.False(prevoteRes.IsPrevotePeriod)

	voteRes, _ = querier.IsVotePeriod(sdk.WrapSDKContext(ctx), &types.QueryIsVotePeriodRequest{})
	require.True(voteRes.IsVotePeriod)
}

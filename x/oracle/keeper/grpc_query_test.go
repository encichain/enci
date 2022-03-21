package keeper_test

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/types"
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
	// No VoteRound set
	res, err := queryClient.VoteRounds(sdk.WrapSDKContext(ctx), &types.QueryVoteRoundsRequest{})
	require.NoError(err)
	require.NotNil(res)

	// Set VoteRound
	claim := types.TestClaim{}
	claimAny, err := codectypes.NewAnyWithValue(&claim)
	require.NoError(err)
	vote := types.Vote{
		Claim:     claimAny,
		Validator: "encivaloper1y8t2xrx5n7tzs5wtszyfeyjdtcq7d3qcjgmeee",
		VotePower: 100,
	}
	app.OracleKeeper.AppendVoteToRound(ctx, vote, "test")
	res, err = queryClient.VoteRounds(sdk.WrapSDKContext(ctx), &types.QueryVoteRoundsRequest{})

	require.NoError(err)
	require.Len(res.VoteRounds, 1)
}

/*
func (suite *KeeperTestSuite) TestQueryPrevoteRounds() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient
	require := suite.Require()

	// No VoteRound set
	res, err := queryClient.VoteRounds(sdk.WrapSDKContext(ctx), &types.QueryVoteRoundsRequest{})
	require.NoError(err)
	require.NotNil(res)
}

*/

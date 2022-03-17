package keeper_test

import (
	"encoding/hex"
	"testing"

	"github.com/encichain/enci/app"
	"github.com/encichain/enci/x/oracle"
	"github.com/encichain/enci/x/oracle/exported"
	"github.com/encichain/enci/x/oracle/keeper"
	"github.com/encichain/enci/x/oracle/testoracle"
	"github.com/encichain/enci/x/oracle/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *app.EnciApp

	queryClient types.QueryClient
	querier     sdk.Querier

	validators []sdk.ValAddress
	pow        []int64
	k          keeper.Keeper
	handler    sdk.Handler
	addrs      []sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest() {
	checkTx := false
	app, ctx := app.CreateTestInput()
	// cdc := app.LegacyAmino()
	powers := []int64{10, 10, 10, 10}
	addrs, validators, _ := testoracle.CreateValidators(suite.T(), ctx, app, powers)

	suite.addrs = addrs
	suite.validators = validators
	suite.pow = powers
	suite.ctx = app.GetBaseApp().NewContext(checkTx, tmproto.Header{Height: 1})
	suite.k = app.OracleKeeper

	suite.app = app

	querier := keeper.Querier{Keeper: app.OracleKeeper}
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)

	suite.queryClient = types.NewQueryClient(queryHelper)
	suite.handler = oracle.NewHandler(app.OracleKeeper)

}

func (suite *KeeperTestSuite) populateClaims(ctx sdk.Context, numClaims int) []exported.Claim {
	claims := make([]exported.Claim, numClaims)
	for i := 0; i < numClaims; i++ {
		claims[i] = types.NewTestClaim(int64(i), "test", "test")
		//suite.k.CreateClaim(ctx, claims[i])
	}
	return claims
}

func (suite *KeeperTestSuite) TestIsVotePeriod() {
	app, ctx := suite.app, suite.ctx
	app.OracleKeeper.SetParams(ctx, types.DefaultParams())
	require := suite.Require()
	params := app.OracleKeeper.GetParams(ctx)
	// not vote period
	ctx = ctx.WithBlockHeight(0)
	isVote := app.OracleKeeper.IsVotePeriod(ctx)
	require.False(isVote)
	isPrevote := app.OracleKeeper.IsPrevotePeriod(ctx)
	require.True(isPrevote)

	//vote period
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod) + 1)
	isVote = app.OracleKeeper.IsVotePeriod(ctx)
	require.True(isVote)
	isPrevote = app.OracleKeeper.IsPrevotePeriod(ctx)
	require.False(isPrevote)

	// First vote period after genesis
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency + params.PrevotePeriod))
	isVote = app.OracleKeeper.IsVotePeriod(ctx)
	require.True(isVote)
	isPrevote = app.OracleKeeper.IsPrevotePeriod(ctx)
	require.False(isPrevote)
}

func (suite *KeeperTestSuite) TestIsPrevotePeriod() {
	app, ctx := suite.app, suite.ctx
	app.OracleKeeper.SetParams(ctx, types.DefaultParams())
	require := suite.Require()
	params := app.OracleKeeper.GetParams(ctx)

	// not prevote period
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod + 1))
	isPrevote := app.OracleKeeper.IsPrevotePeriod(ctx)
	require.False(isPrevote)
	isVote := app.OracleKeeper.IsVotePeriod(ctx)
	require.True(isVote)

	//prevote period
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod - 2))
	isPrevote = app.OracleKeeper.IsPrevotePeriod(ctx)
	require.True(isPrevote)
	isVote = app.OracleKeeper.IsVotePeriod(ctx)
	require.False(isVote)

	// First prevote period after genesis
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency + 1))
	isPrevote = app.OracleKeeper.IsPrevotePeriod(ctx)
	require.True(isPrevote)
	isVote = app.OracleKeeper.IsVotePeriod(ctx)
	require.False(isVote)
}

func (suite *KeeperTestSuite) TestIsVotePeriodEnd() {
	app, ctx := suite.app, suite.ctx
	app.OracleKeeper.SetParams(ctx, types.DefaultParams())
	require := suite.Require()
	params := app.OracleKeeper.GetParams(ctx)

	// Not vote period end
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod))
	isEnd := app.OracleKeeper.IsVotePeriodEnd(ctx)
	require.False(isEnd)

	// vote period end
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod + params.VotePeriod - 2))
	isEnd = app.OracleKeeper.IsVotePeriodEnd(ctx)
	require.True(isEnd)

	// first end of voteperiod after genesis
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency + params.PrevotePeriod + params.VotePeriod - 2))
	isEnd = app.OracleKeeper.IsVotePeriod(ctx)
	require.True(isEnd)
}

func (suite *KeeperTestSuite) TestPreviousPrevotePeriod() {
	app, ctx := suite.app, suite.ctx
	app.OracleKeeper.SetParams(ctx, types.DefaultParams())
	require := suite.Require()
	params := app.OracleKeeper.GetParams(ctx)

	// Last prevote period: genesis
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency - params.PrevotePeriod - params.VotePeriod))
	lastPrevote := app.OracleKeeper.PreviousPrevotePeriod(ctx)
	require.Equal(0, int(lastPrevote))

	// If in current prevote, should return beginning of current prevote
	ctx = ctx.WithBlockHeight(int64(params.PrevotePeriod - 1))
	lastPrevote = app.OracleKeeper.PreviousPrevotePeriod(ctx)
	require.Equal(0, int(lastPrevote))

	// Last prevote after genesis vote period
	ctx = ctx.WithBlockHeight(int64(params.VoteFrequency + params.PrevotePeriod + params.VotePeriod + 5))
	lastPrevote = app.OracleKeeper.PreviousPrevotePeriod(ctx)
	require.Equal(params.VoteFrequency, lastPrevote)
}

func (suite *KeeperTestSuite) TestGetSetDeleteVote() {
	app, ctx := suite.app, suite.ctx
	require := suite.Require()

	claim := types.NewTestClaim(0, "test", "test")
	validator := suite.validators[0]

	vote, err := types.NewVote(claim, validator, 100)
	require.NoError(err)
	valAddr, err := sdk.ValAddressFromBech32(vote.Validator)
	require.NoError(err)

	// Test set vote
	app.OracleKeeper.SetVote(ctx, valAddr, vote, "test")

	//Test get vote
	getVote := app.OracleKeeper.GetVote(ctx, valAddr, "test")
	require.Equal(vote, getVote)

	// Test delete vote
	app.OracleKeeper.DeleteVote(ctx, valAddr, "test")
	getVote = app.OracleKeeper.GetVote(ctx, valAddr, "test")
	require.Equal(types.Vote{}, getVote)
}

func (suite *KeeperTestSuite) TestGetAllVotes() {
	app, ctx := suite.app, suite.ctx
	require := suite.Require()

	claim := types.NewTestClaim(0, "test", "test")
	for _, val := range suite.validators {
		vote, err := types.NewVote(claim, val, 100)
		require.NoError(err)
		app.OracleKeeper.SetVote(ctx, val, vote, "test")
	}
	// Set vote for another claim type
	vote, err := types.NewVote(claim, suite.validators[0], 100)
	require.NoError(err)
	app.OracleKeeper.SetVote(ctx, suite.validators[0], vote, "secondtest")
	votes := app.OracleKeeper.GetAllVotes(ctx)
	require.Equal(len(suite.validators)+1, len(votes))

	// Test GetVotesByClaimType
	votes = app.OracleKeeper.GetVotesByClaimType(ctx, "test")
	require.Equal(len(suite.validators), len(votes))

	// Test DeleteAllVotes()
	app.OracleKeeper.DeleteAllVotes(ctx)

	for _, val := range suite.validators {
		vote := app.OracleKeeper.GetVote(ctx, val, "test")
		require.Equal(types.Vote{}, vote)
	}
	vote = app.OracleKeeper.GetVote(ctx, suite.validators[0], "secondtest")
	require.Equal(types.Vote{}, vote)
}

func (suite *KeeperTestSuite) TestGetSetDeletePrevote() {
	ctx, keeper := suite.ctx, suite.app.OracleKeeper
	require := suite.Require()
	validator := suite.validators[0]

	claim := types.NewTestClaim(0, "test", "test")
	claimHash := hex.EncodeToString(claim.Hash())
	voteHash := types.CreateVoteHash("", claimHash, validator)
	prevote := types.NewPrevote(voteHash, validator, 0)

	// set prevote and get
	err := keeper.SetPrevote(ctx, prevote, "test")
	require.NoError(err)
	getPrevote, err := keeper.GetPrevote(ctx, validator, "test")
	require.NoError(err)
	require.Equal(prevote, getPrevote)

	// delete prevote
	keeper.DeletePrevote(ctx, validator, "test")
	getPrevote, err = keeper.GetPrevote(ctx, validator, "test")
	require.Error(err)
	require.Equal(types.Prevote{}, getPrevote)
}

func (suite *KeeperTestSuite) TestGetAllPrevotes() {
	k, ctx := suite.app.OracleKeeper, suite.ctx
	require := suite.Require()

	claim := types.NewTestClaim(0, "test", "type")
	altClaim := types.NewTestClaim(0, "test", "secondtype")
	altClaimHash := hex.EncodeToString(altClaim.Hash())
	claimHash := hex.EncodeToString(claim.Hash())

	for _, val := range suite.validators {
		voteHash := types.CreateVoteHash("", claimHash, val)
		prevote := types.NewPrevote(voteHash, val, 0)
		err := k.SetPrevote(ctx, prevote, claim.Type())
		require.NoError(err)
	}
	// Get all prevotes
	prevotes := k.GetAllPrevotes(ctx)
	require.Equal(len(suite.validators), len(prevotes))

	// Set alternate claim type prevote to store
	altVoteHash := types.CreateVoteHash("", altClaimHash, suite.validators[0])
	altPrevote := types.NewPrevote(altVoteHash, suite.validators[0], 1)
	err := k.SetPrevote(ctx, altPrevote, altClaim.Type())
	require.NoError(err)

	prevotes = k.GetAllPrevotes(ctx)
	require.Equal(len(suite.validators)+1, len(prevotes))

	// Test GetPrevotesByClaimType
	prevotes = k.GetPrevotesByClaimType(ctx, claim.Type())
	require.Equal(len(suite.validators), len(prevotes))

	// Delete prevotes
	k.DeleteAllPrevotes(ctx)
	prevotes = k.GetAllPrevotes(ctx)
	require.Equal(0, len(prevotes))
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

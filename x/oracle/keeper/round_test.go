package keeper_test

import (
	"encoding/hex"

	"github.com/encichain/enci/x/oracle/types"
)

func (suite *KeeperTestSuite) createTestPrevotesAndVotes() (prevotes []types.Prevote, votes []types.Vote) {
	require := suite.Require()

	claim := types.NewTestClaim(0, "test", "test")
	claimHash := hex.EncodeToString(claim.Hash())

	for _, val := range suite.validators {
		voteHash := types.CreateVoteHash("", claimHash, val)
		prevote := types.NewPrevote(voteHash, val, 0)
		vote, err := types.NewVote(claim, val, 100)
		require.NoError(err)
		votes = append(votes, vote)
		prevotes = append(prevotes, prevote)
	}
	return
}

func (suite *KeeperTestSuite) TestGetSetVoteRound() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()

	_, votes := suite.createTestPrevotesAndVotes()
	voteRound := types.NewVoteRound("test", votes)
	// Set and Get VoteRound
	k.SetVoteRound(ctx, voteRound)
	getVoteRound := k.GetVoteRound(ctx, "test")
	require.Equal(voteRound, getVoteRound)
	require.Len(voteRound.Votes, len(suite.validators))
	require.Equal(voteRound.AggregatePower, uint64(100*len(suite.validators)))

	// Test append vote
	claim := types.NewTestClaim(0, "test", "test")
	newVote, err := types.NewVote(claim, suite.validators[0], 100)
	require.NoError(err)
	k.AppendVoteToRound(ctx, newVote, "test")

	getVoteRound = k.GetVoteRound(ctx, "test")
	require.Len(getVoteRound.Votes, len(suite.validators)+1)
	require.Equal(getVoteRound.AggregatePower, uint64(100*(len(suite.validators)+1)))

	// Clear vote rounds from store
	k.ClearVoteRounds(ctx)
	getVoteRound = k.GetVoteRound(ctx, "test")
	require.Len(getVoteRound.Votes, 0)
}

func (suite *KeeperTestSuite) TestGetAllVoteRound() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()

	voteRounds := k.GetAllVoteRounds(ctx)
	require.Len(voteRounds, 0)
	// Set VoteRounds
	_, votes := suite.createTestPrevotesAndVotes()
	voteRound := types.NewVoteRound("test", votes)
	voteRound2 := types.NewVoteRound("secondtype", []types.Vote{})
	k.SetVoteRound(ctx, voteRound)
	k.SetVoteRound(ctx, voteRound2)

	voteRounds = k.GetAllVoteRounds(ctx)
	require.Len(voteRounds, 2)
	require.Equal(voteRound, voteRounds[1])

	// Clear vote rounds from store
	k.ClearVoteRounds(ctx)
	voteRounds = k.GetAllVoteRounds(ctx)
	require.Len(voteRounds, 0)
}

func (suite *KeeperTestSuite) TestGetSetPrevoteRound() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()

	prevotes, _ := suite.createTestPrevotesAndVotes()
	prevoteRound := types.NewPrevoteRound("test", prevotes)

	// Set and Get PrevoteRound
	k.SetPrevoteRound(ctx, prevoteRound)
	getPrevoteRound := k.GetPrevoteRound(ctx, "test")
	require.Equal(prevoteRound, getPrevoteRound)
	require.Len(prevoteRound.Prevotes, len(suite.validators))

	// clear prevote rounds from store
	k.ClearPrevoteRounds(ctx)
	getPrevoteRound = k.GetPrevoteRound(ctx, "test")
	require.Len(getPrevoteRound.Prevotes, 0)
}

func (suite *KeeperTestSuite) TestGetAllPrevoteRound() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()

	prevoteRounds := k.GetAllPrevoteRounds(ctx)
	require.Len(prevoteRounds, 0)

	// Set prevoteRounds
	prevotes, _ := suite.createTestPrevotesAndVotes()
	prevoteRound := types.NewPrevoteRound("test", prevotes)
	prevoteRound2 := types.NewPrevoteRound("secondtype", []types.Prevote{})
	k.SetPrevoteRound(ctx, prevoteRound)
	k.SetPrevoteRound(ctx, prevoteRound2)

	prevoteRounds = k.GetAllPrevoteRounds(ctx)
	require.Len(prevoteRounds, 2)

	require.Equal(prevoteRound, prevoteRounds[1])

	// Clear prevote rounds from store
	k.ClearPrevoteRounds(ctx)
	prevoteRounds = k.GetAllPrevoteRounds(ctx)
	require.Len(prevoteRounds, 0)
}

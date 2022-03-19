package keeper_test

import (
	"encoding/hex"

	"github.com/encichain/enci/x/oracle/exported"
	"github.com/encichain/enci/x/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func newTestMsgVote(r *require.Assertions, claim exported.Claim, sender sdk.AccAddress) exported.MsgVoteI {
	msg, err := types.NewMsgVote(sender, claim, "")
	r.NoError(err)
	return msg
}

func newTestMsgPrevote(r *require.Assertions, claim exported.Claim, sender sdk.AccAddress, validator sdk.ValAddress) *types.MsgPrevote {
	hash := types.CreateVoteHash("", hex.EncodeToString(claim.Hash()), validator)
	msg := types.NewMsgPrevote(claim.Type(), sender, hash)
	r.NotNil(msg)
	return msg
}

func newTestMsgDelegate(r *require.Assertions, validator sdk.ValAddress, delegate sdk.AccAddress) *types.MsgDelegate {
	msg := types.NewMsgDelegate(validator, delegate)
	r.NotNil(msg)
	return msg
}

func (suite *KeeperTestSuite) TestMsgVote() {
	nonDelegate, _ := sdk.AccAddressFromBech32("enci12p43x8vny2xqpl2z6n9rx2x2gna332ee5xkc9c")
	validator := suite.validators[0]
	val1 := suite.validators[1]
	delegate := suite.addrs[4]
	suite.app.OracleKeeper.SetParams(suite.ctx, types.DefaultParams())
	suite.k.SetVoterDelegation(suite.ctx, delegate, validator)
	params := suite.app.OracleKeeper.GetParams(suite.ctx)
	testCases := []struct {
		description string
		msg         sdk.Msg
		malleate    func()
		expectPass  bool
	}{
		{
			"valid - submitted by validator",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(4, "test", "test"),
				sdk.AccAddress(validator),
			),
			func() {
				suite.ctx = suite.ctx.WithBlockHeight(int64(params.PrevotePeriod))
				claim := types.NewTestClaim(4, "test", "test")
				claimHash := claim.Hash()
				voteHash := types.CreateVoteHash("", hex.EncodeToString(claimHash), validator)
				prevote := types.NewPrevote(voteHash, validator, 0)
				err := suite.k.SetPrevote(suite.ctx, prevote, claim.Type())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"valid - submitted by stored delegate",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(5, "test", "test"),
				sdk.AccAddress(delegate),
			),
			func() {
				claim := types.NewTestClaim(5, "test", "test")
				claimHash := claim.Hash()
				voteHash := types.CreateVoteHash("", hex.EncodeToString(claimHash), validator)
				prevote := types.NewPrevote(voteHash, validator, 0)
				err := suite.k.SetPrevote(suite.ctx, prevote, claim.Type())
				suite.Require().NoError(err)
			},
			true,
		},
		{
			"invalid - submitted by non-delegate",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(5, "test", "test"),
				nonDelegate,
			),
			func() {
				claim := types.NewTestClaim(5, "test", "test")
				claimHash := claim.Hash()
				voteHash := types.CreateVoteHash("", hex.EncodeToString(claimHash), validator)
				prevote := types.NewPrevote(voteHash, validator, 0)
				_ = suite.k.SetPrevote(suite.ctx, prevote, claim.Type())
			},
			false,
		},
		{
			"invalid - no prevote",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(5, "test", "test"),
				sdk.AccAddress(val1),
			),
			func() {
			},
			false,
		},
		{
			"invalid - hash mismatch",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(5, "test", "test"),
				sdk.AccAddress(delegate),
			),
			func() {
				claim := types.NewTestClaim(5, "test", "test")
				claimHash := claim.Hash()
				voteHash := types.CreateVoteHash("1", hex.EncodeToString(claimHash), validator)
				prevote := types.NewPrevote(voteHash, validator, 0)
				_ = suite.k.SetPrevote(suite.ctx, prevote, claim.Type())
			},
			false,
		},
		{
			"invalid - not vote period",
			newTestMsgVote(
				suite.Require(),
				types.NewTestClaim(70, "test", "test"),
				sdk.AccAddress(delegate),
			),
			func() {
				suite.ctx = suite.ctx.WithBlockHeight(70)
				claim := types.NewTestClaim(70, "test", "test")
				claimHash := claim.Hash()
				voteHash := types.CreateVoteHash("", hex.EncodeToString(claimHash), validator)
				prevote := types.NewPrevote(voteHash, validator, 0)
				_ = suite.k.SetPrevote(suite.ctx, prevote, claim.Type())
			},
			false,
		},
	}

	for i, tc := range testCases {
		tc.malleate()
		res, err := suite.handler(suite.ctx, tc.msg)
		if !tc.expectPass {
			suite.Require().Error(err, "expected error; tc #%d", i)
		} else {
			suite.Require().NoError(err, "unexpected error; tc #%d", i)
			suite.Require().NotNil(res, "expected non-nil result; tc #%d", i)

			vote := suite.app.OracleKeeper.GetVote(suite.ctx, validator, "test")
			suite.Require().Greater(vote.VotePower, uint64(0))

		}
	}
}

func (suite *KeeperTestSuite) TestMsgPrevote() {
	delegate := suite.addrs[4]
	nonDelegate, _ := sdk.AccAddressFromBech32("enci12p43x8vny2xqpl2z6n9rx2x2gna332ee5xkc9c")
	validator := suite.validators[0]
	suite.app.OracleKeeper.SetParams(suite.ctx, types.DefaultParams())
	suite.k.SetVoterDelegation(suite.ctx, delegate, validator)
	claim := types.NewTestClaim(5, "test", "test")
	testCases := []struct {
		description string
		msg         sdk.Msg
		malleate    func()
		expectPass  bool
	}{
		{
			"valid - submitted by registered delegate",
			newTestMsgPrevote(
				suite.Require(),
				claim,
				delegate,
				validator,
			),
			func() {
				suite.ctx = suite.ctx.WithBlockHeight(0)
			},
			true,
		},
		{
			"valid - submitted by validator",
			newTestMsgPrevote(
				suite.Require(),
				claim,
				sdk.AccAddress(validator),
				validator,
			),
			func() {},
			true,
		},
		{
			"invalid - submitted by non-delegate",
			newTestMsgPrevote(
				suite.Require(),
				claim,
				nonDelegate,
				validator,
			),
			func() {},
			false,
		},
		{
			"invalid - not prevote period",
			newTestMsgPrevote(
				suite.Require(),
				claim,
				delegate,
				validator,
			),
			func() {
				suite.ctx = suite.ctx.WithBlockHeight(3)
			},
			false,
		},
	}

	for i, tc := range testCases {
		tc.malleate()
		res, err := suite.handler(suite.ctx, tc.msg)

		if !tc.expectPass {
			suite.Require().Error(err, "expected error; tc #%d", i)
		} else {
			suite.Require().NoError(err, "unexpected error; tc #%d", i)
			suite.Require().NotNil(res, "expected non-nil result; tc #%d", i)
		}
	}
}

func (suite *KeeperTestSuite) TestMsgDelegate() {
	validator := suite.validators[0]
	delegate := suite.addrs[4]
	require := suite.Require()

	//valid
	msg := newTestMsgDelegate(suite.Require(), validator, delegate)
	res, err := suite.handler(suite.ctx, msg)
	require.NoError(err)
	require.NotNil(res)

	//invalid validator
	invalidVal, err := sdk.AccAddressFromBech32("enci12p43x8vny2xqpl2z6n9rx2x2gna332ee5xkc9c")
	require.NoError(err)
	msg = newTestMsgDelegate(suite.Require(), sdk.ValAddress(invalidVal), delegate)
	res, err = suite.handler(suite.ctx, msg)
	require.Error(err, "should have errored due to non-existent validator")
	require.Nil(res)

	//invalid address string
	msg = &types.MsgDelegate{
		Delegate:  "enci12p43x8vny2xqpl2z6n9rx2x2gna332ee5xkc",
		Validator: validator.String(),
	}
	res, err = suite.handler(suite.ctx, msg)
	require.Error(err)
	require.Nil(res)
}

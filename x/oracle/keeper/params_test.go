package keeper_test

import (
	"fmt"

	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/encichain/enci/x/oracle/types"
)

func (suite *KeeperTestSuite) TestParams() {
	ctx, k, require := suite.ctx, suite.app.OracleKeeper, suite.Require()
	defaultParams := types.DefaultParams()
	k.SetParams(ctx, types.DefaultParams())

	voteFrequency := k.GetVoteFrequency(ctx)
	require.Equal(defaultParams.VoteFrequency, voteFrequency)

	voteThreshold := k.GetVoteThreshold(ctx)
	require.True(defaultParams.VoteThreshold.Equal(voteThreshold))

	prevotePeriod := k.GetPrevotePeriod(ctx)
	require.Equal(defaultParams.PrevotePeriod, prevotePeriod)

	votePeriod := k.GetVotePeriod(ctx)
	require.Equal(defaultParams.VotePeriod, votePeriod)

	// Custom params
	for i := uint64(1); i < 10; i++ {
		params := types.Params{
			PrevotePeriod: i,
			VotePeriod:    i,
			VoteThreshold: sdk.NewDecWithPrec(55+int64(i), 2),
			VoteFrequency: 100 + i,
		}
		k.SetParams(ctx, params)
		getParams := k.GetParams(ctx)
		require.Equal(params, getParams)
	}
}

func (suite *KeeperTestSuite) TestParamChangeProposal() {
	app, ctx, k := suite.app, suite.ctx, suite.app.OracleKeeper
	require := suite.Require()

	k.SetParams(ctx, types.DefaultParams())
	params := k.GetParams(ctx)
	require.Equal(types.DefaultParams(), params)

	proposalfile := sdktestutil.WriteToNewTempFile(suite.T(), `
	{
		"title": "Oracle Param Change",
		"description": "Change all oracle params",
		"changes": [
		  {"subspace": "oracle", "key": "PrevotePeriod", "value": "5"},
		  {"subspace": "oracle", "key": "VotePeriod", "value": "5"},
		  {"subspace": "oracle", "key": "VoteThreshold", "value": "0.600000000000000000"},
		  {"subspace": "oracle", "key": "VoteFrequency", "value": "100"}
		],
		"deposit": "10000000stake"
	  }
	`)

	// Test parsing param change proposal file
	proposal, err := paramsutils.ParseParamChangeProposalJSON(suite.cdc, proposalfile.Name())
	require.NoError(err)
	require.Equal("Oracle Param Change", proposal.Title)
	require.Equal("Change all oracle params", proposal.Description)
	require.Equal("10000000stake", proposal.Deposit)
	require.Len(proposal.Changes, 4)

	// Create new ParameterChangeProposal{} from proposal
	content := paramproposal.NewParameterChangeProposal(
		proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
	)
	err = content.ValidateBasic()
	require.NoError(err)
	err = paramproposal.ValidateChanges(content.Changes)
	require.NoError(err)

	// call param change handler
	err = handleParameterChangeProposal(ctx, app.ParamsKeeper, content)
	require.NoError(err)

	expectedParams := types.Params{
		PrevotePeriod: 5,
		VotePeriod:    5,
		VoteThreshold: sdk.NewDecWithPrec(60, 2),
		VoteFrequency: 100,
	}

	params = k.GetParams(ctx)
	require.Equal(expectedParams, params)

}

//Taken from params proposal_handler.go. Attempts to update the params with the values in ParameterChangeProposal{}
func handleParameterChangeProposal(ctx sdk.Context, k paramskeeper.Keeper, p *paramproposal.ParameterChangeProposal) error {
	for _, c := range p.Changes {
		ss, ok := k.GetSubspace(c.Subspace)
		if !ok {
			return sdkerrors.Wrap(paramproposal.ErrUnknownSubspace, c.Subspace)
		}

		k.Logger(ctx).Info(
			fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
		)

		if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
			return sdkerrors.Wrapf(paramproposal.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
		}
	}

	return nil
}

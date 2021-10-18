package keeper

import (
	"encoding/json"
	"testing"

	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/stretchr/testify/require"

	"github.com/user/encichain/x/charity/types"
)

func TestParamsFuncs(t *testing.T) {
	app := CreateTestApp(t)
	testParams := types.Params{
		Charities: []types.Charity{
			{CharityName: "test", AccAddress: "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6test", Checksum: "8FF7B399E22B0C99B5FF1B8F0859858797ECB81609BD0088F14A53CC9B417185"}},
		TaxRate: sdk.NewDecWithPrec(1, 2),
		TaxCaps: []types.TaxCap{{
			Denom: "uenci",
			Cap:   sdk.NewInt(int64(2000000)),
		}},
	}
	defaultParams := types.DefaultParams()

	require.Equal(t, defaultParams, app.CharityKeeper.GetAllParams(app.Ctx))
	require.Equal(t, defaultParams.Charities, app.CharityKeeper.GetCharity(app.Ctx))
	require.Equal(t, defaultParams.TaxRate, app.CharityKeeper.GetTaxRate(app.Ctx))
	require.Equal(t, defaultParams.TaxCaps, app.CharityKeeper.GetParamTaxCaps(app.Ctx))

	//Try to set new params
	app.CharityKeeper.SetParams(app.Ctx, testParams)

	require.Equal(t, testParams, app.CharityKeeper.GetAllParams(app.Ctx))
	require.Equal(t, testParams.Charities, app.CharityKeeper.GetCharity(app.Ctx))
	require.Equal(t, testParams.TaxRate, app.CharityKeeper.GetTaxRate(app.Ctx))
	require.Equal(t, testParams.TaxCaps, app.CharityKeeper.GetParamTaxCaps(app.Ctx))

}

func TestCharityParamChangeProposal(t *testing.T) {
	app := CreateTestApp(t)

	proposalfile := sdktestutil.WriteToNewTempFile(t, `
	{
		"title": "Charity Param Change",
		"description": "Update charities",
		"changes": [
		  {
			"subspace": "charity",
			"key": "Charities",
			"value":
			  [{
				"charity_name" : "FOUNDATION OF THE NEEDY CHILDREN TEST",
				"accAddress" : "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6",
				"checksum" : "8FF7B399E22B0C99B5FF1B8F0859858797ECB81609BD0088F14A53CC9B417185"
			  }]
		  }
		],
		"deposit": "10000000stake"
	  }
	`)
	valuebytes := json.RawMessage{0x5b, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x9, 0x22, 0x63, 0x68, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f,
		0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x3a, 0x20, 0x22, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x54,
		0x49, 0x4f, 0x4e, 0x20, 0x4f, 0x46, 0x20, 0x54, 0x48, 0x45, 0x20, 0x4e, 0x45, 0x45, 0x44, 0x59,
		0x20, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x52, 0x45, 0x4e, 0x20, 0x54, 0x45, 0x53, 0x54, 0x22, 0x2c,
		0xa, 0x9, 0x9, 0x9, 0x9, 0x22, 0x61, 0x63, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
		0x20, 0x3a, 0x20, 0x22, 0x65, 0x6e, 0x63, 0x69, 0x31, 0x66, 0x74, 0x78, 0x61, 0x70, 0x72, 0x36,
		0x65, 0x63, 0x6e, 0x72, 0x6d, 0x78, 0x75, 0x6b, 0x70, 0x38, 0x32, 0x33, 0x36, 0x77, 0x79, 0x38, 0x73,
		0x65, 0x77, 0x6e, 0x6e, 0x32, 0x71, 0x35, 0x33, 0x30, 0x73, 0x70, 0x6a, 0x6e, 0x36, 0x22, 0x2c, 0xa, 0x9,
		0x9, 0x9, 0x9, 0x22, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x22, 0x20, 0x3a, 0x20, 0x22, 0x38, 0x46,
		0x46, 0x37, 0x42, 0x33, 0x39, 0x39, 0x45, 0x32, 0x32, 0x42, 0x30, 0x43, 0x39, 0x39, 0x42, 0x35, 0x46, 0x46, 0x31,
		0x42, 0x38, 0x46, 0x30, 0x38, 0x35, 0x39, 0x38, 0x35, 0x38, 0x37, 0x39, 0x37, 0x45, 0x43, 0x42, 0x38, 0x31, 0x36,
		0x30, 0x39, 0x42, 0x44, 0x30, 0x30, 0x38, 0x38, 0x46, 0x31, 0x34, 0x41, 0x35, 0x33, 0x43, 0x43, 0x39, 0x42, 0x34,
		0x31, 0x37, 0x31, 0x38, 0x35, 0x22, 0xa, 0x9, 0x9, 0x9, 0x20, 0x20, 0x7d, 0x5d}

	var charities []types.Charity
	err := json.Unmarshal(valuebytes, &charities)
	require.NoError(t, err)
	require.Equal(t, []types.Charity{
		{CharityName: "FOUNDATION OF THE NEEDY CHILDREN TEST",
			AccAddress: "enci1ftxapr6ecnrmxukp8236wy8sewnn2q530spjn6",
			Checksum:   "8FF7B399E22B0C99B5FF1B8F0859858797ECB81609BD0088F14A53CC9B417185"},
	}, charities)

	// Test parsing param change proposal file
	proposal, err := paramsutils.ParseParamChangeProposalJSON(app.Cdc, proposalfile.Name())

	require.NoError(t, err)
	require.Equal(t, "Charity Param Change", proposal.Title)
	require.Equal(t, "Update charities", proposal.Description)
	require.Equal(t, "10000000stake", proposal.Deposit)

	// Create new ParameterChangeProposal{} from proposal
	content := paramproposal.NewParameterChangeProposal(
		proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
	)

	err = content.ValidateBasic()
	require.NoError(t, err)
	err = paramproposal.ValidateChanges(proposal.Changes.ToParamChanges())
	require.NoError(t, err)
	require.NotEqual(t, paramproposal.ParameterChangeProposal{}, content)

	require.Equal(t, types.DefaultCharities, app.CharityKeeper.GetCharity(app.Ctx))
	// Attempt to update charities
	err = handleParameterChangeProposal(app.Ctx, app.ParamsKeeper, content)
	require.NoError(t, err)
	require.Equal(t, charities, app.CharityKeeper.GetCharity(app.Ctx))

}

func TestTaxRateParamChangeProposal(t *testing.T) {
	app := CreateTestApp(t)

	proposalfile := sdktestutil.WriteToNewTempFile(t, `
	{
		"title": "Charity TaxRate Change",
		"description": "Update taxrate to 1%",
		"changes": [
		  {
			"subspace": "charity",
			"key": "TaxRate",
			"value": "0.010000000000000000"
		  }
		],
		"deposit": "10000000stake"
	  }
	`)
	// Test parsing param change proposal file
	proposal, err := paramsutils.ParseParamChangeProposalJSON(app.Cdc, proposalfile.Name())
	require.NoError(t, err)
	require.Equal(t, "Charity TaxRate Change", proposal.Title)
	require.Equal(t, "Update taxrate to 1%", proposal.Description)
	require.Equal(t, "10000000stake", proposal.Deposit)

	// Create new ParameterChangeProposal{} from proposal
	content := paramproposal.NewParameterChangeProposal(
		proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
	)
	err = content.ValidateBasic()
	require.NoError(t, err)
	err = paramproposal.ValidateChanges(proposal.Changes.ToParamChanges())
	require.NoError(t, err)

	require.Equal(t, types.DefaultTaxRate, app.CharityKeeper.GetTaxRate(app.Ctx))

	// Attempt to update charities
	err = handleParameterChangeProposal(app.Ctx, app.ParamsKeeper, content)
	require.NoError(t, err)
	require.Equal(t, sdk.NewDecWithPrec(1, 2), app.CharityKeeper.GetTaxRate(app.Ctx))
}

func TestTaxCapsParamChangeProposal(t *testing.T) {
	app := CreateTestApp(t)

	proposalfile := sdktestutil.WriteToNewTempFile(t, `
	{
		"title": "Charity TaxCaps Change",
		"description": "Update uenci taxcap to 2enci",
		"changes": [
		  {
			"subspace": "charity",
			"key": "Taxcaps",
			"value": [{ 
				"denom" : "uenci",
				"Cap" :   "2000000"
			}]
		  }
		],
		"deposit": "10000000stake"
	  }
	`)
	// Test parsing param change proposal file
	proposal, err := paramsutils.ParseParamChangeProposalJSON(app.Cdc, proposalfile.Name())
	require.NoError(t, err)
	require.Equal(t, "Charity TaxCaps Change", proposal.Title)
	require.Equal(t, "Update uenci taxcap to 2enci", proposal.Description)
	require.Equal(t, "10000000stake", proposal.Deposit)

	// Create new ParameterChangeProposal{} from proposal
	content := paramproposal.NewParameterChangeProposal(
		proposal.Title, proposal.Description, proposal.Changes.ToParamChanges(),
	)
	err = content.ValidateBasic()
	require.NoError(t, err)
	err = paramproposal.ValidateChanges(proposal.Changes.ToParamChanges())
	require.NoError(t, err)

	require.Equal(t, types.DefaultTaxCaps, app.CharityKeeper.GetParamTaxCaps(app.Ctx))

	// Attempt to update charities
	err = handleParameterChangeProposal(app.Ctx, app.ParamsKeeper, content)
	require.NoError(t, err)
	require.Equal(t, []types.TaxCap{{Denom: "uenci", Cap: sdk.NewInt(int64(2000000))}}, app.CharityKeeper.GetParamTaxCaps(app.Ctx))
}

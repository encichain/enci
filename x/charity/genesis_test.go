package charity_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity"
	"github.com/encichain/enci/x/charity/keeper"
	"github.com/encichain/enci/x/charity/types"

	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	coreapp "github.com/encichain/enci/app"
)

type GenesisTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *coreapp.EnciApp
}

func (suite *GenesisTestSuite) TestExportGenesis() {
	suite.app = coreapp.Setup(false)
	app := suite.app
	//ctx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerEpoch) * 30)
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	ctx := suite.ctx.WithBlockHeight(int64(coretypes.BlocksPerEpoch) * 30)
	params := types.Params{
		Charities: []types.Charity{
			{CharityName: "foo", AccAddress: "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55", Checksum: keeper.CreateCharitySha256("foo", "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55")},
		},
		TaxRate:  sdk.NewDecWithPrec(5, 3),
		TaxCaps:  []types.TaxCap{{Denom: "bar", Cap: sdk.NewInt(5000000)}},
		BurnRate: sdk.NewDecWithPrec(10, 2),
	}
	taxCaps := []types.TaxCap{{Denom: "bar", Cap: sdk.NewInt(5000000)}}
	taxProceeds := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(2000000)))
	taxRateLimits := types.TaxRateLimits{
		RateMin:     sdk.ZeroDec(),
		TaxRateMax:  sdk.NewDecWithPrec(10, 2),
		BurnRateMax: sdk.NewDecWithPrec(40, 2),
	}
	collectionEpochs := []types.CollectionEpoch{}

	newGenesis := types.NewGenesisState(params, taxRateLimits, taxCaps, taxProceeds, collectionEpochs)
	charity.InitGenesis(ctx, app.CharityKeeper, *newGenesis)

	exportGenesis := charity.ExportGenesis(ctx, app.CharityKeeper)

	// create new app that does not share persistent or in-memory state
	// and initialize app from exported genesis state above.()
	newApp := coreapp.Setup(false)
	newCtx := newApp.BaseApp.NewContext(false, tmproto.Header{}).WithBlockHeight(int64(coretypes.BlocksPerEpoch) * 30)
	charity.InitGenesis(newCtx, newApp.CharityKeeper, *exportGenesis)
	secondExpGen := charity.ExportGenesis(newCtx, newApp.CharityKeeper)

	// Exported genesisState should be the same as origin
	suite.Require().Equal(*exportGenesis, *secondExpGen)

}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

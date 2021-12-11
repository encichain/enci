package charity_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"
	"github.com/user/encichain/x/charity"
	"github.com/user/encichain/x/charity/keeper"
	"github.com/user/encichain/x/charity/types"
	//bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	//coreapp "github.com/user/encichain/app"
)

func TestExportGenesis(t *testing.T) {
	app := keeper.CreateKeeperTestApp(t)
	ctx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerPeriod) * 30)
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
	collectionPeriods := []types.CollectionPeriod{}

	newGenesis := types.NewGenesisState(params, taxRateLimits, taxCaps, taxProceeds, collectionPeriods)
	charity.InitGenesis(ctx, app.CharityKeeper, *newGenesis)
	exportGenesis := charity.ExportGenesis(ctx, app.CharityKeeper)
	// Exported genesisState should be the same as origin
	require.Equal(t, newGenesis, exportGenesis)

}

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
)

func TestExportGenesis(t *testing.T) {
	app := keeper.CreateKeeperTestApp(t)
	ctx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerMonth))

	params := types.Params{
		Charities: []types.Charity{
			{CharityName: "foo", AccAddress: "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55", Checksum: keeper.CreateCharitySha256("foo", "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55")},
		},
		TaxRate:  sdk.NewDecWithPrec(5, 3),
		TaxCaps:  []types.TaxCap{{Denom: "bar", Cap: sdk.NewInt(5000000)}},
		BurnRate: sdk.NewDecWithPrec(10, 2),
	}
	app.CharityKeeper.SetParams(ctx, params)
	app.CharityKeeper.SetTaxCap(ctx, "bar", sdk.NewInt(5000000))
	app.CharityKeeper.SetTaxProceeds(ctx, sdk.Coins{})
	genesis := charity.ExportGenesis(ctx, app.CharityKeeper)

	newApp := keeper.CreateKeeperTestApp(t)
	newCtx := app.Ctx.WithBlockHeight(int64(coretypes.BlocksPerMonth))
	charity.InitGenesis(newCtx, newApp.CharityKeeper, *genesis)
	newGenesis := charity.ExportGenesis(newCtx, newApp.CharityKeeper)

	require.Equal(t, genesis, newGenesis)

}

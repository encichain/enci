package keeper_test

import (
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

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

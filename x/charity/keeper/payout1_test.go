package keeper_test

/*
import (
	"crypto/sha256"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	coretypes "github.com/user/encichain/types"

	"github.com/user/encichain/x/charity/types"
	//"go.uber.org/goleak"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/user/encichain/x/charity/keeper"
)

func (suite *CharityTestSuite) TestPayoutFunctions() {
	// Note: Goroutine leaks detected in App. Will cause abnormalities and failed tests in subsequent test functions if CreateKeeperTestApp is reinitialized.
	//defer goleak.VerifyNone(t)
	suite.SetupTest(false)
	app := suite.app
	ctx := suite.ctx
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")

	bech32addr := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	addr, err := sdk.AccAddressFromBech32(bech32addr)
	suite.Require().NoError(err)
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	suite.Require().NotNil(acc)

	// Create checksum and encode to string
	csb := sha256.Sum256([]byte("Test Charity" + bech32addr))
	checksum := hex.EncodeToString(csb[:])
	suite.Require().NotEqual("", checksum)

	params := types.DefaultParamsSet
	params.Charities = []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr,
			Checksum:   checksum,
		},
	}
	app.CharityKeeper.SetParams(ctx, params)
	suite.Require().Equal(app.CharityKeeper.GetAllParams(ctx), params)

	// Check if able to send coins to addr
	coins := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(1000))))
	// Fund charity tax collector
	modAcc := app.AccountKeeper.GetModuleAccount(ctx, types.CharityCollectorName)
	suite.Require().NotNil(modAcc)
	suite.app.AccountKeeper.SetAccount(ctx, modAcc)
	err = simapp.FundModuleAccount(app.BankKeeper, ctx, types.CharityCollectorName, coins)
	suite.Require().NoError(err)
	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.CharityCollectorName, addr, coins)
	suite.Require().NoError(err)

	// Test DonateCharity for a valid account
	charities := app.CharityKeeper.GetCharity(ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(ctx, coins, charity)
		suite.Require().NoError(err)
	}

	// Test non-existent account
	bech32addr = "enci1vvcw744ck9kzczrhf282lqmset47jnxe9090qt"
	suite.Require().NoError(err)
	csb = sha256.Sum256([]byte("Test Charity" + bech32addr))
	checksum = hex.EncodeToString(csb[:])
	params.Charities[0].AccAddress = bech32addr
	params.Charities[0].Checksum = checksum
	app.CharityKeeper.SetParams(ctx, params)

	charities = app.CharityKeeper.GetCharity(ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(ctx, coins, charity)
		suite.Require().Errorf(err, "account does not exist for the provided address")
	}

	// Test Donate with invalid Checksum
	bech32addr = "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	addrforinvchecksum := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv54"
	csb = sha256.Sum256([]byte("Test Charity" + addrforinvchecksum))
	checksum = hex.EncodeToString(csb[:])
	params.Charities = []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr,
			Checksum:   checksum,
		},
	}
	app.CharityKeeper.SetParams(ctx, params)
	suite.Require().Equal(app.CharityKeeper.GetAllParams(ctx), params)
	// Test DonateCharity with invalid checksum: expect error
	charities = app.CharityKeeper.GetCharity(ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(ctx, coins, charity)
		suite.Require().Errorf(err, "checksum is invalid")
	}

	//Test CalculateSplit and DisburseDonations
	// Configure test charity accounts
	bech32addr1 := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	bech32addr2 := "enci1vvcw744ck9kzczrhf282lqmset47jnxe9090qt"
	addr1, err := sdk.AccAddressFromBech32(bech32addr1)
	suite.Require().NoError(err)
	addr2, err := sdk.AccAddressFromBech32(bech32addr2)
	suite.Require().NoError(err)
	acc1 := app.AccountKeeper.NewAccountWithAddress(ctx, addr1)
	suite.Require().NotNil(acc1)
	acc2 := app.AccountKeeper.NewAccountWithAddress(ctx, addr2)
	suite.Require().NotNil(acc2)

	// Set accounts to store
	app.AccountKeeper.SetAccount(ctx, acc1)
	app.AccountKeeper.SetAccount(ctx, acc2)
	// Create checksums and encode to strings
	csb1 := sha256.Sum256([]byte("Test Charity" + bech32addr1))
	checksum1 := hex.EncodeToString(csb1[:])
	suite.Require().NotEqual("", checksum1)

	csb2 := sha256.Sum256([]byte("Test Charity 2" + bech32addr2))
	checksum2 := hex.EncodeToString(csb2[:])

	validCharities := []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr1,
			Checksum:   checksum1,
		},
		{CharityName: "Test Charity 2",
			AccAddress: bech32addr2,
			Checksum:   checksum2},
	}

	invalidCharities := []types.Charity{
		{CharityName: "Invalid charity1",
			//invalid accAddress
			AccAddress: bech32addr1 + "a",
			Checksum:   keeper.CreateCharitySha256("Invalid charity1", (bech32addr1 + "a")),
		},
		{CharityName: "Invalid charity2",
			AccAddress: bech32addr2,
			//invalid checksum
			Checksum: checksum2},
	}

	params.Charities = invalidCharities

	app.CharityKeeper.SetParams(ctx, params)
	suite.Require().Equal(app.CharityKeeper.GetAllParams(ctx), params)
	charities = app.CharityKeeper.GetCharity(ctx)
	suite.Require().Equal(params.Charities, charities)

	taxaddr := app.AccountKeeper.GetModuleAddress(types.CharityCollectorName)
	balance := app.BankKeeper.SpendableCoins(ctx, taxaddr)
	baseamt := app.BankKeeper.GetAllBalances(ctx, taxaddr)
	hasbalance := app.BankKeeper.HasBalance(ctx, taxaddr, sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(1000))))
	suite.Require().Equal(baseamt, balance)
	suite.Require().True(hasbalance)

	//Test calculate split
	split := app.CharityKeeper.CalculateSplit(ctx, app.CharityKeeper.GetCharity(ctx))
	suite.Require().Equal(baseamt[0].Amount.Quo(sdk.NewInt(int64(2))), split[0].Amount)

	//DisburseDonations with invalid charities should return errors
	payouts, errs := app.CharityKeeper.DisburseDonations(ctx, charities)
	suite.Require().NotEmpty(errs)
	suite.Require().Equal(true, len(errs) == 2)
	suite.Require().Empty(payouts)
	expectedErrs := []string{"Payout failed for charity: Invalid charity1, with error: invalid address", "Payout failed for charity: Invalid charity2, with error: checksum is invalid"}
	suite.Require().Equal(expectedErrs, errs)

	//Test DisburseDonations with valid charity
	params.Charities = validCharities
	app.CharityKeeper.SetParams(ctx, params)
	charities = app.CharityKeeper.GetCharity(ctx)
	suite.Require().Equal(params.Charities, charities)
	suite.Require().Equal(app.CharityKeeper.GetAllParams(ctx), params)

	payouts, errs = app.CharityKeeper.DisburseDonations(ctx, charities)
	suite.Require().Empty(errs)
	suite.Require().Equal(true, len(payouts) == 2)
	suite.Require().Equal([]types.Payout{{Recipientaddr: bech32addr1, Coins: split}, {Recipientaddr: bech32addr2, Coins: split}}, payouts)
}

func (suite *CharityTestSuite) TestBurnCoins() {
	suite.SetupTest(false)
	app := suite.app
	ctx := suite.ctx
	// Get burner account address and ensure it has no balance
	burnAddr := app.AccountKeeper.GetModuleAddress(types.BurnAccName)
	isZeroBal := app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero()
	suite.Require().True(isZeroBal)

	// Fund burner account
	coins := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(10000000))))
	err := simapp.FundModuleAccount(app.BankKeeper, ctx, types.BurnAccName, coins)
	suite.Require().NoError(err)
	hasBal := app.BankKeeper.HasBalance(ctx, burnAddr, sdk.NewCoin(coretypes.MicroTokenDenom, coins[0].Amount))
	suite.Require().True(hasBal)

	// Test BurnCoinsFromBurner
	err = app.CharityKeeper.BurnCoinsFromBurner(ctx)
	suite.Require().NoError(err)
	isZeroBal = app.BankKeeper.GetAllBalances(ctx, burnAddr).IsZero()
	suite.Require().False(isZeroBal)
}

*/

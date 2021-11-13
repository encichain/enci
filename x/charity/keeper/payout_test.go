package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"

	"github.com/user/encichain/x/charity/types"
	//"go.uber.org/goleak"
)

func TestPayoutFunctions(t *testing.T) {
	// Note: Goroutine leaks detected in App. Will cause abnormalities and failed tests in subsequent test functions if CreateTestApp is reinitialized.
	//defer goleak.VerifyNone(t)
	app := CreateTestApp(t)
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")

	bech32addr := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	addr, err := sdk.AccAddressFromBech32(bech32addr)
	require.NoError(t, err)
	acc := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr)
	require.NotNil(t, acc)

	// Create checksum and encode to string
	csb := sha256.Sum256([]byte("Test Charity" + bech32addr))
	checksum := hex.EncodeToString(csb[:])
	require.NotEqual(t, "", checksum)

	params := types.DefaultParamsSet
	params.Charities = []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr,
			Checksum:   checksum,
		},
	}
	app.CharityKeeper.SetParams(app.Ctx, params)
	require.Equal(t, app.CharityKeeper.GetAllParams(app.Ctx), params)

	// Check if able to send coins to addr
	coins := sdk.Coins{{Denom: coretypes.MicroTokenDenom, Amount: sdk.NewInt(int64(1000))}}
	err = app.BankKeeper.SendCoinsFromModuleToAccount(app.Ctx, types.CharityCollectorName, addr, coins)
	require.NoError(t, err)

	// Test DonateCharity for a valid account
	charities := app.CharityKeeper.GetCharity(app.Ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(app.Ctx, coins, charity)
		require.NoError(t, err)
	}
	// Reimburse sent amount
	err = FundModuleAccount(app.BankKeeper, app.Ctx, types.CharityCollectorName, coins)
	require.NoError(t, err)
	// Test non-existent account
	bech32addr = "enci1vvcw744ck9kzczrhf282lqmset47jnxe9090qt"
	require.NoError(t, err)
	csb = sha256.Sum256([]byte("Test Charity" + bech32addr))
	checksum = hex.EncodeToString(csb[:])
	params.Charities[0].AccAddress = bech32addr
	params.Charities[0].Checksum = checksum
	app.CharityKeeper.SetParams(app.Ctx, params)

	charities = app.CharityKeeper.GetCharity(app.Ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(app.Ctx, coins, charity)
		require.Errorf(t, err, "account does not exist for the provided address")
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
	app.CharityKeeper.SetParams(app.Ctx, params)
	require.Equal(t, app.CharityKeeper.GetAllParams(app.Ctx), params)
	// Test DonateCharity with invalid checksum: expect error
	charities = app.CharityKeeper.GetCharity(app.Ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(app.Ctx, coins, charity)
		require.Errorf(t, err, "checksum is invalid")
	}

	//Test CalculateSplit and DisburseDonations
	// Configure test charity accounts
	bech32addr1 := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	bech32addr2 := "enci1vvcw744ck9kzczrhf282lqmset47jnxe9090qt"
	addr1, err := sdk.AccAddressFromBech32(bech32addr1)
	require.NoError(t, err)
	addr2, err := sdk.AccAddressFromBech32(bech32addr2)
	require.NoError(t, err)
	acc1 := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr1)
	require.NotNil(t, acc1)
	acc2 := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr2)
	require.NotNil(t, acc2)

	// Set accounts to store
	app.AccountKeeper.SetAccount(app.Ctx, acc1)
	app.AccountKeeper.SetAccount(app.Ctx, acc2)
	// Create checksums and encode to strings
	csb1 := sha256.Sum256([]byte("Test Charity" + bech32addr1))
	checksum1 := hex.EncodeToString(csb1[:])
	require.NotEqual(t, "", checksum1)

	csb2 := sha256.Sum256([]byte("Test Charity 2" + bech32addr2))
	checksum2 := hex.EncodeToString(csb2[:])

	params.Charities = []types.Charity{
		{CharityName: "Test Charity",
			AccAddress: bech32addr1,
			Checksum:   checksum1,
		},
		{CharityName: "Test Charity 2",
			AccAddress: bech32addr2,
			Checksum:   checksum2},
	}

	app.CharityKeeper.SetParams(app.Ctx, params)
	require.Equal(t, app.CharityKeeper.GetAllParams(app.Ctx), params)
	charities = app.CharityKeeper.GetCharity(app.Ctx)
	require.Equal(t, params.Charities, charities)

	taxaddr := app.AccountKeeper.GetModuleAddress(types.CharityCollectorName)
	balance := app.BankKeeper.SpendableCoins(app.Ctx, taxaddr)
	baseamt := app.BankKeeper.GetAllBalances(app.Ctx, taxaddr)
	hasbalance := app.BankKeeper.HasBalance(app.Ctx, taxaddr, sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(int64(1000))))
	require.Equal(t, baseamt, balance)
	require.True(t, hasbalance)

	//Test calculate split
	split := app.CharityKeeper.CalculateSplit(app.Ctx, app.CharityKeeper.GetCharity(app.Ctx))
	require.Equal(t, baseamt[0].Amount.Quo(sdk.NewInt(int64(2))), split[0].Amount)
	//Test DisburseDonations
	payouts, _ := app.CharityKeeper.DisburseDonations(app.Ctx, charities)
	require.Equal(t, []types.Payout{{Recipientaddr: bech32addr1, Coins: split}, {Recipientaddr: bech32addr2, Coins: split}}, payouts)

}

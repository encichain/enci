package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	coretypes "github.com/user/encichain/types"

	"github.com/user/encichain/x/charity/types"
)

func TestDonateCharity(t *testing.T) {
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

}

func TestDonateCharityInvalidChecksum(t *testing.T) {
	app := CreateTestApp(t)
	sdk.GetConfig().SetBech32PrefixForAccount("enci", "encipub")

	bech32addr := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv55"
	addrforinvchecksum := "enci1aag23fr2qjxan9aktyfsywp3udxg036c9zxv54"
	addr, err := sdk.AccAddressFromBech32(bech32addr)
	require.NoError(t, err)
	acc := app.AccountKeeper.NewAccountWithAddress(app.Ctx, addr)
	require.NotNil(t, acc)
	// Create invalid checksum and encode to string
	csb := sha256.Sum256([]byte("Test Charity" + addrforinvchecksum))
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

	// Test DonateCharity: expect error
	charities := app.CharityKeeper.GetCharity(app.Ctx)
	for _, charity := range charities {
		err = app.CharityKeeper.DonateCharity(app.Ctx, coins, charity)
		require.Errorf(t, err, "checksum is invalid")
	}
}

package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	customante "github.com/encichain/enci/customcore/ante"
	coretypes "github.com/encichain/enci/types"
	"github.com/encichain/enci/x/charity/types"
)

func (suite *AnteTestSuite) TestDeductFees() {
	suite.SetupTest(false) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// Set 0 TaxRateMin *RateMin*
	newTaxLimit := types.TaxRateLimits{RateMin: sdk.NewDec(int64(0)), TaxRateMax: sdk.NewDecWithPrec(2, 2), BurnRateMax: types.DefaultBurnRateMax}
	suite.app.CharityKeeper.SetTaxRateLimits(suite.ctx, newTaxLimit)
	taxRateLim := suite.app.CharityKeeper.GetTaxRateLimits(suite.ctx)
	suite.Require().Equal(newTaxLimit, taxRateLim)

	// Set 0 taxrate
	suite.app.CharityKeeper.SetTaxRate(suite.ctx, sdk.NewDec(int64(0)))
	suite.Require().True(suite.app.CharityKeeper.GetTaxRate(suite.ctx).IsZero())

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// msg and signatures
	amt := int64(1000000)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, amt))
	msg := banktypes.NewMsgSend(addr1, addr1, coins)
	feeAmount := NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.Require().NoError(suite.txBuilder.SetMsgs(msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	// Set account with insufficient funds
	acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, addr1)
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	coins = sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(10)))
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, coins)
	suite.Require().NoError(err)

	dtd := customante.NewDeductTaxFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, nil)
	antehandler := sdk.ChainAnteDecorators(dtd)

	_, err = antehandler(suite.ctx, tx, false)

	suite.Require().NotNil(err, "Tx did not error when fee payer had insufficient funds")

	// Set account with sufficient funds
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(200))))
	suite.Require().NoError(err)

	// Ensure fee is not in fee collector account
	fcacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	feeAccBal := suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), feeAmount[0])
	suite.Require().False(feeAccBal)

	_, err = antehandler(suite.ctx, tx, false)

	suite.Require().Nil(err, "Tx errored after account has been set with sufficient funds")

	// Ensure fee has been sent to fee collector account
	feeAccBal = suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), feeAmount[0])
	suite.Require().True(feeAccBal)
}

package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	customante "github.com/user/encichain/customcore/ante"
	coretypes "github.com/user/encichain/types"
	charitytypes "github.com/user/encichain/x/charity/types"
)

func (suite *AnteTestSuite) TestDeductTaxes() {
	suite.SetupTest(false) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

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

	dtd := customante.NewDeductTaxDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, suite.app.FeeGrantKeeper)
	antehandler := sdk.ChainAnteDecorators(dtd)

	_, err = antehandler(suite.ctx, tx, false)

	suite.Require().NotNil(err, "Tx did not error when fee payer had insufficient funds")
	taxRate := suite.app.CharityKeeper.GetTaxRate(suite.ctx)
	expectedTax := sdk.NewCoin(coretypes.MicroTokenDenom, taxRate.MulInt64(amt).TruncateInt())
	ctaxaddr := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)
	taxAccBal := suite.app.BankKeeper.HasBalance(suite.ctx, ctaxaddr.GetAddress(), expectedTax)
	suite.Require().False(taxAccBal)

	// Set account with sufficient funds
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(2000000))))
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)

	taxAccBal = suite.app.BankKeeper.HasBalance(suite.ctx, ctaxaddr.GetAddress(), expectedTax)
	suite.Require().True(taxAccBal)

	suite.Require().Nil(err, "Tx errored after account has been set with sufficient funds")
}

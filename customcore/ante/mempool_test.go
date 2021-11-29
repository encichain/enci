package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	customante "github.com/user/encichain/customcore/ante"
	coretypes "github.com/user/encichain/types"
)

func (suite *AnteTestSuite) TestEnsureMempoolFeeTax() {
	suite.SetupTest(true) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	mfd := customante.NewMempoolFeeTaxDecorator(suite.app.CharityKeeper)
	antehandler := sdk.ChainAnteDecorators(mfd)

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// msg and signatures
	msg := testdata.NewTestMsg(addr1)
	feeAmount := NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.Require().NoError(suite.txBuilder.SetMsgs(msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	// Set high gas price so standard test fee fails
	uenciPrice := sdk.NewDecCoinFromDec(coretypes.MicroTokenDenom, sdk.NewDec(200).Quo(sdk.NewDec(100000)))
	highGasPrice := []sdk.DecCoin{uenciPrice}
	suite.ctx = suite.ctx.WithMinGasPrices(highGasPrice)

	// Set IsCheckTx to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	// antehandler errors with insufficient fees
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NotNil(err, "Decorator should have errored on too low fee for local gasPrice")

	// Set IsCheckTx to false
	suite.ctx = suite.ctx.WithIsCheckTx(false)

	// antehandler should not error since we do not check minGasPrice in DeliverTx
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Nil(err, "MempoolFeeDecorator returned error in DeliverTx")

	// Set IsCheckTx back to true for testing sufficient mempool fee
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	uenciPrice = sdk.NewDecCoinFromDec(coretypes.MicroTokenDenom, sdk.NewDec(0).Quo(sdk.NewDec(100000)))
	lowGasPrice := []sdk.DecCoin{uenciPrice}
	suite.ctx = suite.ctx.WithMinGasPrices(lowGasPrice)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Nil(err, "Decorator should not have errored on fee higher than local gasPrice")

	// Low gasprice but fees insufficient for tax
	taxRate := suite.app.CharityKeeper.GetTaxRate(suite.ctx)
	amt := sdk.NewDecFromInt(feeAmount[0].Amount).Quo(taxRate).TruncateInt().Add(sdk.NewInt(int64(10000)))
	highTaxCoin := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, amt))

	bankmsg := banktypes.NewMsgSend(addr1, addr1, highTaxCoin)
	suite.Require().NoError(suite.txBuilder.SetMsgs(bankmsg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	tx, err = suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NotNil(err, "Decorator should have errored on too low fee for tax")

	// Set IsCheckTx back to False for testing mempool sufficient charity tax fee check in DeliverTx
	suite.ctx = suite.ctx.WithIsCheckTx(false)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NotNil(err, "Decorator should have errored on too low fee for tax in DeliverTx")

	// Set IsCheckTx back to True
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	// non-Zero gas fee and tax. Sufficient fees
	uenciPrice = sdk.NewDecCoinFromDec(coretypes.MicroTokenDenom, sdk.NewDec(100).Quo(sdk.NewDec(100000)))
	lowerGasPrice := []sdk.DecCoin{uenciPrice}
	suite.ctx = suite.ctx.WithMinGasPrices(lowerGasPrice)

	sufficientfee := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, feeAmount[0].Amount.Add(sdk.NewInt(int64(1000)))))
	suite.txBuilder.SetFeeAmount(sufficientfee)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Nil(err, "Decorator should not have errored on sufficient fee for non-zero gas fee and tax")
}

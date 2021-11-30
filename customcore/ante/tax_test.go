package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	customante "github.com/user/encichain/customcore/ante"
	coretypes "github.com/user/encichain/types"
	charitytypes "github.com/user/encichain/x/charity/types"
)

func (suite *AnteTestSuite) TestDeductTaxesNoGasFee() {
	suite.SetupTest(false) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// fees, gas limit, msg and signatures
	//amt == DefaultTaxCap
	amt := int64(1000000)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, amt))
	msg := banktypes.NewMsgSend(addr1, addr1, coins)
	taxRate := suite.app.CharityKeeper.GetTaxRate(suite.ctx)
	// Assume 0 gas fees, only tax
	feeAmount := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, taxRate.MulInt64(amt).TruncateInt()))
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

	dtd := customante.NewDeductTaxFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, suite.app.FeeGrantKeeper)
	antehandler := sdk.ChainAnteDecorators(dtd)

	_, err = antehandler(suite.ctx, tx, false)

	suite.Require().NotNil(err, "Tx did not error when fee payer had insufficient funds")
	expectedTax := sdk.NewCoin(coretypes.MicroTokenDenom, feeAmount[0].Amount)
	ctaxacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)
	taxAccBal := suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), expectedTax)
	suite.Require().False(taxAccBal)

	// Set account with sufficient funds
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(2000000))))
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)

	taxAccBal = suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), expectedTax)
	suite.Require().True(taxAccBal)

	suite.Require().Nil(err, "Tx errored after account has been set with sufficient funds")
}

func (suite *AnteTestSuite) TestDeductTaxesFees() {
	suite.SetupTest(true)
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// fees, gas limit, msg and signatures
	//amt == DefaultTaxCap
	amt := int64(1000000)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, amt))
	msg := banktypes.NewMsgSend(addr1, addr1, coins)
	taxRate := suite.app.CharityKeeper.GetTaxRate(suite.ctx)

	// Both tax and gas fees
	feeAmount := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, taxRate.MulInt64(amt).TruncateInt().Add(sdk.NewInt(int64(100)))))
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

	dtd := customante.NewDeductTaxFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, suite.app.FeeGrantKeeper)
	antehandler := sdk.ChainAnteDecorators(dtd)

	_, err = antehandler(suite.ctx, tx, false)

	suite.Require().NotNil(err, "Tx did not error when fee payer had insufficient funds")
	expectedTax := sdk.NewCoin(coretypes.MicroTokenDenom, feeAmount[0].Amount.Sub(sdk.NewInt(int64(100))))
	ctaxacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)
	taxAccBal := suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), expectedTax)
	suite.Require().False(taxAccBal)

	// Set account with sufficient funds
	suite.app.AccountKeeper.SetAccount(suite.ctx, acc)
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(2000000))))
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)

	taxAccBal = suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), expectedTax)
	suite.Require().True(taxAccBal)

	suite.Require().Nil(err, "Tx errored after account has been set with sufficient funds")
}

func (suite *AnteTestSuite) TestAnteMultiSend() {
	suite.SetupTest(true) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// msg and signatures
	amount := int64(1000000)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, amount))
	msg := banktypes.NewMsgMultiSend(
		[]banktypes.Input{
			banktypes.NewInput(addr1, coins),
			banktypes.NewInput(addr1, coins),
		},
		[]banktypes.Output{
			banktypes.NewOutput(addr1, coins.Add(coins...)),
		},
	)

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
	coins = sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(100)))
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, coins)
	suite.Require().NoError(err)

	mftd := customante.NewMempoolFeeTaxDecorator(suite.app.CharityKeeper)
	dtd := customante.NewDeductTaxFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, suite.app.FeeGrantKeeper)
	antehandler := sdk.ChainAnteDecorators(mftd, dtd)

	//Set IsCheckTx to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	//Set gas prices, low gas
	uenciPrice := sdk.NewDecCoinFromDec(coretypes.MicroTokenDenom, sdk.NewDec(50).Quo(sdk.NewDec(100000)))
	lowGasPrice := []sdk.DecCoin{uenciPrice}
	suite.ctx = suite.ctx.WithMinGasPrices(lowGasPrice)

	expectedTaxAmt := suite.app.CharityKeeper.GetTaxRate(suite.ctx).MulInt64(amount).TruncateInt()
	expectedTax := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt))

	// Set too low fee for gas + tax
	suite.txBuilder.SetFeeAmount(expectedTax)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "Mempool Decorator should have errored on too low fee for local gas price + tax")

	//Set IsCheckTx to false to check deliverTx
	suite.ctx = suite.ctx.WithIsCheckTx(true)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "Mempool Decorator in DeliverTx should have errored on too low fee for local gas price + tax")

	//Set IsCheckTX to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	//Sufficient tax but not gasfee
	suite.txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt.Add(expectedTaxAmt))))
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "Mempool Decorator should have errored on too low fee for local gas price + tax")

	// Ensure fee+tax collector accounts do not have balance
	ctaxacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)
	fcacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	hasFeeBal := suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), sdk.NewInt64Coin(coretypes.MicroTokenDenom, int64(1)))
	hasTaxBal := suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), sdk.NewInt64Coin(coretypes.MicroTokenDenom, int64(1)))
	suite.Require().False(hasFeeBal)
	suite.Require().False(hasTaxBal)

	//Sufficient fees but insufficient funds
	suffFee := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt.Add(expectedTaxAmt).Add(sdk.NewInt(int64(200)))))
	suite.txBuilder.SetFeeAmount(suffFee)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "DeductTaxFee decorator should have errored on insufficient funds to pay fee for local gas price + tax")

	// Fund account with sufficient funds
	coins = sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(1000000)))
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, coins)
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NoError(err, "Decorator should not have errored due to sufficient fees and sufficient funds.")

	// Check balances
	hasFeeBal = suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), suffFee[0].SubAmount(expectedTaxAmt).SubAmount(expectedTaxAmt))
	suite.Require().True(hasFeeBal)

	hasTaxBal = suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), suffFee[0].SubAmount(sdk.NewInt(int64(200))))
	suite.Require().True(hasTaxBal)

	bal := suite.app.BankKeeper.GetBalance(suite.ctx, ctaxacc.GetAddress(), coretypes.MicroTokenDenom)
	suite.Require().Equal(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt.Add(expectedTaxAmt)), bal)
}

func (suite *AnteTestSuite) TestAnteAuthzExec() {

	suite.SetupTest(true) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	// keys and addresses
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// msg and signatures
	amount := int64(1000000)
	coins := sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, amount))
	msg := authz.NewMsgExec(addr1, []sdk.Msg{banktypes.NewMsgSend(addr1, addr1, coins)})

	feeAmount := NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.Require().NoError(suite.txBuilder.SetMsgs(&msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	//Set gas prices, low gas
	uenciPrice := sdk.NewDecCoinFromDec(coretypes.MicroTokenDenom, sdk.NewDec(50).Quo(sdk.NewDec(100000)))
	lowGasPrice := []sdk.DecCoin{uenciPrice}
	suite.ctx = suite.ctx.WithMinGasPrices(lowGasPrice)

	expectedTaxAmt := suite.app.CharityKeeper.GetTaxRate(suite.ctx).MulInt64(amount).TruncateInt()
	expectedTax := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt))

	// Set antehandler
	mtd := customante.NewMempoolFeeTaxDecorator(suite.app.CharityKeeper)
	dtd := customante.NewDeductTaxFeeDecorator(suite.app.AccountKeeper, suite.app.BankKeeper, suite.app.CharityKeeper, suite.app.FeeGrantKeeper)
	antehandler := sdk.ChainAnteDecorators(mtd, dtd)

	// Set IsCheckTx to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	// Too low fee for gas + tax
	suite.txBuilder.SetFeeAmount(expectedTax)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "Mempool Decorator should have errored on too low fee for local gas price + tax")

	//Set IsCheckTx to false to check deliverTx
	suite.ctx = suite.ctx.WithIsCheckTx(false)
	antemem := sdk.ChainAnteDecorators(mtd)
	_, err = antemem(suite.ctx, tx, false)
	suite.Require().NoError(err, "Mempool Decorator in DeliverTx should not have errored due to sufficient fee for tax in deliverTx")

	//Set IsCheckTX to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	//Sufficient tax but not gasfee
	suite.txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt.Add(expectedTaxAmt))))
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "Mempool Decorator should have errored on too low fee for local gas price + tax")

	// Ensure fee+tax collector accounts do not have balance
	ctaxacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, charitytypes.CharityCollectorName)
	fcacc := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName)
	hasFeeBal := suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), sdk.NewInt64Coin(coretypes.MicroTokenDenom, int64(1)))
	hasTaxBal := suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), sdk.NewInt64Coin(coretypes.MicroTokenDenom, int64(1)))
	suite.Require().False(hasFeeBal)
	suite.Require().False(hasTaxBal)

	//Sufficient fees but insufficient account funds
	suffFee := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt.Add(sdk.NewInt(int64(200)))))
	suite.txBuilder.SetFeeAmount(suffFee)
	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err, "DeductTaxFee decorator should have errored on insufficient funds to pay fee for local gas price + tax")

	// Fund account with sufficient funds
	coins = sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, sdk.NewInt(1000000)))
	err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr1, coins)
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NoError(err, "Decorator should not have errored due to sufficient fees and sufficient funds.")

	// Check balances
	hasFeeBal = suite.app.BankKeeper.HasBalance(suite.ctx, fcacc.GetAddress(), suffFee[0].SubAmount(expectedTaxAmt))
	suite.Require().True(hasFeeBal)

	hasTaxBal = suite.app.BankKeeper.HasBalance(suite.ctx, ctaxacc.GetAddress(), suffFee[0].SubAmount(sdk.NewInt(int64(200))))
	suite.Require().True(hasTaxBal)
	bal := suite.app.BankKeeper.GetBalance(suite.ctx, ctaxacc.GetAddress(), coretypes.MicroTokenDenom)
	suite.Require().Equal(sdk.NewCoin(coretypes.MicroTokenDenom, expectedTaxAmt), bal)
}

func (suite *AnteTestSuite) TestCalculateTaxLim() {
	suite.SetupTest(false) // setup
	app := suite.app
	ctx := suite.ctx
	defaultRateAmt := sdk.NewDecFromInt(sdk.NewInt(100)).Mul(charitytypes.DefaultTaxRate).TruncateInt()
	defaultRateCoins := sdk.NewCoins(sdk.NewCoin(coretypes.MicroTokenDenom, defaultRateAmt))
	// Set too high tax rate
	app.CharityKeeper.SetTaxRate(ctx, sdk.NewDecWithPrec(6, 2))
	taxRate := app.CharityKeeper.GetTaxRate(ctx)
	suite.Require().NotEqual(charitytypes.DefaultTaxRate, taxRate)

	tax := customante.ComputeTax(ctx, app.CharityKeeper, sdk.NewCoins(sdk.NewInt64Coin(coretypes.MicroTokenDenom, 100)))
	suite.Require().Equal(defaultRateCoins, tax)
}

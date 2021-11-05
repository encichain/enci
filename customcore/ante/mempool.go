package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MempoolFeeTaxDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config) + charity tax.
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeTaxDecorator
type MempoolFeeTaxDecorator struct {
	charityKeeper CharityKeeper
}

func NewMempoolFeeTaxDecorator(ck CharityKeeper) MempoolFeeTaxDecorator {
	return MempoolFeeTaxDecorator{
		charityKeeper: ck,
	}
}

func (mfd MempoolFeeTaxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()
	msgs := feeTx.GetMsgs()

	if !simulate {
		// Calculate taxes
		tax := ParseMsgAndComputeTax(ctx, mfd.charityKeeper, msgs...)

		// Ensure that the provided fees meet a minimum threshold for the validator + charity tax,
		// if this is a CheckTx. This is only for local mempool purposes, and thus
		// is only ran on check tx.
		if ctx.IsCheckTx() {
			err := CheckMempoolFeeTax(ctx, gas, feeCoins, tax)
			if err != nil {
				return ctx, err
			}
		}
		// Ensure supplied fee is not lower than tax amount
		if _, hasNeg := feeCoins.SafeSub(tax); hasNeg {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, tax)
		}
	}

	return next(ctx, tx, simulate)
}

// EnsureSufficientMempoolFees verifies that the given transaction has supplied
// enough fees to cover a proposer's minimum fees.
// Error is returned on failure
func CheckMempoolFeeTax(ctx sdk.Context, gas uint64, fees sdk.Coins, tax sdk.Coins) error {
	minGasPrices := ctx.MinGasPrices()
	requiredFees := sdk.Coins{}
	if !minGasPrices.IsZero() {
		requiredFees = make(sdk.Coins, len(minGasPrices))

		// Determine the required fees by multiplying each required minimum gas
		// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
		glDec := sdk.NewDec(int64(gas))
		for i, gp := range minGasPrices {
			fee := gp.Amount.Mul(glDec)
			requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}

	// Deduct taxes from fees to obtain fee(gas) without charity tax.
	// Checks for both insufficient tax and insufficient supplied gas fee
	feegas, neg := fees.SafeSub(tax)
	if neg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "1; insufficient fees(tax); got: %s, required: %s = %s(gas) +%s(charity tax)",
			feegas.Add(tax...), requiredFees.Add(tax...), requiredFees, tax)
	}

	// Ensure supplied fees(gas w/o tax) can cover required fees
	if !requiredFees.IsZero() && !feegas.IsAnyGTE(requiredFees) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "2; insufficient fees(gas); got: %s, required: %s = %s(gas) +%s(charity tax)",
			feegas.Add(tax...), requiredFees.Add(tax...), requiredFees, tax)
	}
	return nil
}

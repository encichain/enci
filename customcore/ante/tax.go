package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	charitytypes "github.com/user/encichain/x/charity/types"
)

// DeductTaxDecorator deducts charity tax from the first signer of the tx
// Tax is calculated based on tx.Amount, independent of Fee.
// If the first signer does not have the funds to pay for the tax, return with InsufficientFunds error
// Call next AnteHandler if tax successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductTaxDecorator
type DeductTaxDecorator struct {
	ak             AccountKeeper
	bankKeeper     types.BankKeeper
	CharityKeeper  CharityKeeper
	feegrantKeeper FeegrantKeeper
}

func NewDeductTaxDecorator(ak AccountKeeper, bk types.BankKeeper, ck CharityKeeper, fk FeegrantKeeper) DeductTaxDecorator {
	return DeductTaxDecorator{
		ak:             ak,
		bankKeeper:     bk,
		CharityKeeper:  ck,
		feegrantKeeper: fk,
	}
}

func (dtd DeductTaxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}
	// Ensure charity tax collector module account has been set
	if addr := dtd.ak.GetModuleAddress(charitytypes.CharityCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", charitytypes.CharityCollectorName))
	}

	tax := ParseMsgAndComputeTax(ctx, dtd.CharityKeeper, tx.GetMsgs()...)
	taxPayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()
	deductFeesFrom := taxPayer

	// if feegranter set deduct tax from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if dtd.feegrantKeeper == nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee grants are not enabled")
		} else if !feeGranter.Equals(taxPayer) {
			err := dtd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, taxPayer, tax, tx.GetMsgs())

			if err != nil {
				return ctx, sdkerrors.Wrapf(err, "%s not allowed to pay fees from %s", feeGranter, taxPayer)
			}
		}

		deductFeesFrom = feeGranter
	}

	deductTaxFromAcc := dtd.ak.GetAccount(ctx, deductFeesFrom)
	if deductTaxFromAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "tax payer address: %s does not exist", deductFeesFrom)
	}

	// deduct the fees
	if !tax.IsZero() {
		err = DeductFees(dtd.bankKeeper, ctx, deductTaxFromAcc, tax)
		if err != nil {
			return ctx, err
		}
		// Record the tax proceeds to store
		dtd.CharityKeeper.AddTaxProceeds(ctx, tax)
	}

	return next(ctx, tx, simulate)
}

// DeductFees deducts taxes from the given account. sending the taxes to the specified collector account
func DeductFees(bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, tax sdk.Coins) error {
	if !tax.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid tax amount: %s", tax)
	}

	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), charitytypes.CharityCollectorName, tax)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}

// Filter bank messages and compute tax on each transaction.
func ParseMsgAndComputeTax(ctx sdk.Context, ck CharityKeeper, msgs ...sdk.Msg) sdk.Coins {
	taxFinal := sdk.Coins{}

	for _, msg := range msgs {
		switch msg := msg.(type) {
		case *banktypes.MsgSend:
			taxFinal = taxFinal.Add(ComputeTax(ctx, ck, msg.Amount)...)

		case *banktypes.MsgMultiSend:
			for _, input := range msg.Inputs {
				taxFinal = taxFinal.Add(ComputeTax(ctx, ck, input.Coins)...)
			}
		case *authz.MsgExec:
			authzmsgs, err := msg.GetMessages()
			if err != nil {
				panic(err)
			}
			taxFinal = taxFinal.Add(ParseMsgAndComputeTax(ctx, ck, authzmsgs...)...)
		}
	}
	return taxFinal
}

// Compute the charity tax due
func ComputeTax(ctx sdk.Context, ck CharityKeeper, coins sdk.Coins) sdk.Coins {
	taxRate := ck.GetTaxRate(ctx)
	taxFinal := sdk.Coins{}
	if taxRate.Equal(sdk.ZeroDec()) {
		return taxFinal
	}

	for _, coin := range coins {
		taxOwed := sdk.NewDecFromInt(coin.Amount).Mul(taxRate).TruncateInt()

		if taxOwed.IsZero() {
			continue
		}

		// Check if taxOwed is greater than denom taxcap
		taxCap := ck.GetTaxCap(ctx, coin.Denom)
		if taxCap.IsNegative() || taxCap.IsZero() {
			taxCap = charitytypes.DefaultCap
		}
		if taxOwed.GT(taxCap) {
			taxOwed = taxCap
		}
		taxFinal = taxFinal.Add(sdk.NewCoin(coin.Denom, taxOwed))
	}
	return taxFinal
}

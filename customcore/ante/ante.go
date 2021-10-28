package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	cosmosante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// HandlerOptions are the options required for constructing a default SDK AnteHandler.
type HandlerOptions struct {
	AccountKeeper   AccountKeeper
	BankKeeper      types.BankKeeper
	FeegrantKeeper  FeegrantKeeper
	CharityKeeper   CharityKeeper
	SignModeHandler authsigning.SignModeHandler
	SigGasConsumer  func(meter sdk.GasMeter, sig signing.SignatureV2, params types.Params) error
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	var sigGasConsumer = options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = cosmosante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		cosmosante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		cosmosante.NewRejectExtensionOptionsDecorator(),
		cosmosante.NewMempoolFeeDecorator(),
		cosmosante.NewValidateBasicDecorator(),
		cosmosante.TxTimeoutHeightDecorator{},
		cosmosante.NewValidateMemoDecorator(options.AccountKeeper),
		cosmosante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		cosmosante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		NewDeductTaxDecorator(options.AccountKeeper, options.BankKeeper, options.CharityKeeper, options.FeegrantKeeper),
		cosmosante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		cosmosante.NewValidateSigCountDecorator(options.AccountKeeper),
		cosmosante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		cosmosante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		cosmosante.NewIncrementSequenceDecorator(options.AccountKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

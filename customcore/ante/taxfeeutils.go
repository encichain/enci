package ante

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/pflag"
	charitytypes "github.com/user/encichain/x/charity/types"
)

// Compute fees and tax with CLI options
func ComputeFeeTaxCli(clientCtx client.Context, flagset *pflag.FlagSet, msgs ...sdk.Msg) (fees sdk.Coins, gas uint64, err error) {
	// Create new tx factory
	txF := tx.NewFactoryCLI(clientCtx, flagset)
	gas = txF.Gas()

	if txF.SimulateAndExecute() {
		_, adjusted, err := tx.CalculateGas(clientCtx, txF, msgs...)
		if err != nil {
			return nil, gas, err
		}
		gas = adjusted
	}
	fees = txF.Fees()

	// Compute gasFee from gas price, if gas price is not zero
	// fee = ceil(gasPrice * gasLimit).
	gasPrices := txF.GasPrices()
	if !gasPrices.IsZero() {
		glDec := sdk.NewDec(int64(gas))
		gasFees := make(sdk.Coins, len(gasPrices))
		for i, gp := range gasPrices {
			fee := gp.Amount.Mul(glDec)
			gasFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}

		fees = fees.Add(gasFees.Sort()...)
	}
	// Compute taxes and add to current fees
	taxes, err := ClientParseMsgAndComputeTax(clientCtx, msgs...)
	if err != nil {
		return nil, gas, err
	}

	fees = fees.Add(taxes...)
	return
}

// Filter messages and compute tax on each MsgSend/MsgMultiSend
func ClientParseMsgAndComputeTax(clientCtx client.Context, msgs ...sdk.Msg) (taxFinal sdk.Coins, err error) {
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case *banktypes.MsgSend:
			taxFinal, err = clientComputeTax(clientCtx, msg.Amount)
			if err != nil {
				return nil, err
			}

		case *banktypes.MsgMultiSend:
			for _, input := range msg.Inputs {
				taxFinal, err = clientComputeTax(clientCtx, input.Coins)
			}
			if err != nil {
				return nil, err
			}
		case *authz.MsgExec:
			authzmsgs, err := msg.GetMessages()
			if err != nil {
				panic(err)
			}
			taxFinal, err = ClientParseMsgAndComputeTax(clientCtx, authzmsgs...)
			if err != nil {
				return nil, err
			}
		}
	}
	return
}

func clientComputeTax(clientCtx client.Context, coins sdk.Coins) (sdk.Coins, error) {
	taxRate, err := queryTaxRate(clientCtx)
	tax := sdk.Coins{}
	if err != nil {
		return nil, err
	}
	for _, coin := range coins {
		taxAmt := sdk.NewDecFromInt(coin.Amount).Mul(taxRate).TruncateInt()

		taxCap, err := queryTaxCap(clientCtx, coin.Denom)
		if err != nil {
			return nil, err
		}
		// Ensure tax due is not greater than tax cap
		if taxAmt.GT(taxCap) {
			taxAmt = taxCap
		}

		if taxAmt.IsZero() {
			continue
		}

		tax = tax.Add(sdk.NewCoin(coin.Denom, taxAmt))
	}

	return tax, nil
}

func queryTaxRate(clientCtx client.Context) (sdk.Dec, error) {
	queryClient := charitytypes.NewQueryClient(clientCtx)

	res, err := queryClient.TaxRate(context.Background(), &charitytypes.QueryTaxRateRequest{})
	return res.TaxRate, err
}

func queryTaxCap(clientCtx client.Context, denom string) (sdk.Int, error) {
	queryClient := charitytypes.NewQueryClient(clientCtx)

	res, err := queryClient.TaxCap(context.Background(), &charitytypes.QueryTaxCapRequest{Denom: denom})
	return res.Cap, err
}

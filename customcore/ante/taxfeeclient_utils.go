package ante

import (
	"context"

	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	charitytypes "github.com/encichain/enci/x/charity/types"
)

// Compute fees and tax with CLI options
func ComputeFeeTaxCli(clientCtx client.Context, flagset *pflag.FlagSet, msgs ...sdk.Msg) (fees sdk.Coins, gas uint64, err error) {
	// Create new tx factory
	txF := tx.NewFactoryCLI(clientCtx, flagset)
	gas = txF.Gas()

	if txF.SimulateAndExecute() {
		txF, err := prepareFactory(clientCtx, txF)
		if err != nil {
			return nil, gas, err
		}

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
	taxFinal := sdk.Coins{}
	if err != nil {
		return nil, err
	}
	if taxRate.Equal(sdk.ZeroDec()) {
		return taxFinal, nil
	}

	taxLims, err := queryTaxRateLimits(clientCtx)
	if err != nil {
		return nil, err
	}
	// Set default tax rate if the queried taxRate does not satisfy TaxRateLimits
	if taxRate.LT(taxLims.RateMin) || taxRate.GT(taxLims.TaxRateMax) {
		taxRate = charitytypes.DefaultTaxRate
	}

	//Compute tax on each sdk.Coin
	for _, coin := range coins {
		taxOwed := sdk.NewDecFromInt(coin.Amount).Mul(taxRate).TruncateInt()

		taxCap, err := queryTaxCap(clientCtx, coin.Denom)
		if err != nil {
			return nil, err
		}
		// Ensure tax owed is not greater than tax cap
		if taxCap.IsNegative() {
			taxCap = charitytypes.DefaultCap
		}
		if taxOwed.GT(taxCap) || taxCap.IsZero() {
			taxOwed = taxCap
		}
		if taxOwed.IsZero() {
			continue
		}
		taxFinal = taxFinal.Add(sdk.NewCoin(coin.Denom, taxOwed))
	}
	return taxFinal, nil
}

// queryTaxRate queries the set taxRate
func queryTaxRate(clientCtx client.Context) (sdk.Dec, error) {
	queryClient := charitytypes.NewQueryClient(clientCtx)

	res, err := queryClient.TaxRate(context.Background(), &charitytypes.QueryTaxRateRequest{})
	return res.TaxRate, err
}

// queryTaxCap queries the KVStore taxCaps
func queryTaxCap(clientCtx client.Context, denom string) (sdk.Int, error) {
	queryClient := charitytypes.NewQueryClient(clientCtx)

	res, err := queryClient.TaxCap(context.Background(), &charitytypes.QueryTaxCapRequest{Denom: denom})
	return res.Cap, err
}

// queryTaxRateLimits queries the set TaxRateLimits
func queryTaxRateLimits(clientCtx client.Context) (charitytypes.TaxRateLimits, error) {
	queryClient := charitytypes.NewQueryClient(clientCtx)

	res, err := queryClient.TaxRateLimits(context.Background(), &charitytypes.QueryTaxRateLimitsRequest{})

	return res.TaxRateLimits, err
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func prepareFactory(clientCtx client.Context, txf tx.Factory) (tx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

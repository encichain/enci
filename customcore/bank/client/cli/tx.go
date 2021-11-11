package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/user/encichain/customcore/ante"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Bank transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(NewSendTxCmd())

	return txCmd
}

// NewSendTxCmd returns a CLI command handler for creating a MsgSend transaction.
func NewSendTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "send [from_key_or_address] [to_address] [amount]",
		Short: `Send funds from one account to another. Note, the'--from' flag is
ignored as it is implied from [from_key_or_address].`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			toAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(clientCtx.GetFromAddress(), toAddr, coins)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			//Create new transaction factory
			txF := tx.NewFactoryCLI(clientCtx, cmd.Flags())

			if !clientCtx.GenerateOnly && txF.Fees().IsZero() {
				// estimate tax and gas fees
				fees, gas, err := ante.ComputeFeeTaxCli(clientCtx, cmd.Flags(), msg)
				if err != nil {
					return err
				}
				// Update tx
				txF = txF.WithFees(fees.String()).
					WithGas(gas).
					WithSimulateAndExecute(false).
					WithGasPrices("")
			}
			// Generate tx and print if clientCtx.GenerateOnly == true ; sign and broadcast if false
			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txF, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

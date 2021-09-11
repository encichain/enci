package cli

import (
	"context"
	"fmt"

	//"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/user/charity/x/charity/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group charity queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(
		CmdQueryTaxRate(),
		CmdQueryParams(),
		CmdQueryCharityOne(),
		CmdQueryCharityTwo(),
	)

	return cmd
}

// CmdQueryTaxRate implements the query taxrate command and returns the TaxRate
func CmdQueryTaxRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taxrate",
		Args:  cobra.NoArgs,
		Short: "Query the tax rate",
		Long:  "Query the tax rate. Returned is a decimal.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TaxRate(context.Background(), &types.QueryTaxRateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryParams implements the query params command and returns all params
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query all params",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCharityOne implements the query Charity One command and returns the charity one object
func CmdQueryCharityOne() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charity-one",
		Args:  cobra.NoArgs,
		Short: "Query charity one",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CharityOne(context.Background(), &types.QueryCharityOneRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCharityTwo implements the query Charity Two command and returns the charity two object
func CmdQueryCharityTwo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charity-two",
		Args:  cobra.NoArgs,
		Short: "Query charity two",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CharityTwo(context.Background(), &types.QueryCharityTwoRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

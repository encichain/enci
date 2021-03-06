package cli

import (
	"context"
	"fmt"

	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/encichain/enci/x/charity/types"
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
		CmdQueryCharities(),
		CmdQueryTaxCap(),
		CmdQueryTaxCaps(),
		CmdQueryBurnRate(),
		CmdQueryTaxRateLimits(),
		CmdQueryTaxProceeds(),
		CmdQueryCollectionEpochs(),
		CmdQueryCollectionEpoch(),
		//CmdQueryChecksum(),
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
func CmdQueryCharities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charities",
		Args:  cobra.NoArgs,
		Short: "Query all charities",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Charities(context.Background(), &types.QueryCharitiesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTaxCap implements the query taxcap command and returns the taxcap of a *denom*
func CmdQueryTaxCap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taxcap [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query the tax cap of a denom",
		Long:  "Query the tax cap of a denom. A TaxCap is the maximum amount of tax that can be charged per transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			denom := args[0]
			res, err := queryClient.TaxCap(context.Background(), &types.QueryTaxCapRequest{Denom: denom})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTaxCaps implements the query taxcaps command and returns all taxcaps for all denoms
func CmdQueryTaxCaps() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taxcaps",
		Args:  cobra.NoArgs,
		Short: "Query all existing Taxcaps for all denom assets.",
		Long:  "Query all existing Taxcaps for all denom assets.. A TaxCap is the maximum amount of tax that can be charged per transaction.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TaxCaps(context.Background(), &types.QueryTaxCapsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryBurnRate implements the query BurnRate command and returns the BurnRate
func CmdQueryBurnRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-rate",
		Args:  cobra.NoArgs,
		Short: "Query the charity burn rate",
		Long: `Query the charity burn rate. Returned is a decimal, representing the percent
				of the charity tax proceeds that is burned at the end of each epoch`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.BurnRate(context.Background(), &types.QueryBurnRateRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTaxRateLimits implements the query TaxRateLimits command and returns the limits to TaxRate
func CmdQueryTaxRateLimits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taxratelimits",
		Args:  cobra.NoArgs,
		Short: "Query Tax Rate Limits",
		Long:  "Query Tax Rate Limits. Tax Rate Limits determines the minimum and maximum tax rate that can be set.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TaxRateLimits(context.Background(), &types.QueryTaxRateLimitsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryTaxProceeds implements the query taxproceeds command and returns the tax collected for the current CollectionEpoch
func CmdQueryTaxProceeds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taxproceeds",
		Args:  cobra.NoArgs,
		Short: "Query the tax collected for the current collection epoch",
		Long:  "Query the tax collected for the current collection epoch. This represents the proceeds from the charity tax.",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.TaxProceeds(context.Background(), &types.QueryTaxProceedsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCollectionEpochs implements the query collectionepochs command and returns all CollectionEpochs
func CmdQueryCollectionEpochs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collection-epochs",
		Args:  cobra.NoArgs,
		Short: "Query all collection epochs",
		Long: strings.TrimSpace(`
		Query all collection epochs.
		This returns the data from all previous collection epochs.
	
		$ encid query charity collection-epochs
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CollectionEpochs(context.Background(), &types.QueryAllCollectionEpochsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCollectionEpochs implements the query collectionepoch command and returns a CollectionEpoch based on epoch
func CmdQueryCollectionEpoch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collection-epoch [epoch]",
		Args:  cobra.ExactArgs(1),
		Short: "Query a collection epoch based on epoch",
		Long: strings.TrimSpace(`
		Query a collection epoch.
	
		$ encid query charity collection-epoch 0
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			epoch, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			res, err := queryClient.CollectionEpoch(context.Background(), &types.QueryCollectionEpochRequest{Epoch: uint64(epoch)})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

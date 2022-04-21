package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/encichain/enci/x/oracle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group oracle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryDelegateAddress())
	cmd.AddCommand(CmdQueryDelegatorAddress())
	cmd.AddCommand(CmdQueryClaimTypes())
	cmd.AddCommand(CmdQueryVoteRounds())
	cmd.AddCommand(CmdQueryPrevoteRounds())
	cmd.AddCommand(CmdQueryNextVotePeriod())
	cmd.AddCommand(CmdQueryNextPrevotePeriod())
	cmd.AddCommand(CmdQueryIsVotePeriod())
	cmd.AddCommand(CmdQueryIsPrevotePeriod())

	return cmd
}

// CmdParams implements a command to fetch oracle parameters.
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "query the current oracle parameters",
		Args:  cobra.NoArgs,
		Long: strings.TrimSpace(`Query genesis parameters for the oracle module:

$ <appd> query oracle params
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}
			res, err := queryClient.Params(context.Background(), params)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// CmdDelegeteAddress queries the delegate address from the chain given validators address
func CmdQueryDelegateAddress() *cobra.Command {
	return &cobra.Command{
		Use:     "delegate-address [validator-address]",
		Aliases: []string{"del"},
		Args:    cobra.ExactArgs(1),
		Short:   "query delegate address from the chain given validators address",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryDelegateAddressRequest{Validator: args[0]}

			res, err := queryClient.DelegateAddress(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
}

// CmdValidatorAddress queries the validator address from the chain given the addres of a voter delegate
func CmdQueryDelegatorAddress() *cobra.Command {
	return &cobra.Command{
		Use:     "validator-address [delegate-address]",
		Aliases: []string{"val"},
		Args:    cobra.ExactArgs(1),
		Short:   "query validator address from the chain given the address that validator delegated to",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(ctx)
			req := &types.QueryDelegatorAddressRequest{Delegate: args[0]}

			res, err := queryClient.DelegatorAddress(cmd.Context(), req)
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}
}

// CmdQueryClaimTypes queries the registered claim types
func CmdQueryClaimTypes() *cobra.Command {
	return &cobra.Command{
		Use:   "claim-types",
		Args:  cobra.NoArgs,
		Short: "Query all registered claim types",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ClaimTypes(cmd.Context(), &types.QueryClaimTypesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
}

// CmdQueryVoteRounds queries all votes for each claim type
func CmdQueryVoteRounds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote-rounds",
		Args:  cobra.NoArgs,
		Short: "Query all vote rounds containing all votes for current voting period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.VoteRounds(cmd.Context(), &types.QueryVoteRoundsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryPrevoteRounds queries all prevotes for each claim type
func CmdQueryPrevoteRounds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prevote-rounds",
		Args:  cobra.NoArgs,
		Short: "Query all prevote rounds containing all prevotes for current voting period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PrevoteRounds(cmd.Context(), &types.QueryPrevoteRoundsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryNextVotePeriod queries the block height of the next vote period
func CmdQueryNextVotePeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-vote-period",
		Args:  cobra.NoArgs,
		Short: "Query block height of the next vote period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			height, err := rpc.GetChainHeight(clientCtx)
			if err != nil {
				return err
			}
			clientCtx.Height = height
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Printf("Current height: %d", clientCtx.Height)

			res, err := queryClient.NextVotePeriod(cmd.Context(), &types.QueryNextVotePeriodRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryNextPrevotePeriod queries the block height of the next prevote period
func CmdQueryNextPrevotePeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-prevote-period",
		Args:  cobra.NoArgs,
		Short: "Query block height of the next prevote period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			height, err := rpc.GetChainHeight(clientCtx)
			if err != nil {
				return err
			}
			clientCtx.Height = height
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Printf("Current height: %d", clientCtx.Height)

			res, err := queryClient.NextPrevote(cmd.Context(), &types.QueryNextPrevoteRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryIsVotePeriod queries if the current block is part of a voting period
func CmdQueryIsVotePeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-vote-period",
		Args:  cobra.NoArgs,
		Short: "Query if current block is part of a vote period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			height, err := rpc.GetChainHeight(clientCtx)
			if err != nil {
				return err
			}
			clientCtx.Height = height
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Printf("Current height: %d", clientCtx.Height)

			res, err := queryClient.IsVotePeriod(cmd.Context(), &types.QueryIsVotePeriodRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryIsPrevotePeriod queries if the current block is part of a prevote period
func CmdQueryIsPrevotePeriod() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-prevote-period",
		Args:  cobra.NoArgs,
		Short: "Query if current block is part of a prevote period",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			height, err := rpc.GetChainHeight(clientCtx)
			if err != nil {
				return err
			}
			clientCtx.Height = height
			queryClient := types.NewQueryClient(clientCtx)
			fmt.Printf("Current height: %d", clientCtx.Height)

			res, err := queryClient.IsPrevotePeriod(cmd.Context(), &types.QueryIsPrevotePeriodRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

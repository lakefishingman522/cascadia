package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for the inflation module.
func GetQueryCmd() *cobra.Command {
	inflationQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the inflation module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	inflationQueryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryPenaltyAccount(),
	)

	return inflationQueryCmd
}

// GetCmdQueryParams implements a command to return the current inflation
// parameters.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current inflation parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryParamsRequest{}
			res, err := queryClient.Params(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryInflation implements a command to return the current inflation
// inflation value.
func GetCmdQueryPenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "Query the current penalty account address",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryPenaltyAccountRequest{}
			res, err := queryClient.PenaltyAccount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintString(fmt.Sprintf("%s\n", res.Address))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

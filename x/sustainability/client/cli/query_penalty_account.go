package cli

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowPenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-penalty-account",
		Short: "shows penalty_account",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPenaltyAccountRequest{}

			res, err := queryClient.PenaltyAccount(context.Background(), params)

			if res == nil {
				return clientCtx.PrintProto(res)
			}

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

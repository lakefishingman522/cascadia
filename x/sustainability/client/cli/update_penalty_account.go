package cli

import (
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

func CmdUpdatePenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-penalty-account [multi-sig account address]",
		Short: "Update penalty account in staking module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if _, err = sdk.AccAddressFromBech32(argAddress); err != nil {
				return err
			}
			msg := types.NewMsgUpdatePenaltyAccountRequest(clientCtx.GetFromAddress().String(), argAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

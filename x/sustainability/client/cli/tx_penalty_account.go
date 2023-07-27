package cli

import (
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

func CmdCreatePenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-penalty-account [multisig-address]",
		Short: "Create penalty_account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMultisigAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			if _, err = sdktypes.AccAddressFromBech32(argMultisigAddress); err != nil {
				return err
			}

			msg := types.NewMsgCreatePenaltyAccount(clientCtx.GetFromAddress().String(), argMultisigAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-penalty-account [multisig-address]",
		Short: "Update penalty_account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMultisigAddress := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePenaltyAccount(clientCtx.GetFromAddress().String(), argMultisigAddress)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePenaltyAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-penalty-account",
		Short: "Delete penalty_account",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePenaltyAccount(clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

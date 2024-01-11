package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdUpdatePriceStatistics() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-price-statistics [p-360] [p-180] [p-90] [p-30] [p-14] [p-7] [p-1]",
		Short: "Broadcast message update-price-statistics",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argP360 := args[0]
             argP180 := args[1]
             argP90 := args[2]
             argP30 := args[3]
             argP14 := args[4]
             argP7 := args[5]
             argP1 := args[6]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePriceStatistics(
				clientCtx.GetFromAddress().String(),
				argP360,
				argP180,
				argP90,
				argP30,
				argP14,
				argP7,
				argP1,
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}
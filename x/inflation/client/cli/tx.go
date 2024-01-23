package cli

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func NewSubmitCreateInflationControlParamsProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-inflation-control-params-proposal [title] [description] [lambda] [w360] [w180] [w90] [w30] [w14] [w7] [w1]",
		Short: "Submit a new proposal for creating inflation control params",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content, err := parseCreateInflationControlParamsArgsToContent(cast.ToString(from), args)
			if err != nil {
				return err
			}

			depositCoinsDetail, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositCoinsDetail)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func NewSubmitUpdateInflationControlParamsProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-inflation-control-params-proposal [title] [description] [lambda] [w360] [w180] [w90] [w30] [w14] [w7] [w1]",
		Short: "Submit a new proposal for updating inflation control params",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content, err := parseUpdateInflationControlParamsArgsToContent(cast.ToString(from), args)
			if err != nil {
				return err
			}

			depositCoinsDetail, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositCoinsDetail)
			if err != nil {
				return err
			}

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")

	return cmd
}

func parseCreateInflationControlParamsArgsToContent(from string, args []string) (govtypes.Content, error) {
	title := args[0]

	description := args[1]

	lambda, err := math.LegacyNewDecFromStr(args[2])
	if err != nil {
		return nil, err
	}
	// Get value arguments
	w360, err := math.LegacyNewDecFromStr(args[3])
	if err != nil {
		return nil, err
	}
	w180, err := math.LegacyNewDecFromStr(args[4])
	if err != nil {
		return nil, err
	}
	w90, err := math.LegacyNewDecFromStr(args[5])
	if err != nil {
		return nil, err
	}
	w30, err := math.LegacyNewDecFromStr(args[6])
	if err != nil {
		return nil, err
	}
	w14, err := math.LegacyNewDecFromStr(args[7])
	if err != nil {
		return nil, err
	}
	w7, err := math.LegacyNewDecFromStr(args[8])
	if err != nil {
		return nil, err
	}
	w1, err := math.LegacyNewDecFromStr(args[9])
	if err != nil {
		return nil, err
	}

	if !(w360.Add(w180.Add(w90.Add(w30.Add(w14.Add(w7.Add(w1)))))).Equal(math.LegacyNewDec(1))) {
		return nil, fmt.Errorf("sum of all weights must be one")
	}
	content := &types.InflationCreateControlParamsProposal{
		Title:       title,
		Description: description,
		ControlParams: &types.InflationControlParams{
			Lambda: lambda,
			W360:   w360,
			W180:   w180,
			W90:    w90,
			W30:    w30,
			W14:    w14,
			W7:     w7,
			W1:     w1,
		},
	}

	return content, nil
}

func parseUpdateInflationControlParamsArgsToContent(from string, args []string) (govtypes.Content, error) {
	title := args[0]

	description := args[1]

	lambda, err := math.LegacyNewDecFromStr(args[2])
	if err != nil {
		return nil, err
	}
	// Get value arguments
	w360, err := math.LegacyNewDecFromStr(args[3])
	if err != nil {
		return nil, err
	}
	w180, err := math.LegacyNewDecFromStr(args[4])
	if err != nil {
		return nil, err
	}
	w90, err := math.LegacyNewDecFromStr(args[5])
	if err != nil {
		return nil, err
	}
	w30, err := math.LegacyNewDecFromStr(args[6])
	if err != nil {
		return nil, err
	}
	w14, err := math.LegacyNewDecFromStr(args[7])
	if err != nil {
		return nil, err
	}
	w7, err := math.LegacyNewDecFromStr(args[8])
	if err != nil {
		return nil, err
	}
	w1, err := math.LegacyNewDecFromStr(args[9])
	if err != nil {
		return nil, err
	}

	if !(w360.Add(w180.Add(w90.Add(w30.Add(w14.Add(w7.Add(w1)))))).Equal(math.LegacyNewDec(1))) {
		return nil, fmt.Errorf("sum of all weights must be one")
	}
	content := &types.InflationUpdateControlParamsProposal{
		Title:       title,
		Description: description,
		ControlParams: &types.InflationControlParams{
			Lambda: lambda,
			W360:   w360,
			W180:   w180,
			W90:    w90,
			W30:    w30,
			W14:    w14,
			W7:     w7,
			W1:     w1,
		},
	}

	return content, nil
}

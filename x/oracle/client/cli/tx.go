package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRequestBandPrice())
	cmd.AddCommand(CmdSubmitAddAssetInfoProposal())
	cmd.AddCommand(CmdSubmitRemoveAssetInfoProposal())
	cmd.AddCommand(CmdSubmitAddPriceFeedersProposal())
	cmd.AddCommand(CmdSubmitRemovePriceFeedersProposal())
	cmd.AddCommand(CmdFeedPrice())
	cmd.AddCommand(CmdSetPriceFeeder())
	cmd.AddCommand(CmdDeletePriceFeeder())
	cmd.AddCommand(CmdFeedMultiplePrices())
	cmd.AddCommand(CmdUpdateChannel())
	cmd.AddCommand(CmdUpdatePriceStatistics())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdSubmitAddAssetInfoProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-asset-info-proposal [denom] [display] [bandTicker] [cascadiaTicker]",
		Args:  cobra.ExactArgs(4),
		Short: "Submit an add asset info proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			content := types.NewProposalAddAssetInfo(
				title,
				description,
				args[0],
				args[1],
				args[2],
				args[3],
			)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := v1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSubmitRemoveAssetInfoProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-asset-info-proposal [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an add asset info proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			content := types.NewProposalRemoveAssetInfo(
				title,
				description,
				args[0],
			)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := v1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSubmitAddPriceFeedersProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-price-feeders-proposal [feeders]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an add price feeders proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			feeders := strings.Split(args[0], ",")
			content := types.NewProposalAddPriceFeeders(
				title,
				description,
				feeders,
			)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := v1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSubmitRemovePriceFeedersProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-price-feeders-proposal [feeders]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a remove price feeders proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			title, err := cmd.Flags().GetString(cli.FlagTitle)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(cli.FlagDescription)
			if err != nil {
				return err
			}

			feeders := strings.Split(args[0], ",")
			content := types.NewProposalRemovePriceFeeders(
				title,
				description,
				feeders,
			)

			from := clientCtx.GetFromAddress()

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := v1beta1.NewMsgSubmitProposal(content, deposit, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewSubmitUpdatePriceFeederInfoProposalTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pricefeeder-info-proposal [title] [description] [name] [address] [active]",
		Short: "Submit a Update pricefeeder info proposal",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress()

			content, err := parsePriceFeederInfoArgsToContent(cast.ToString(from), args)
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

func parsePriceFeederInfoArgsToContent(from string, args []string) (govtypes.Content, error) {
	title := args[0]

	description := args[1]

	name := args[2]

	address := args[3]

	isActive, err := cast.ToBoolE(args[4])
	if err != nil {
		return nil, err
	}

	priceFeedInfo := types.PriceFeederInfo{
		Name:    name,
		Address: address,
		Active:  isActive,
	}
	content := &types.UpdatePriceFeederInfoProposal{
		Title:           title,
		Description:     description,
		PriceFeederInfo: priceFeedInfo,
	}

	return content, nil
}

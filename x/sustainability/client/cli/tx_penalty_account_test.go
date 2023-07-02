package cli_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/cascadiafoundation/github.com/cascadiafoundation/cascadia/testutil/network"

	"github.com/cascadiafoundation/cascadia/x/sustainability/client/cli"
)

func TestCreatePenaltyAccount(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		code uint32
	}{
		{
			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreatePenaltyAccount(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestUpdatePenaltyAccount(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	var args []string
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreatePenaltyAccount(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc string
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			args: common,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdatePenaltyAccount(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

func TestDeletePenaltyAccount(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	var args []string
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreatePenaltyAccount(), args)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc string
		args []string
		code uint32
		err  error
	}{
		{
			desc: "valid",
			args: common,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeletePenaltyAccount(), append([]string{}, tc.args...))
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)
			}
		})
	}
}

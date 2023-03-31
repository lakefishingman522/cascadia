package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/cascadiafoundation/cascadia/app"
	cascadiad "github.com/cascadiafoundation/cascadia/cmd/cascadiad"
	"github.com/cascadiafoundation/cascadia/utils"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := cascadiad.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",          // Test the init cmd
		"cascadia-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestnetChainID+"-1"),
	})

	err := svrcmd.Execute(rootCmd, "cascadiad", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := cascadiad.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "cascadiad", app.DefaultNodeHome)
	require.Error(t, err)
}

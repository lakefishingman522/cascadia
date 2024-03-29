package oracle_test

import (
	"testing"

	"github.com/cascadiafoundation/cascadia/utils"

	simapp "github.com/cascadiafoundation/cascadia/app"
	"github.com/cascadiafoundation/cascadia/testutil/nullify"
	feemarkettypes "github.com/cascadiafoundation/cascadia/x/feemarket/types"
	"github.com/cascadiafoundation/cascadia/x/oracle"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	initChain = true
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		AssetInfos: []types.AssetInfo{
			{
				Denom:   "satoshi",
				Display: "BTC",
			},
			{
				Denom:   "wei",
				Display: "ETH",
			},
		},
		Prices: []types.Price{
			{
				Asset: "BTC",
				Price: sdk.NewDec(30000),
			},
			{
				Asset: "ETH",
				Price: sdk.NewDec(2000),
			},
		},
		PriceFeeders: []types.PriceFeeder{
			{
				Feeder:   "elys10d07y265gmmuvt4z0w9aw880jnsr700j6z2zm3",
				IsActive: true,
			},
			{
				Feeder:   "elys12tzylat4udvjj56uuhu3vj2n4vgp7cf9fwna9w",
				IsActive: false,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	isCheckTx := false

	chainID := utils.TestnetChainID + "-1"

	app := simapp.Setup(isCheckTx, feemarkettypes.DefaultGenesisState(), chainID)
	ctx := app.BaseApp.NewContext(initChain, tmproto.Header{})
	oracle.InitGenesis(ctx, app.OracleKeeper, genesisState)
	got := oracle.ExportGenesis(ctx, app.OracleKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.AssetInfos, got.AssetInfos)
	require.ElementsMatch(t, genesisState.Prices, got.Prices)
	require.ElementsMatch(t, genesisState.PriceFeeders, got.PriceFeeders)
	// this line is used by starport scaffolding # genesis/test/assert
}

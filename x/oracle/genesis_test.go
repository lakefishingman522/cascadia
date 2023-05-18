package oracle_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simapp "github.com/cascadiafoundation/cascadia/app"
	"github.com/cascadiafoundation/cascadia/testutil/nullify"
	"github.com/cascadiafoundation/cascadia/x/oracle"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
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
				Feeder:   "cascadia1uvlapzqatexsw0erk43lpy2ye0vdkkr2rzc7nj",
				IsActive: true,
			},
			{
				Feeder:   "cascadia1x8lscy3egp6rxmm87dzsu7wtd5rjf6myv8vux3",
				IsActive: false,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	app := simapp.InitcascadiaTestApp(initChain)
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

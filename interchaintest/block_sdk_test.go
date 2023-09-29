package interchain_test

import (
	"fmt"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/cascadiafoundation/cascadia/encoding"
	"github.com/cascadiafoundation/cascadia/app"
	"github.com/skip-mev/block-sdk/tests/integration"
	"github.com/stretchr/testify/suite"
	ictestutil "github.com/strangelove-ventures/interchaintest/v7/testutil"
	cascadiakeyring "github.com/cascadiafoundation/cascadia/crypto/keyring"
)

var (
	// config params
	numValidators = 4
	numFullNodes = 0
	denom         = "aCC"

	image = ibc.DockerImage{
		Repository: "tharsishq/cascadia",
		Version:    "latest",
		UidGid:     "1000:1000",
	}
	noHostMount    = false
	gasAdjustment  = float64(2.0)
	encodingConfig = MakeEncodingConfig()

	genesisKV = []cosmos.GenesisKV{
		{
			Key:   "app_state.auction.params.max_bundle_size",
			Value: 3,
		},
		{
			Key:   "app_state.auction.params.reserve_fee.denom",
			Value: denom,
		},
		{
			Key:   "app_state.auction.params.reserve_fee.amount",
			Value: "1",
		},
		{
			Key:   "app_state.auction.params.min_bid_increment.denom",
			Value: denom,
		},
		{
			Key:  "app_state.staking.params.bond_denom",
			Value: denom,
		},
		{
			Key: "app_state.crisis.constant_fee.denom",
			Value: denom,
		},
		{
			Key: "app_state.feemarket.params.no_base_fee",
			Value: true,
		},
		{
			Key: "consensus_params.max_gas",
			Value: -1,
		},
	}

	initCoins = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	// interchain specification
	spec = &interchaintest.ChainSpec{
		ChainName:     "cascadia",
		Name:          "cascadia",
		NumValidators: &numValidators,
		NumFullNodes:  &numFullNodes,
		Version:       "latest",
		NoHostMount:   &noHostMount,
		ChainConfig: ibc.ChainConfig{
			EncodingConfig: encodingConfig,
			Images: []ibc.DockerImage{
				image,
			},
			Type:                   "cosmos",
			Name:                   "cascadia",
			Denom:                  denom,
			ChainID:                "cascadia_9001-1",
			Bin:                    "cascadiad",
			Bech32Prefix:           "cascadia",
			CoinType:               "118",
			GasAdjustment:          gasAdjustment,
			GasPrices:              fmt.Sprintf("0%s", denom),
			TrustingPeriod:         "48h",
			NoHostMount:            noHostMount,
			ModifyGenesis:          cosmos.ModifyGenesis(genesisKV),
			ModifyGenesisAmounts: func() (sdk.Coin, sdk.Coin) {
				return sdk.NewCoin(denom, sdk.NewIntFromBigInt(initCoins)), sdk.NewCoin(denom, sdk.NewIntFromBigInt(initCoins))
			},
			ConfigFileOverrides: map[string]any{
				"config/client.toml" : ictestutil.Toml{
					"chain-id": "cascadia_9001-1",
				},
			},
			UsingChainIDFlagCLI: true,
		},
	}
)

func MakeEncodingConfig() *testutil.TestEncodingConfig {
	ec := encoding.MakeConfig(app.ModuleBasics)
	return &testutil.TestEncodingConfig{
		InterfaceRegistry: ec.InterfaceRegistry,
		Codec: ec.Codec,
		TxConfig: ec.TxConfig,
		Amino: ec.Amino,
	}
}

func TestBlockSDKSuite(t *testing.T) {
	s := integration.NewIntegrationTestSuiteFromSpec(spec)
	s.WithDenom(denom)
	s.WithKeyringOptions(encodingConfig.Codec, cascadiakeyring.Option())
	suite.Run(t, s)
}
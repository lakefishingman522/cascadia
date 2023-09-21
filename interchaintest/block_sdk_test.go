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
)

var (
	// config params
	numValidators = 4
	numFullNodes = 0
	denom         = "inj"

	image = ibc.DockerImage{
		Repository: "gcr.io/injective-core/core",
		Version:    "latest",
		UidGid:     "1000:1000",
	}
	noHostMount    = false
	gasAdjustment  = float64(2.0)
	encodingConfig = MakeEncodingConfig()

	genesisKV = []cosmos.GenesisKV{
		{
			Key:   "app_state.builder.params.max_bundle_size",
			Value: 3,
		},
		{
			Key:   "app_state.builder.params.reserve_fee.denom",
			Value: denom,
		},
		{
			Key:   "app_state.builder.params.reserve_fee.amount",
			Value: "1",
		},
		{
			Key:   "app_state.builder.params.min_bid_increment.denom",
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
	}

	initCoins = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	initDelegation = new(big.Int).Exp(big.NewInt(11), big.NewInt(18), nil) 
	// interchain specification
	spec = &interchaintest.ChainSpec{
		ChainName:     "cascadia",
		Name:          "injective",
		NumValidators: &numValidators,
		NumFullNodes:  &numFullNodes,
		Version:       "latest",
		NoHostMount:   &noHostMount,
		GasAdjustment: &gasAdjustment,
		ChainConfig: ibc.ChainConfig{
			EncodingConfig: encodingConfig,
			Images: []ibc.DockerImage{
				image,
			},
			Type:                   "cosmos",
			Name:                   "injective",
			Denom:                  denom,
			ChainID:                "injective-1",
			Bin:                    "injectived",
			Bech32Prefix:           "inj",
			CoinType:               "118",
			GasAdjustment:          gasAdjustment,
			GasPrices:              fmt.Sprintf("0%s", denom),
			TrustingPeriod:         "48h",
			NoHostMount:            noHostMount,
			UsingNewGenesisCommand: false,
			ModifyGenesis:          cosmos.ModifyGenesis(genesisKV),
			ModifyGenesisAmounts: func() (sdk.Coin, sdk.Coin) {
				return sdk.NewCoin(denom, sdk.NewIntFromBigInt(initCoins)), sdk.NewCoin(denom, sdk.NewIntFromBigInt(initCoins))
			},
			UsingChainIDFlagCLI: true,
		},
	}
)

func MakeEncodingConfig() *testutil.TestEncodingConfig {
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
}
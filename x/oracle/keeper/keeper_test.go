package keeper_test

import (
	"testing"

	"github.com/cascadiafoundation/cascadia/utils"

	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	simapp "github.com/cascadiafoundation/cascadia/app"
	feemarkettypes "github.com/cascadiafoundation/cascadia/x/feemarket/types"
)

const (
	initChain = true
)

type KeeperTestSuite struct {
	suite.Suite

	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	app         *simapp.Cascadia
}

func (suite *KeeperTestSuite) SetupTest() {
	isCheckTx := false

	chainID := utils.TestnetChainID + "-1"

	app := simapp.Setup(isCheckTx, feemarkettypes.DefaultGenesisState(), chainID)

	suite.legacyAmino = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(initChain, tmproto.Header{})
	suite.app = app
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

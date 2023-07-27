package keeper_test

import (
	"strconv"

	"github.com/cascadiafoundation/cascadia/testutil/nullify"
	"github.com/cascadiafoundation/cascadia/x/oracle/keeper"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPriceFeeder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.PriceFeeder {
	items := make([]types.PriceFeeder, n)
	for i := range items {
		items[i].Feeder = strconv.Itoa(i)

		keeper.SetPriceFeeder(ctx, items[i])
	}
	return items
}

func (suite *KeeperTestSuite) TestPriceFeederGet() {
	keeper, ctx := suite.app.OracleKeeper, suite.ctx
	items := createNPriceFeeder(&keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPriceFeeder(ctx, item.Feeder)
		suite.Require().True(found)
		suite.Require().Equal(
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func (suite *KeeperTestSuite) TestPriceFeederRemove() {
	keeper, ctx := suite.app.OracleKeeper, suite.ctx
	items := createNPriceFeeder(&keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePriceFeeder(ctx, item.Feeder)
		_, found := keeper.GetPriceFeeder(ctx, item.Feeder)
		suite.Require().False(found)
	}
}

func (suite *KeeperTestSuite) TestPriceFeederGetAll() {
	keeper, ctx := suite.app.OracleKeeper, suite.ctx
	items := createNPriceFeeder(&keeper, ctx, 10)
	suite.Require().ElementsMatch(
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPriceFeeder(ctx)),
	)
}

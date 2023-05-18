package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
)

func (suite *KeeperTestSuite) TestParamsQuery() {
	keeper, ctx := suite.app.OracleKeeper, suite.ctx
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(&types.QueryParamsResponse{Params: params}, response)
}

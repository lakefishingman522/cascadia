package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"

	"github.com/cascadiafoundation/cascadia/x/sustainability/keeper"
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SustainabilityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

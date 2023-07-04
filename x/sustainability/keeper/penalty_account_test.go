package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"

	"github.com/cascadiafoundation/cascadia/testutil/nullify"

	"github.com/cascadiafoundation/cascadia/x/sustainability/keeper"
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
)

func createTestPenaltyAccount(keeper *keeper.Keeper, ctx sdk.Context) types.PenaltyAccount {
	item := types.PenaltyAccount{}
	keeper.SetPenaltyAccount(ctx, item)
	return item
}

func TestPenaltyAccountGet(t *testing.T) {
	keeper, ctx := keepertest.SustainabilityKeeper(t)
	item := createTestPenaltyAccount(keeper, ctx)
	rst, found := keeper.GetPenaltyAccount(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPenaltyAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.SustainabilityKeeper(t)
	createTestPenaltyAccount(keeper, ctx)
	keeper.RemovePenaltyAccount(ctx)
	_, found := keeper.GetPenaltyAccount(ctx)
	require.False(t, found)
}

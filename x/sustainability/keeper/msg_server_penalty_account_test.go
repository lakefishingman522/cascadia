package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"

	"github.com/cascadiafoundation/cascadia/x/sustainability/keeper"
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
)

func TestPenaltyAccountMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.SustainabilityKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreatePenaltyAccount{Creator: creator}
	_, err := srv.CreatePenaltyAccount(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetPenaltyAccount(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestPenaltyAccountMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdatePenaltyAccount
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePenaltyAccount{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePenaltyAccount{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SustainabilityKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreatePenaltyAccount{Creator: creator}
			_, err := srv.CreatePenaltyAccount(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePenaltyAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPenaltyAccount(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPenaltyAccountMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeletePenaltyAccount
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePenaltyAccount{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePenaltyAccount{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.SustainabilityKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreatePenaltyAccount(wctx, &types.MsgCreatePenaltyAccount{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeletePenaltyAccount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPenaltyAccount(ctx)
				require.False(t, found)
			}
		})
	}
}

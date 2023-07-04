package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PenaltyAccount(goCtx context.Context, req *types.QueryGetPenaltyAccountRequest) (*types.QueryGetPenaltyAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPenaltyAccount(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPenaltyAccountResponse{PenaltyAccount: val}, nil
}

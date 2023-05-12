package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/penalty/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

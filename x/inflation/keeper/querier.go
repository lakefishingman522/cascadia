package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) InflationRate(c context.Context, _ *types.QueryInflationRateRequest) (*types.QueryInflationRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryInflationRateResponse{InflationRate: minter.Inflation}, nil
}

func (k Keeper) AnnualProvisions(c context.Context, _ *types.QueryAnnualProvisionsRequest) (*types.QueryAnnualProvisionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryAnnualProvisionsResponse{AnnualProvisions: minter.AnnualProvisions}, nil
}

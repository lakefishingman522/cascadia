package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/reward/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RewardContractAll(c context.Context, req *types.QueryAllRewardContractRequest) (*types.QueryAllRewardContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rewardContracts []types.RewardContract
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	rewardContractStore := prefix.NewStore(store, types.KeyPrefix(types.RewardContractKey))

	pageRes, err := query.Paginate(rewardContractStore, req.Pagination, func(key []byte, value []byte) error {
		var rewardContract types.RewardContract
		if err := k.cdc.Unmarshal(value, &rewardContract); err != nil {
			return err
		}

		rewardContracts = append(rewardContracts, rewardContract)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRewardContractResponse{RewardContract: rewardContracts, Pagination: pageRes}, nil
}

func (k Keeper) RewardContract(c context.Context, req *types.QueryGetRewardContractRequest) (*types.QueryGetRewardContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	rewardContract, found := k.GetRewardContract(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRewardContractResponse{RewardContract: rewardContract}, nil
}

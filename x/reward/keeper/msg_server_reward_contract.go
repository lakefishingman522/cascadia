package keeper

import (
	"context"
	"fmt"

	"github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRewardContract(goCtx context.Context, msg *types.MsgCreateRewardContract) (*types.MsgCreateRewardContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var rewardContract = types.RewardContract{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	id := k.AppendRewardContract(
		ctx,
		rewardContract,
	)

	return &types.MsgCreateRewardContractResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRewardContract(goCtx context.Context, msg *types.MsgUpdateRewardContract) (*types.MsgUpdateRewardContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var rewardContract = types.RewardContract{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetRewardContract(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetRewardContract(ctx, rewardContract)

	return &types.MsgUpdateRewardContractResponse{}, nil
}

func (k msgServer) DeleteRewardContract(goCtx context.Context, msg *types.MsgDeleteRewardContract) (*types.MsgDeleteRewardContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetRewardContract(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRewardContract(ctx, msg.Id)

	return &types.MsgDeleteRewardContractResponse{}, nil
}

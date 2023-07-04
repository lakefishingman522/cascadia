package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePenaltyAccount(goCtx context.Context, msg *types.MsgCreatePenaltyAccount) (*types.MsgCreatePenaltyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetPenaltyAccount(ctx)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var penaltyAccount = types.PenaltyAccount{
		Creator:         msg.Creator,
		MultisigAddress: msg.MultisigAddress,
	}

	k.SetPenaltyAccount(
		ctx,
		penaltyAccount,
	)
	return &types.MsgCreatePenaltyAccountResponse{}, nil
}

func (k msgServer) UpdatePenaltyAccount(goCtx context.Context, msg *types.MsgUpdatePenaltyAccount) (*types.MsgUpdatePenaltyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPenaltyAccount(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var penaltyAccount = types.PenaltyAccount{
		Creator:         msg.Creator,
		MultisigAddress: msg.MultisigAddress,
	}

	k.SetPenaltyAccount(ctx, penaltyAccount)

	return &types.MsgUpdatePenaltyAccountResponse{}, nil
}

func (k msgServer) DeletePenaltyAccount(goCtx context.Context, msg *types.MsgDeletePenaltyAccount) (*types.MsgDeletePenaltyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPenaltyAccount(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePenaltyAccount(ctx)

	return &types.MsgDeletePenaltyAccountResponse{}, nil
}

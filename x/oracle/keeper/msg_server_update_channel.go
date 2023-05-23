package keeper

import (
	"context"

	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Update source channel
func (k msgServer) UpdateChannel(goCtx context.Context, msg *types.MsgUpdateChannel) (*types.MsgUpdateChannelResponse, error) {
	if len(msg.Channel) < 1 {
		return &types.MsgUpdateChannelResponse{}, nil
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	// Update channel name
	params.BandChannelSource = msg.Channel
	k.SetParams(ctx, params)

	return &types.MsgUpdateChannelResponse{}, nil
}

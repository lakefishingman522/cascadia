package reward

import (
	"fmt"

	"github.com/cascadiafoundation/cascadia/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgCreateRewardContract:
			res, err := msgServer.CreateRewardContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgUpdateRewardContract:
			res, err := msgServer.UpdateRewardContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types.MsgDeleteRewardContract:
			res, err := msgServer.DeleteRewardContract(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

			// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

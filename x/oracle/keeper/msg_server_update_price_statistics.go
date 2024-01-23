package keeper

import (
	"context"
	"fmt"

	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePriceStatistics(goCtx context.Context, msg *types.MsgUpdatePriceStatistics) (*types.MsgUpdatePriceStatisticsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	feeder, found := k.GetPriceFeederInfo(ctx, "inflation")
	if !found {
		return nil, fmt.Errorf("feeder not found")
	}

	if msg.Creator != feeder.Address {
		return nil, fmt.Errorf("Unauthorized oracle provider")
	}

	// TODO: Handling the message
	priceStatistics, found := k.GetPriceStatistics(ctx)

	if !found {
		priceStatistics = types.DefaultPriceStatistics()
	}

	priceStatistics.P360 = sdk.MustNewDecFromStr(msg.P360)
	priceStatistics.P180 = sdk.MustNewDecFromStr(msg.P180)
	priceStatistics.P90 = sdk.MustNewDecFromStr(msg.P90)
	priceStatistics.P30 = sdk.MustNewDecFromStr(msg.P30)
	priceStatistics.P14 = sdk.MustNewDecFromStr(msg.P14)
	priceStatistics.P7 = sdk.MustNewDecFromStr(msg.P7)
	priceStatistics.P1 = sdk.MustNewDecFromStr(msg.P1)

	// Update Price Statistics
	k.SetPriceStatistics(ctx, priceStatistics)

	return &types.MsgUpdatePriceStatisticsResponse{}, nil
}

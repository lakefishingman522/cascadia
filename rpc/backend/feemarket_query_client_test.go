package backend

import (
	"github.com/cascadiafoundation/cascadia/rpc/backend/mocks"
	rpc "github.com/cascadiafoundation/cascadia/rpc/types"
	feemarkettypes "github.com/cascadiafoundation/cascadia/x/feemarket/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ feemarkettypes.QueryClient = &mocks.FeeMarketQueryClient{}

// Params
func RegisterFeeMarketParams(feeMarketClient *mocks.FeeMarketQueryClient, height int64) {
	feeMarketClient.On("Params", rpc.ContextWithHeight(height), &feemarkettypes.QueryParamsRequest{}).
		Return(&feemarkettypes.QueryParamsResponse{Params: feemarkettypes.DefaultParams()}, nil)
}

func RegisterFeeMarketParamsError(feeMarketClient *mocks.FeeMarketQueryClient, height int64) {
	feeMarketClient.On("Params", rpc.ContextWithHeight(height), &feemarkettypes.QueryParamsRequest{}).
		Return(nil, sdkerrors.ErrInvalidRequest)
}

package inflation

import (
	"github.com/cascadiafoundation/cascadia/x/inflation/keeper"
	"github.com/cascadiafoundation/cascadia/x/inflation/keeper/gov"
	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewInflationControlParamsProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.InflationCreateControlParamsProposal:
			return handleInflationCreateControlParamsProposal(ctx, k, c)
		case *types.InflationUpdateControlParamsProposal:
			return handleInflationUpdateControlParamsProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized inflation config proposal type: %T", c)
		}
	}
}

func handleInflationCreateControlParamsProposal(ctx sdk.Context, k keeper.Keeper, p *types.InflationCreateControlParamsProposal) error {
	return gov.HandleCreateInflationControlParamsProposal(ctx, k, p)
}

func handleInflationUpdateControlParamsProposal(ctx sdk.Context, k keeper.Keeper, p *types.InflationUpdateControlParamsProposal) error {
	return gov.HandleUpdateInflationControlParamsProposal(ctx, k, p)
}

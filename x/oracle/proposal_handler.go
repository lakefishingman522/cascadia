package oracle

import (
	"github.com/cascadiafoundation/cascadia/x/oracle/keeper"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func NewAssetInfoProposalHandler(k *keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ProposalAddAssetInfo:
			return handleAddAssetInfoProposal(ctx, k, c)

		case *types.ProposalRemoveAssetInfo:
			return handleRemoveAssetInfoProposal(ctx, k, c)

		case *types.ProposalAddPriceFeeders:
			return handleAddPriceFeedersProposal(ctx, k, c)

		case *types.ProposalRemovePriceFeeders:
			return handleRemovePriceFeedersProposal(ctx, k, c)

		case *types.UpdatePriceFeederInfoProposal:
			return handlePriceFeederInfoProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized software upgrade proposal content type: %T", c)
		}
	}
}

func handleAddAssetInfoProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ProposalAddAssetInfo) error {
	k.SetAssetInfo(ctx, types.AssetInfo{
		Denom:          p.Denom,
		Display:        p.Display,
		BandTicker:     p.BandTicker,
		CascadiaTicker: p.CascadiaTicker,
	})
	return nil
}

func handleRemoveAssetInfoProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ProposalRemoveAssetInfo) error {
	k.RemoveAssetInfo(ctx, p.Denom)
	return nil
}

func handleAddPriceFeedersProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ProposalAddPriceFeeders) error {
	for _, feeder := range p.Feeders {
		k.SetPriceFeeder(ctx, types.PriceFeeder{
			Feeder:   feeder,
			IsActive: true,
		})
	}
	return nil
}

func handleRemovePriceFeedersProposal(ctx sdk.Context, k *keeper.Keeper, p *types.ProposalRemovePriceFeeders) error {
	for _, feeder := range p.Feeders {
		k.RemovePriceFeeder(ctx, feeder)
	}
	return nil
}

// handleUpdateSymbolRequestProposal is a function that handles the update symbol request proposal.
func handlePriceFeederInfoProposal(ctx sdk.Context, k *keeper.Keeper, p *types.UpdatePriceFeederInfoProposal) error {
	// Set the symbol requests in the keeper with the new symbol requests specified in the proposal.
	k.SetPriceFeederInfo(ctx, p.PriceFeederInfo)
	return nil
}

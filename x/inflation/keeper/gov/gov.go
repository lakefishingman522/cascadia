package gov

import (
	"github.com/cascadiafoundation/cascadia/x/inflation/keeper"
	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HandleCreateInflationControlParamsProposal(ctx sdk.Context, k keeper.Keeper, p *types.InflationCreateControlParamsProposal) error {

	// configuration to be added
	var msg = p.ControlParams

	// Check if the value already exists
	_, isFound := k.GetInflationControlParams(
		ctx,
	)

	if isFound {
		return nil
	}

	var inflationControlParams = types.InflationControlParams{
		Lambda: msg.Lambda,
		W360:   msg.W360,
		W180:   msg.W180,
		W90:    msg.W90,
		W30:    msg.W30,
		W14:    msg.W14,
		W7:     msg.W7,
		W1:     msg.W1,
	}

	k.SetInflationControlParams(
		ctx,
		inflationControlParams,
	)

	return nil
}

func HandleUpdateInflationControlParamsProposal(ctx sdk.Context, k keeper.Keeper, p *types.InflationUpdateControlParamsProposal) error {

	// configuration to be added
	var msg = p.ControlParams

	// Check if the value already exists
	_, isFound := k.GetInflationControlParams(
		ctx,
	)

	if !isFound {
		return nil
	}

	var inflationControlParams = types.InflationControlParams{
		Lambda: msg.Lambda,
		W360:   msg.W360,
		W180:   msg.W180,
		W90:    msg.W90,
		W30:    msg.W30,
		W14:    msg.W14,
		W7:     msg.W7,
		W1:     msg.W1,
	}

	k.SetInflationControlParams(
		ctx,
		inflationControlParams,
	)

	return nil
}

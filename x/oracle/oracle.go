package oracle

import (
	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

// handleOraclePacket handles the result of the received BandChain oracles
// packet and saves the data into the KV database
func (im IBCModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData packet.OracleResponsePacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}

	switch modulePacketData.GetClientID() {

	case types.BandPriceClientIDKey:
		var OraclePriceResult types.OraclePriceResults_UInt8Version
		if err := obi.Decode(modulePacketData.Result, &OraclePriceResult); err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "cannot decode the BandPrice received packet")
		}
		reqID := types.OracleRequestID(modulePacketData.RequestID)

		var BandPriceResult types.BandPriceResult
		for _, r := range OraclePriceResult.Responses {
			BandPriceResult.Rates = append(BandPriceResult.Rates, r.Rate)
		}

		im.keeper.SetBandPriceResult(ctx, reqID, BandPriceResult)

		params := im.keeper.GetParams(ctx)
		request, err := im.keeper.GetBandRequest(ctx, reqID)
		if err != nil {
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "historical request does not exist")
		}

		if len(request.Symbols) != len(BandPriceResult.Rates) {
			return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "request and result count does not match")
		}

		for index, symbol := range request.Symbols {
			im.keeper.SetPrice(ctx, types.Price{
				Asset:     symbol,
				Price:     sdk.NewDecWithPrec(int64(BandPriceResult.Rates[index]), int64(params.Multiplier)),
				Source:    types.BAND,
				Provider:  "automation",
				Timestamp: uint64(ctx.BlockTime().Unix()),
			})
		}
		// this line is used by starport scaffolding # oracle/module/recv

	default:
		err := sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "oracle received packet not found: %s", modulePacketData.GetClientID())
		ack = channeltypes.NewErrorAcknowledgement(err)
		return ack, err

	}
	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			packet.NewOracleRequestPacketAcknowledgement(modulePacketData.RequestID),
		),
	)
	return ack, nil
}

// handleOracleAcknowledgment handles the acknowledgment result from the BandChain
// request and saves the request-id into the KV database
func (im IBCModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck packet.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		var data packet.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}
		requestID := types.OracleRequestID(oracleAck.RequestID)

		switch data.GetClientID() {

		case types.BandPriceClientIDKey:
			var RequestBandPrice types.Calldata
			if err = obi.Decode(data.GetCalldata(), &RequestBandPrice); err != nil {
				return nil, sdkerrors.Wrap(err, "cannot decode the BandPrice oracle acknowledgment packet")
			}
			im.keeper.SetLastBandRequestId(ctx, requestID)

			RequestBandPrice_ := types.BandPriceCallData{
				Symbols:    RequestBandPrice.Symbols,
				Multiplier: uint64(RequestBandPrice.MinimumSourceCount),
			}
			im.keeper.SetBandRequest(ctx, requestID, RequestBandPrice_)
			return &sdk.Result{}, nil
			// this line is used by starport scaffolding # oracle/module/ack

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"oracle acknowledgment packet not found: %s", data.GetClientID())
		}
	}
	return nil, nil
}

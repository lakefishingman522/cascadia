package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePriceStatistics = "update_price_statistics"

var _ sdk.Msg = &MsgUpdatePriceStatistics{}

func NewMsgUpdatePriceStatistics(creator string, p360 string, p180 string, p90 string, p30 string, p14 string, p7 string, p1 string) *MsgUpdatePriceStatistics {
  return &MsgUpdatePriceStatistics{
		Creator: creator,
    P360: p360,
    P180: p180,
    P90: p90,
    P30: p30,
    P14: p14,
    P7: p7,
    P1: p1,
	}
}

func (msg *MsgUpdatePriceStatistics) Route() string {
  return RouterKey
}

func (msg *MsgUpdatePriceStatistics) Type() string {
  return TypeMsgUpdatePriceStatistics
}

func (msg *MsgUpdatePriceStatistics) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePriceStatistics) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePriceStatistics) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}


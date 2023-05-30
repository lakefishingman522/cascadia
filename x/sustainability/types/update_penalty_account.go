package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgUpdatePenaltyAccount = "update_penalty_account"
)

var _ sdk.Msg = &MsgUpdatePenaltyAccountRequest{}

func NewMsgUpdatePenaltyAccountRequest(oldCreator, newAccount string) *MsgUpdatePenaltyAccountRequest {
	return &MsgUpdatePenaltyAccountRequest{
		Creator:    oldCreator,
		NewAddress: newAccount,
	}
}

func (msg *MsgUpdatePenaltyAccountRequest) Route() string {
	return QuerierRoute
}

func (msg *MsgUpdatePenaltyAccountRequest) Type() string {
	return TypeMsgUpdatePenaltyAccount
}

func (msg *MsgUpdatePenaltyAccountRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePenaltyAccountRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePenaltyAccountRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

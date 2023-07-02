package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePenaltyAccount = "create_penalty_account"
	TypeMsgUpdatePenaltyAccount = "update_penalty_account"
	TypeMsgDeletePenaltyAccount = "delete_penalty_account"
)

var _ sdk.Msg = &MsgCreatePenaltyAccount{}

func NewMsgCreatePenaltyAccount(creator string, multisigAddress string) *MsgCreatePenaltyAccount {
	return &MsgCreatePenaltyAccount{
		Creator:         creator,
		MultisigAddress: multisigAddress,
	}
}

func (msg *MsgCreatePenaltyAccount) Route() string {
	return RouterKey
}

func (msg *MsgCreatePenaltyAccount) Type() string {
	return TypeMsgCreatePenaltyAccount
}

func (msg *MsgCreatePenaltyAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePenaltyAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePenaltyAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePenaltyAccount{}

func NewMsgUpdatePenaltyAccount(creator string, multisigAddress string) *MsgUpdatePenaltyAccount {
	return &MsgUpdatePenaltyAccount{
		Creator:         creator,
		MultisigAddress: multisigAddress,
	}
}

func (msg *MsgUpdatePenaltyAccount) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePenaltyAccount) Type() string {
	return TypeMsgUpdatePenaltyAccount
}

func (msg *MsgUpdatePenaltyAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePenaltyAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePenaltyAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePenaltyAccount{}

func NewMsgDeletePenaltyAccount(creator string) *MsgDeletePenaltyAccount {
	return &MsgDeletePenaltyAccount{
		Creator: creator,
	}
}
func (msg *MsgDeletePenaltyAccount) Route() string {
	return RouterKey
}

func (msg *MsgDeletePenaltyAccount) Type() string {
	return TypeMsgDeletePenaltyAccount
}

func (msg *MsgDeletePenaltyAccount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePenaltyAccount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePenaltyAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

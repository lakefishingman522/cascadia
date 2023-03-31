package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRewardContract = "create_reward_contract"
	TypeMsgUpdateRewardContract = "update_reward_contract"
	TypeMsgDeleteRewardContract = "delete_reward_contract"
)

var _ sdk.Msg = &MsgCreateRewardContract{}

func NewMsgCreateRewardContract(creator string, address string) *MsgCreateRewardContract {
	return &MsgCreateRewardContract{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateRewardContract) Route() string {
	return RouterKey
}

func (msg *MsgCreateRewardContract) Type() string {
	return TypeMsgCreateRewardContract
}

func (msg *MsgCreateRewardContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRewardContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRewardContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRewardContract{}

func NewMsgUpdateRewardContract(creator string, id uint64, address string) *MsgUpdateRewardContract {
	return &MsgUpdateRewardContract{
		Id:      id,
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateRewardContract) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRewardContract) Type() string {
	return TypeMsgUpdateRewardContract
}

func (msg *MsgUpdateRewardContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRewardContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRewardContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRewardContract{}

func NewMsgDeleteRewardContract(creator string, id uint64) *MsgDeleteRewardContract {
	return &MsgDeleteRewardContract{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteRewardContract) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRewardContract) Type() string {
	return TypeMsgDeleteRewardContract
}

func (msg *MsgDeleteRewardContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRewardContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRewardContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

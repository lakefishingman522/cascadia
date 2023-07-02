package types

import (
	"testing"

	"github.com/cascadiafoundation/github.com/cascadiafoundation/cascadia/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePenaltyAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePenaltyAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePenaltyAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePenaltyAccount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdatePenaltyAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePenaltyAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePenaltyAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdatePenaltyAccount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeletePenaltyAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeletePenaltyAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeletePenaltyAccount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeletePenaltyAccount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

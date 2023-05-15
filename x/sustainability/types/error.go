package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/reward module sentinel errors
var (
	ErrInvalidCreator = sdkerrors.Register(ModuleName, 1500, "only original creator can update this address")
)

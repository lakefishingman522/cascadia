package types // noalias

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	SetPenaltyAccount(ctx sdk.Context, address sdk.AccAddress)
	GetPenaltyAccount(ctx sdk.Context) sdk.AccAddress
}

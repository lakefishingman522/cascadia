// Copyright 2022 Cascadia Foundation
// This file is part of the Cascadia Network packages.
//
// Cascadia is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Cascadia packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Cascadia packages. If not, see https://github.com/cascadiafoundation/cascadia/blob/main/LICENSE
package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
)

var _ types.MsgServer = &Keeper{}


func (k *Keeper) UpdatePenaltyAccount(goCtx context.Context, msg *types.MsgUpdatePenaltyAccountRequest) (*types.MsgUpdatePenaltyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentPenaltyAccount := k.stakingKeeper.GetPenaltyAccount(ctx)
	if msg.Creator != currentPenaltyAccount.String() {
		return nil, fmt.Errorf("Failed to update penalty account %s", types.ErrInvalidCreator)
	}

	oldAddr := k.stakingKeeper.GetPenaltyAccount(ctx)
	k.stakingKeeper.SetPenaltyAccount(ctx, sdk.MustAccAddressFromBech32(msg.NewAddress))
	return &types.MsgUpdatePenaltyAccountResponse{
		OldAddress: oldAddr.String(),
	}, nil
}

func (k Keeper) PenaltyAccount(goCtx context.Context, msg *types.QueryPenaltyAccountRequest) (*types.QueryPenaltyAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	acc := k.stakingKeeper.GetPenaltyAccount(ctx)
	return &types.QueryPenaltyAccountResponse{
		Address: acc.String(),
	}, nil
}

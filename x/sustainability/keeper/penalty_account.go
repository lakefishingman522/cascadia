package keeper

import (
	"fmt"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPenaltyAccount set penaltyAccount in the store
func (k Keeper) SetPenaltyAccount(ctx sdk.Context, penaltyAccount types.PenaltyAccount) (err error) {
	multisigAddress, err := sdk.GetFromBech32(penaltyAccount.GetMultisigAddress(), "cascadia")
	if err != nil {
		return err
	}

	found := k.accountKeeper.HasAccount(ctx, multisigAddress)

	if found == false {
		return fmt.Errorf("Can't set non-exist multisigAddress")
	}

	k.stakingKeeper.SetPenaltyAccount(ctx, penaltyAccount)
	return nil
}

// GetPenaltyAccount returns penaltyAccount
func (k Keeper) GetPenaltyAccount(ctx sdk.Context) (val types.PenaltyAccount, found bool) {
	val, found = k.stakingKeeper.GetPenaltyAccount(ctx)
	return val, found
}

// RemovePenaltyAccount removes penaltyAccount from the store
func (k Keeper) RemovePenaltyAccount(ctx sdk.Context) {
	k.stakingKeeper.RemovePenaltyAccount(ctx)
}

package keeper

import (
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPenaltyAccount set penaltyAccount in the store
func (k Keeper) SetPenaltyAccount(ctx sdk.Context, penaltyAccount types.PenaltyAccount) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PenaltyAccountKey))
	b := k.cdc.MustMarshal(&penaltyAccount)
	store.Set([]byte{0}, b)
}

// GetPenaltyAccount returns penaltyAccount
func (k Keeper) GetPenaltyAccount(ctx sdk.Context) (val types.PenaltyAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PenaltyAccountKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePenaltyAccount removes penaltyAccount from the store
func (k Keeper) RemovePenaltyAccount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PenaltyAccountKey))
	store.Delete([]byte{0})
}

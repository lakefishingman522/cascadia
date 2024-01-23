package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
)

// SetPriceFeederInfo set a specific priceFeederInfo in the store from its index
func (k Keeper) SetPriceFeederInfo(ctx sdk.Context, priceFeederInfo types.PriceFeederInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeederInfoKeyPrefix))
	b := k.cdc.MustMarshal(&priceFeederInfo)
	store.Set(types.PriceFeederInfoKey(
		priceFeederInfo.Name,
	), b)
}

// GetPriceFeederInfo returns a priceFeederInfo from its index
func (k Keeper) GetPriceFeederInfo(
	ctx sdk.Context,
	name string,

) (val types.PriceFeederInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeederInfoKeyPrefix))

	b := store.Get(types.PriceFeederInfoKey(
		name,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceFeederInfo removes a priceFeederInfo from the store
func (k Keeper) RemovePriceFeederInfo(
	ctx sdk.Context,
	name string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeederInfoKeyPrefix))
	store.Delete(types.PriceFeederInfoKey(
		name,
	))
}

// GetAllPriceFeederInfo returns all priceFeederInfo
func (k Keeper) GetAllPriceFeederInfo(ctx sdk.Context) (list []types.PriceFeederInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceFeederInfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PriceFeederInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

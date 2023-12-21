package keeper

import (
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPriceStatistics set priceStatistics in the store
func (k Keeper) SetPriceStatistics(ctx sdk.Context, priceStatistics types.PriceStatistics) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceStatisticsKey))
	b := k.cdc.MustMarshal(&priceStatistics)
	store.Set([]byte{0}, b)
}

// GetPriceStatistics returns priceStatistics
func (k Keeper) GetPriceStatistics(ctx sdk.Context) (val types.PriceStatistics, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceStatisticsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePriceStatistics removes priceStatistics from the store
func (k Keeper) RemovePriceStatistics(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PriceStatisticsKey))
	store.Delete([]byte{0})
}

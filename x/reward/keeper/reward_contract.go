package keeper

import (
	"encoding/binary"

	"github.com/cascadiafoundation/cascadia/x/reward/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetRewardContractCount get the total number of rewardContract
func (k Keeper) GetRewardContractCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RewardContractCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRewardContractCount set the total number of rewardContract
func (k Keeper) SetRewardContractCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RewardContractCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRewardContract appends a rewardContract in the store with a new id and update the count
func (k Keeper) AppendRewardContract(
	ctx sdk.Context,
	rewardContract types.RewardContract,
) uint64 {
	// Create the rewardContract
	count := k.GetRewardContractCount(ctx)

	// Set the ID of the appended value
	rewardContract.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardContractKey))
	appendedValue := k.cdc.MustMarshal(&rewardContract)
	store.Set(GetRewardContractIDBytes(rewardContract.Id), appendedValue)

	// Update rewardContract count
	k.SetRewardContractCount(ctx, count+1)

	return count
}

// SetRewardContract set a specific rewardContract in the store
func (k Keeper) SetRewardContract(ctx sdk.Context, rewardContract types.RewardContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardContractKey))
	b := k.cdc.MustMarshal(&rewardContract)
	store.Set(GetRewardContractIDBytes(rewardContract.Id), b)
}

// GetRewardContract returns a rewardContract from its id
func (k Keeper) GetRewardContract(ctx sdk.Context, id uint64) (val types.RewardContract, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardContractKey))
	b := store.Get(GetRewardContractIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRewardContract removes a rewardContract from the store
func (k Keeper) RemoveRewardContract(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardContractKey))
	store.Delete(GetRewardContractIDBytes(id))
}

// GetAllRewardContract returns all rewardContract
func (k Keeper) GetAllRewardContract(ctx sdk.Context) (list []types.RewardContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RewardContractKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RewardContract
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRewardContractIDBytes returns the byte representation of the ID
func GetRewardContractIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRewardContractIDFromBytes returns ID in uint64 format from a byte array
func GetRewardContractIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

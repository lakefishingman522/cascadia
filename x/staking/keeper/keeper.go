package keeper

import (
	"fmt"

	"cosmossdk.io/math"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	sustainabilitytypes "github.com/cascadiafoundation/cascadia/x/sustainability/types"
)

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}

// Implements DelegationSet interface
var _ types.DelegationSet = Keeper{}

// keeper of the staking store
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	authKeeper types.AccountKeeper
	bankKeeper types.BankKeeper
	hooks      types.StakingHooks
	paramstore paramtypes.Subspace
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, ak types.AccountKeeper, bk types.BankKeeper,
	ps paramtypes.Subspace,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	// ensure bonded and not bonded module accounts are set
	if addr := ak.GetModuleAddress(types.BondedPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BondedPoolName))
	}

	if addr := ak.GetModuleAddress(types.NotBondedPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.NotBondedPoolName))
	}

	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		authKeeper: ak,
		bankKeeper: bk,
		paramstore: ps,
		hooks:      nil,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdktypes.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// Set the validator hooks
func (k *Keeper) SetHooks(sh types.StakingHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}

	k.hooks = sh

	return k
}

// Load the last total validator power.
func (k Keeper) GetLastTotalPower(ctx sdktypes.Context) math.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastTotalPowerKey)

	if bz == nil {
		return sdktypes.ZeroInt()
	}

	ip := sdktypes.IntProto{}
	k.cdc.MustUnmarshal(bz, &ip)

	return ip.Int
}

// Set the last total validator power.
func (k Keeper) SetLastTotalPower(ctx sdktypes.Context, power math.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdktypes.IntProto{Int: power})
	store.Set(types.LastTotalPowerKey, bz)
}

// SetPenaltyAccount set penaltyAccount in the store
func (k Keeper) SetPenaltyAccount(ctx sdktypes.Context, penaltyAccount sustainabilitytypes.PenaltyAccount) {
	
	store := prefix.NewStore(ctx.KVStore(k.storeKey), sustainabilitytypes.KeyPrefix(sustainabilitytypes.PenaltyAccountKey))
	b := k.cdc.MustMarshal(&penaltyAccount)
	store.Set([]byte{0}, b)
}

// GetPenaltyAccount returns penaltyAccount
func (k Keeper) GetPenaltyAccount(ctx sdktypes.Context) (val sustainabilitytypes.PenaltyAccount, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), sustainabilitytypes.KeyPrefix(sustainabilitytypes.PenaltyAccountKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	
	return val, true
}

// RemovePenaltyAccount removes penaltyAccount from the store
func (k Keeper) RemovePenaltyAccount(ctx sdktypes.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), sustainabilitytypes.KeyPrefix(sustainabilitytypes.PenaltyAccountKey))
	store.Delete([]byte{0})
}

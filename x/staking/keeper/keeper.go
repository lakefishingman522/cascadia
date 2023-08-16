package keeper

import (
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	sustainabilitytypes "github.com/cascadiafoundation/cascadia/x/sustainability/types"
)

// Implements ValidatorSet interface
var _ types.ValidatorSet = Keeper{}

// Implements DelegationSet interface
var _ types.DelegationSet = Keeper{}

// Keeper of the x/staking store
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	authKeeper types.AccountKeeper
	bankKeeper types.BankKeeper
	hooks      types.StakingHooks
	authority  string
}

// NewKeeper creates a new staking Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	authority string,
) *Keeper {
	// ensure bonded and not bonded module accounts are set
	if addr := ak.GetModuleAddress(types.BondedPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.BondedPoolName))
	}

	if addr := ak.GetModuleAddress(types.NotBondedPoolName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.NotBondedPoolName))
	}

	// ensure that authority is a valid AccAddress
	if _, err := sdktypes.AccAddressFromBech32(authority); err != nil {
		panic("authority is not a valid acc address")
	}

	return &Keeper{
		storeKey:   key,
		cdc:        cdc,
		authKeeper: ak,
		bankKeeper: bk,
		hooks:      nil,
		authority:  authority,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdktypes.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// Hooks gets the hooks for staking *Keeper {
func (k *Keeper) Hooks() types.StakingHooks {
	if k.hooks == nil {
		// return a no-op implementation if no hooks are set
		return types.MultiStakingHooks{}
	}

	return k.hooks
}

// SetHooks Set the validator hooks.  In contrast to other receivers, this method must take a pointer due to nature
// of the hooks interface and SDK start up sequence.
func (k *Keeper) SetHooks(sh types.StakingHooks) {
	if k.hooks != nil {
		panic("cannot set validator hooks twice")
	}

	k.hooks = sh
}

// GetLastTotalPower Load the last total validator power.
func (k Keeper) GetLastTotalPower(ctx sdktypes.Context) math.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastTotalPowerKey)

	if bz == nil {
		return math.ZeroInt()
	}

	ip := sdktypes.IntProto{}
	k.cdc.MustUnmarshal(bz, &ip)

	return ip.Int
}

// SetLastTotalPower Set the last total validator power.
func (k Keeper) SetLastTotalPower(ctx sdktypes.Context, power math.Int) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdktypes.IntProto{Int: power})
	store.Set(types.LastTotalPowerKey, bz)
}

// GetAuthority returns the x/staking module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// SetValidatorUpdates sets the ABCI validator power updates for the current block.
func (k Keeper) SetValidatorUpdates(ctx sdktypes.Context, valUpdates []abci.ValidatorUpdate) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.ValidatorUpdates{Updates: valUpdates})
	store.Set(types.ValidatorUpdatesKey, bz)
}

// GetValidatorUpdates returns the ABCI validator power updates within the current block.
func (k Keeper) GetValidatorUpdates(ctx sdktypes.Context) []abci.ValidatorUpdate {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ValidatorUpdatesKey)

	var valUpdates types.ValidatorUpdates
	k.cdc.MustUnmarshal(bz, &valUpdates)

	return valUpdates.Updates
}

// *****
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

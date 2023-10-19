package keeper

import (
	"cosmossdk.io/math"
	"github.com/cometbft/cometbft/libs/log"

	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	otypes "github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryCodec
	storeKey         storetypes.StoreKey
	paramSpace       paramtypes.Subspace
	stakingKeeper    types.StakingKeeper
	bankKeeper       types.BankKeeper
	rewardKeeper     types.RewardKeeper
	accountKeeper    types.AccountKeeper
	oracleKeeper     types.OracleKeeper
	feeCollectorName string
}

const (
	FeeDistrContract uint64 = 1
	Nprotocol               = 2
)

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	sk types.StakingKeeper, ak types.AccountKeeper, ok types.OracleKeeper,
	bk types.BankKeeper, rk types.RewardKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		stakingKeeper:    sk,
		bankKeeper:       bk,
		rewardKeeper:     rk,
		accountKeeper:    ak,
		oracleKeeper:     ok,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshal(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&minter)
	store.Set(types.MinterKey, b)
}

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// StakingTokenSupply implements an alias call to the underlying staking keeper's
// StakingTokenSupply to be used in BeginBlocker.
func (k Keeper) StakingTokenSupply(ctx sdk.Context) math.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx)
}

// BondedRatio implements an alias call to the underlying staking keeper's
// BondedRatio to be used in BeginBlocker.
func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	return k.stakingKeeper.BondedRatio(ctx)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// AddCollectedFees implements an alias call to the underlying supply keeper's
// AddCollectedFees to be used in BeginBlocker.
func (k Keeper) AddCollectedFees(ctx sdk.Context) error {
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	staking := k.bankKeeper.GetAllBalances(ctx, moduleAddr)

	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, staking)
}

// Allocate block rewards to fee distributor
func (k Keeper) AddContractIncentives(ctx sdk.Context, amount sdk.Coin, contractnumber uint64) (sdk.Coin, error) {
	params := k.GetParams(ctx)

	if contractnumber != FeeDistrContract && contractnumber != Nprotocol {
		return sdk.Coin{Denom: params.MintDenom, Amount: math.ZeroInt()}, nil
	}

	feeDistrContract, contractFound := k.rewardKeeper.GetRewardContract(ctx, contractnumber)
	proportions := params.InflationDistribution

	if contractFound {
		var contractReward sdk.Coins

		if contractnumber == FeeDistrContract {
			contractReward = sdk.NewCoins(k.GetProportions(ctx, amount, proportions.VecontractRewards))
		} else if contractnumber == Nprotocol {
			contractReward = sdk.NewCoins(k.GetProportions(ctx, amount, proportions.NprotocolRewards))
		}

		address, err := sdk.AccAddressFromHexUnsafe(feeDistrContract.Address[2:])
		if err != nil {
			panic(err)
		}

		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, contractReward)
		return amount, nil
	}

	return sdk.Coin{Denom: params.MintDenom, Amount: math.ZeroInt()}, nil
}

// GetAllocationProportion calculates the proportion of coins that is to be
// allocated during inflation for a given distribution.
func (k Keeper) GetProportions(
	ctx sdk.Context,
	coin sdk.Coin,
	distribution sdk.Dec,
) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		sdk.NewDecFromInt(coin.Amount).Mul(distribution).TruncateInt(),
	)
}

func (k Keeper) GetLatestPriceFromAssetAndSource(ctx sdk.Context, asset string, source string) (otypes.Price, bool) {
	return k.oracleKeeper.GetLatestPriceFromAssetAndSource(ctx, asset, source)
}

package types // noalias

import (
	"cosmossdk.io/math"
	otypes "github.com/cascadiafoundation/cascadia/x/oracle/types"
	rewardtypes "github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// StakingKeeper defines the expected staking keeper
type StakingKeeper interface {
	StakingTokenSupply(ctx sdk.Context) math.Int
	BondedRatio(ctx sdk.Context) sdk.Dec
}

// AccountKeeper defines the contract required for account APIs.
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, types.ModuleAccountI)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
}

// BankKeeper defines the contract needed to be fulfilled for banking and supply
// dependencies.
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type RewardKeeper interface {
	GetRewardContract(ctx sdk.Context, id uint64) (val rewardtypes.RewardContract, found bool)
}

type OracleKeeper interface {
	// Fetch latest price from asset and source
	GetLatestPriceFromAssetAndSource(sdk.Context, string, string) (otypes.Price, bool)
	// Fetch latest price from any source
	GetLatestPriceFromAnySource(sdk.Context, string) (otypes.Price, bool)
	// GetPriceStatistics returns priceStatistics
	GetPriceStatistics(ctx sdk.Context) (val otypes.PriceStatistics, found bool)
	// SetPriceStatistics set priceStatistics in the oracle KVstore
	SetPriceStatistics(ctx sdk.Context, priceStatistics otypes.PriceStatistics)
}

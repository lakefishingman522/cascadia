package v0_1_6

import (
	"cosmossdk.io/math"

	stakingkeeper "github.com/cascadiafoundation/cascadia/x/staking/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	auctionkeeper "github.com/skip-mev/block-sdk/x/auction/keeper"

	auctiontypes "github.com/skip-mev/block-sdk/x/auction/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v8
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	// upgrade dependencies
	ak auctionkeeper.Keeper,
	sk stakingkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		logger.Debug("running module migrations (v0.1.6) ...")

		// update the auction module's escrow address
		params := auctiontypes.DefaultGenesisState().GetParams()

		// update the auction module's escrow address
		penalty_account, _ := sk.GetPenaltyAccount(ctx)
		multiSigAddress, err := sdk.GetFromBech32(penalty_account.MultisigAddress, "cascadia")
		if err != nil {
			return nil, err
		}

		params.EscrowAccountAddress = multiSigAddress
		params.MaxBundleSize = 4
		params.FrontRunningProtection = false
		params.MinBidIncrement.Denom = BaseDenom
		params.MinBidIncrement.Amount = math.NewInt(1000000)
		params.ReserveFee.Denom = BaseDenom
		params.ReserveFee.Amount = math.NewInt(1000000)

		// set to state
		if err := ak.SetParams(ctx, params); err != nil {
			return nil, err
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}

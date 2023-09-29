package v0_1_6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	auctionkeeper "github.com/skip-mev/block-sdk/x/auction/keeper"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v8
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	// upgrade dependencies
	ak auctionkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		logger.Debug("running module migrations (v0.1.6) ...")

		// update the auction module's escrow address
		params, err := ak.GetParams(ctx)
		if err != nil {
			return nil, err
		}

		// update the auction module's escrow address
		multiSigAddress, err := sdk.AccAddressFromBech32(EscrowAddress)
		if err != nil {
			return nil, err
		}

		params.EscrowAccountAddress = multiSigAddress

		// set to state
		if err := ak.SetParams(ctx, params); err != nil {
			return nil, err
		}

		return mm.RunMigrations(ctx, configurator, vm)
	}
}

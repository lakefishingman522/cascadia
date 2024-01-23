package v0_3_0

import (
	oraclekeeper "github.com/cascadiafoundation/cascadia/x/oracle/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v8
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	// upgrade dependencies
	oracleKeeper oraclekeeper.Keeper,

) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		logger.Info("running module migrations (v0.3.0) ...")
		oracleKeeper.MigrateAddOracleProviderAddress(ctx)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}

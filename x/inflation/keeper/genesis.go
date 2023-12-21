package keeper

import (
	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new mint genesis
func (keeper Keeper) InitGenesis(ctx sdk.Context, ak types.AccountKeeper, data *types.GenesisState) {
	keeper.SetMinter(ctx, data.Minter)
	keeper.SetParams(ctx, data.Params)
	keeper.SetInflationControlParams(ctx, data.InflationControlParams)
	ak.GetModuleAccount(ctx, types.ModuleName)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (keeper Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	minter := keeper.GetMinter(ctx)
	params := keeper.GetParams(ctx)
	inflation_control_params, _ := keeper.GetInflationControlParams(ctx)
	return types.NewGenesisState(minter, params, inflation_control_params)
}

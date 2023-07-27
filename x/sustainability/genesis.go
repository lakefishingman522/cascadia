package sustainability

import (
	"github.com/cascadiafoundation/cascadia/x/sustainability/keeper"
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.PenaltyAccount != nil {
		k.SetPenaltyAccount(ctx, *genState.PenaltyAccount)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all penaltyAccount
	penaltyAccount, found := k.GetPenaltyAccount(ctx)
	if found {
		genesis.PenaltyAccount = &penaltyAccount
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

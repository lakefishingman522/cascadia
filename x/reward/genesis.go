package reward

import (
	"github.com/cascadiafoundation/cascadia/x/reward/keeper"
	"github.com/cascadiafoundation/cascadia/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the rewardContract
	for _, elem := range genState.RewardContractList {
		k.SetRewardContract(ctx, elem)
	}

	// Set rewardContract count
	k.SetRewardContractCount(ctx, genState.RewardContractCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RewardContractList = k.GetAllRewardContract(ctx)
	genesis.RewardContractCount = k.GetRewardContractCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

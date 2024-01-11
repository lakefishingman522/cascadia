package oracle

import (
	"math/rand"

	"github.com/cascadiafoundation/cascadia/testutil/sample"
	oraclesimulation "github.com/cascadiafoundation/cascadia/x/oracle/simulation"
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = oraclesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
    opWeightMsgUpdatePriceStatistics = "op_weight_msg_update_price_statistics"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePriceStatistics int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	oracleGenesis := types.GenesisState{
		Params:	types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&oracleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgUpdatePriceStatistics int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePriceStatistics, &weightMsgUpdatePriceStatistics, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePriceStatistics = defaultWeightMsgUpdatePriceStatistics
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePriceStatistics,
		oraclesimulation.SimulateMsgUpdatePriceStatistics(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
	    simulation.NewWeightedProposalMsg(
	opWeightMsgUpdatePriceStatistics,
	defaultWeightMsgUpdatePriceStatistics,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		oraclesimulation.SimulateMsgUpdatePriceStatistics(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

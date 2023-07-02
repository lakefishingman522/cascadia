package sustainability_test

import (
	"testing"

	keepertest "github.com/cascadiafoundation/cascadia/testutil/keeper"
	"github.com/cascadiafoundation/cascadia/testutil/nullify"

	"github.com/cascadiafoundation/cascadia/x/sustainability"
	"github.com/cascadiafoundation/cascadia/x/sustainability/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PenaltyAccount: &types.PenaltyAccount{
			MultisigAddress: "99",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SustainabilityKeeper(t)
	sustainability.InitGenesis(ctx, *k, genesisState)
	got := sustainability.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PenaltyAccount, got.PenaltyAccount)
	// this line is used by starport scaffolding # genesis/test/assert
}

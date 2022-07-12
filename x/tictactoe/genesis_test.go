package tictactoe_test

import (
	"testing"

	keepertest "github.com/alice/tictactoe/testutil/keeper"
	"github.com/alice/tictactoe/testutil/nullify"
	"github.com/alice/tictactoe/x/tictactoe"
	"github.com/alice/tictactoe/x/tictactoe/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TictactoeKeeper(t)
	tictactoe.InitGenesis(ctx, *k, genesisState)
	got := tictactoe.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

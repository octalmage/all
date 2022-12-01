package all_test

import (
	"testing"

	keepertest "github.com/octalmage/all/testutil/keeper"
	"github.com/octalmage/all/testutil/nullify"
	"github.com/octalmage/all/x/all"
	"github.com/octalmage/all/x/all/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AllKeeper(t)
	all.InitGenesis(ctx, *k, genesisState)
	got := all.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

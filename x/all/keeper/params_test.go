package keeper_test

import (
	"testing"

	testkeeper "github.com/octalmage/all/testutil/keeper"
	"github.com/octalmage/all/x/all/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AllKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

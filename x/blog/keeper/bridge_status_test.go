package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/blog/keeper"
	"planet/x/blog/types"
)

func createTestBridgeStatus(keeper *keeper.Keeper, ctx sdk.Context) types.BridgeStatus {
	item := types.BridgeStatus{}
	keeper.SetBridgeStatus(ctx, item)
	return item
}

func TestBridgeStatusGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	item := createTestBridgeStatus(keeper, ctx)
	rst, found := keeper.GetBridgeStatus(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestBridgeStatusRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	createTestBridgeStatus(keeper, ctx)
	keeper.RemoveBridgeStatus(ctx)
	_, found := keeper.GetBridgeStatus(ctx)
	require.False(t, found)
}

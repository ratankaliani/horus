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

func createNSentAction(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentAction {
	items := make([]types.SentAction, n)
	for i := range items {
		items[i].Id = keeper.AppendSentAction(ctx, items[i])
	}
	return items
}

func TestSentActionGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentAction(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentAction(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSentActionRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentAction(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentAction(ctx, item.Id)
		_, found := keeper.GetSentAction(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentActionGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentAction(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSentAction(ctx)),
	)
}

func TestSentActionCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentAction(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentActionCount(ctx))
}

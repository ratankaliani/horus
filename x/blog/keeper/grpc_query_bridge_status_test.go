package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/blog/types"
)

func TestBridgeStatusQuery(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestBridgeStatus(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBridgeStatusRequest
		response *types.QueryGetBridgeStatusResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBridgeStatusRequest{},
			response: &types.QueryGetBridgeStatusResponse{BridgeStatus: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BridgeStatus(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

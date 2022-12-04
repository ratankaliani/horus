package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"planet/testutil/network"
	"planet/testutil/nullify"
	"planet/x/blog/client/cli"
	"planet/x/blog/types"
)

func networkWithBridgeStatusObjects(t *testing.T) (*network.Network, types.BridgeStatus) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	bridgeStatus := &types.BridgeStatus{}
	nullify.Fill(&bridgeStatus)
	state.BridgeStatus = bridgeStatus
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.BridgeStatus
}

func TestShowBridgeStatus(t *testing.T) {
	net, obj := networkWithBridgeStatusObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.BridgeStatus
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowBridgeStatus(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetBridgeStatusResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.BridgeStatus)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.BridgeStatus),
				)
			}
		})
	}
}

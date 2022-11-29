package keeper

import (
	"context"

	"planet/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

func (k msgServer) SendIbcHorusAction(goCtx context.Context, msg *types.MsgSendIbcHorusAction) (*types.MsgSendIbcHorusActionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcHorusActionPacketData

	packet.Title = msg.Title
	packet.Action = msg.Action
	packet.Creator = msg.Creator

	// Transmit the packet
	err := k.TransmitIbcHorusActionPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcHorusActionResponse{}, nil
}

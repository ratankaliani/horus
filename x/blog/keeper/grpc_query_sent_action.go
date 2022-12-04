package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"planet/x/blog/types"
)

func (k Keeper) SentActionAll(c context.Context, req *types.QueryAllSentActionRequest) (*types.QueryAllSentActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentActions []types.SentAction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sentActionStore := prefix.NewStore(store, types.KeyPrefix(types.SentActionKey))

	pageRes, err := query.Paginate(sentActionStore, req.Pagination, func(key []byte, value []byte) error {
		var sentAction types.SentAction
		if err := k.cdc.Unmarshal(value, &sentAction); err != nil {
			return err
		}

		sentActions = append(sentActions, sentAction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentActionResponse{SentAction: sentActions, Pagination: pageRes}, nil
}

func (k Keeper) SentAction(c context.Context, req *types.QueryGetSentActionRequest) (*types.QueryGetSentActionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	sentAction, found := k.GetSentAction(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentActionResponse{SentAction: sentAction}, nil
}

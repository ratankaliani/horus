package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/types"
)

// SetBridgeStatus set bridgeStatus in the store
func (k Keeper) SetBridgeStatus(ctx sdk.Context, bridgeStatus types.BridgeStatus) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BridgeStatusKey))
	b := k.cdc.MustMarshal(&bridgeStatus)
	store.Set([]byte{0}, b)
}

// GetBridgeStatus returns bridgeStatus
func (k Keeper) GetBridgeStatus(ctx sdk.Context) (val types.BridgeStatus, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BridgeStatusKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBridgeStatus removes bridgeStatus from the store
func (k Keeper) RemoveBridgeStatus(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BridgeStatusKey))
	store.Delete([]byte{0})
}

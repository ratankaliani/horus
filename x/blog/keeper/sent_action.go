package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/types"
)

// GetSentActionCount get the total number of sentAction
func (k Keeper) GetSentActionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentActionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentActionCount set the total number of sentAction
func (k Keeper) SetSentActionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentActionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentAction appends a sentAction in the store with a new id and update the count
func (k Keeper) AppendSentAction(
	ctx sdk.Context,
	sentAction types.SentAction,
) uint64 {
	// Create the sentAction
	count := k.GetSentActionCount(ctx)

	// Set the ID of the appended value
	sentAction.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentActionKey))
	appendedValue := k.cdc.MustMarshal(&sentAction)
	store.Set(GetSentActionIDBytes(sentAction.Id), appendedValue)

	// Update sentAction count
	k.SetSentActionCount(ctx, count+1)

	return count
}

// SetSentAction set a specific sentAction in the store
func (k Keeper) SetSentAction(ctx sdk.Context, sentAction types.SentAction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentActionKey))
	b := k.cdc.MustMarshal(&sentAction)
	store.Set(GetSentActionIDBytes(sentAction.Id), b)
}

// GetSentAction returns a sentAction from its id
func (k Keeper) GetSentAction(ctx sdk.Context, id uint64) (val types.SentAction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentActionKey))
	b := store.Get(GetSentActionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentAction removes a sentAction from the store
func (k Keeper) RemoveSentAction(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentActionKey))
	store.Delete(GetSentActionIDBytes(id))
}

// GetAllSentAction returns all sentAction
func (k Keeper) GetAllSentAction(ctx sdk.Context) (list []types.SentAction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentActionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentAction
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentActionIDBytes returns the byte representation of the ID
func GetSentActionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentActionIDFromBytes returns ID in uint64 format from a byte array
func GetSentActionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

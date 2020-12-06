package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kkennis/nameservice/x/nameservice/types"
)

// Get creator of the item
func (k Keeper) GetNameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	name, err := k.GetName(ctx, key)
	if err != nil {
		return nil
	}
	return name.Owner
}


// Check if the key exists in the store
func (k Keeper) NameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.NamePrefix + key))
}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/sdk-tutorials/nameservice/nameservice/x/nameservice/types"
)

// CreateName creates a name. This function is included in starport type scaffolding.
// We won't use this function in our application, so it can be commented out.
// func (k Keeper) CreateName(ctx sdk.Context, name types.Name) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := []byte(types.NamePrefix + name.Value)
// 	value := k.cdc.MustMarshalBinaryLengthPrefixed(name)
// 	store.Set(key, value)
// }

// GetName returns the name information
func (k Keeper) GetName(ctx sdk.Context, key string) (types.Name, error) {
	store := ctx.KVStore(k.storeKey)
	var name types.Name
	byteKey := []byte(types.NamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &name)
	if err != nil {
		return name, err
	}
	return name, nil
}

// SetName sets a name. We modified this function to use the `name` value as the key instead of msg.ID
func (k Keeper) SetName(ctx sdk.Context, newName string, name types.Name) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(name)
	key := []byte(types.NamePrefix + newName)
	store.Set(key, bz)
}

// DeleteName deletes a name
func (k Keeper) DeleteName(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.NamePrefix + key))
}

//
// Functions used by querier
//

func listName(ctx sdk.Context, k Keeper) ([]byte, error) {
	var nameList []types.Name
	iterator := k.GetNamesIterator(ctx);
	for ; iterator.Valid(); iterator.Next() {
		var name types.Name
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &name)
		nameList = append(nameList, name)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, nameList)
	return res, nil
}

func getName(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	name, err := k.GetName(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, name)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Resolves a name, returns the value
func resolveName(ctx sdk.Context, path []string, keeper Keeper) ([]byte, error) {
	value := keeper.ResolveName(ctx, path[0])

	if value == "" {
		return []byte{}, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "could not resolve name")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: value})
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get owner of the item
func (k Keeper) GetOwner(ctx sdk.Context, key string) sdk.AccAddress {
	name, _ := k.GetName(ctx, key)
	return name.Owner
}

// Check if the key exists in the store
func (k Keeper) Exists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.NamePrefix + key))
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	nameObj, _ := k.GetName(ctx, name)
	return nameObj.Value
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetNameValue(ctx sdk.Context, name string, value string) {
	nameObj, _ := k.GetName(ctx, name)
	nameObj.Value = value
	k.SetName(ctx, name, nameObj)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	nameObj, _ := k.GetName(ctx, name)
	return !nameObj.Owner.Empty()
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	nameObj, _ := k.GetName(ctx, name)
	nameObj.Owner = owner
	k.SetName(ctx, name, nameObj)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	nameObj, _ := k.GetName(ctx, name)
	return nameObj.Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	nameObj, _ := k.GetName(ctx, name)
	nameObj.Price = price
	k.SetName(ctx, name, nameObj)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// Get an iterator over all names in which the keys are the names and the values are the name
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.NamePrefix))
}

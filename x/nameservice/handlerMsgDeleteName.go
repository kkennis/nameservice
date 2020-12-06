package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/kkennis/nameservice/x/nameservice/types"
	"github.com/kkennis/nameservice/x/nameservice/keeper"
)

// Handle a message to delete name
func handleMsgDeleteName(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteName) (*sdk.Result, error) {
	if !keeper.Exists(ctx, msg.Name) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Name)
	}
	if !msg.Owner.Equals(k.GetOwner(ctx, msg.Name)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteName(ctx, msg.Name)
	return &sdk.Result{}, nil
}

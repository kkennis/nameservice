package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/user/nameservice/x/nameservice/types"
	"github.com/user/nameservice/x/nameservice/keeper"
)

func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) (*sdk.Result, error) {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { // Checks if the msg sender is same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	keeper.SetName(ctx, msg.Name, msg.Value) // If so, set the name to the value specified in the msg.
	return &sdk.Result{}, nil // return
}
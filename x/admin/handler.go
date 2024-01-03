package admin

import (
	"fmt"

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	sdkerrors "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/errors"
	"github.com/CudoVentures/cudos-node/x/admin/keeper"
	"github.com/CudoVentures/cudos-node/x/admin/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgAdminSpendCommunityPool:
			res, err := msgServer.AdminSpendCommunityPool(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

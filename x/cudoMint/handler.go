package cudoMint

import (
	"fmt"

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	sdkerrors "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/errors"
	"github.com/CudoVentures/cudos-node/x/cudoMint/keeper"
	"github.com/CudoVentures/cudos-node/x/cudoMint/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

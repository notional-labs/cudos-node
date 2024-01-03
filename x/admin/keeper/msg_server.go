package keeper

import (
	"context"

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	sdkerrors "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/errors"
	"github.com/CudoVentures/cudos-node/x/admin/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) AdminSpendCommunityPool(goCtx context.Context, proposal *types.MsgAdminSpendCommunityPool) (*types.MsgAdminSpendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	initiatorAddr, err := sdk.AccAddressFromBech32(proposal.Initiator)
	if err != nil {
		return nil, err
	}

	balance := m.Keeper.bankKeeper.GetBalance(ctx, initiatorAddr, types.AdminDenom)

	if !balance.IsPositive() {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Insufficient permissions. Address '%s' has no %s tokens", initiatorAddr, types.AdminDenom)
	}

	to, err := sdk.AccAddressFromBech32(proposal.ToAddress)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.AdminDistributeFromFeePool(ctx, proposal.Coins, to)
	if err != nil {
		return nil, err
	}
	return &types.MsgAdminSpendResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

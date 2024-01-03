package keeper

import (
	// this line is used by starport scaffolding # 1

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	sdkerrors "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/errors"
	"github.com/CudoVentures/cudos-node/x/admin/types"
	"github.com/cosmos/cosmos-sdk/codec"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 2
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	"github.com/CudoVentures/cudos-node/x/admin/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type (
	Keeper struct {
		cdc                codec.Codec
		storeKey           sdk.StoreKey
		memKey             sdk.StoreKey
		distributionKeeper types.DistributionKeeper
		bankKeeper         types.BankKeeper
	}
)

func NewKeeper(cdc codec.Codec, storeKey, memKey sdk.StoreKey,
	dk types.DistributionKeeper, bk types.BankKeeper) *Keeper {
	return &Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		distributionKeeper: dk,
		bankKeeper:         bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AdminDistributeFromFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error {
	return k.distributionKeeper.DistributeFromFeePool(ctx, amount, receiveAddr)
}

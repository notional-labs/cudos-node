package module

import (
	"github.com/CudoVentures/cudos-node/x/group/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	if err := k.TallyProposalsAtVPEnd(ctx); err != nil {
		panic(err)
	}
	pruneProposals(ctx, k)
}

func pruneProposals(ctx sdk.Context, k keeper.Keeper) {
	err := k.PruneProposals(ctx)
	if err != nil {
		panic(err)
	}
}

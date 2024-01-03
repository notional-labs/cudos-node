package module

import (
	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	"github.com/CudoVentures/cudos-node/x/group/keeper"
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

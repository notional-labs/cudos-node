package types

// DONTCOVER

import (
	sdkerrors "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/errors"
)

// x/cudoMint module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors
)

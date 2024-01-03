package sample

import (
	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

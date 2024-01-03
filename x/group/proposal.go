package group

import (
	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	"github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

func (p *Proposal) GetMsgs() ([]sdk.Msg, error) {
	return tx.GetMsgs(p.Messages, "proposal")
}

func (p *Proposal) SetMsgs(msgs []sdk.Msg) error {
	anys, err := tx.SetMsgs(msgs)
	if err != nil {
		return err
	}
	p.Messages = anys
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	return tx.UnpackInterfaces(unpacker, p.Messages)
}

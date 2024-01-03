package types

import (
	"fmt"

	sdk "github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Params struct {
	BlocksPerDay sdk.Int `protobuf:"bytes,1,opt,name=blocks_per_day,json=blocksPerDay,proto3,customtype=github.com/CudoVentures/cudos-node/third_party/pkg/cudos-forked-sdk/types.Int" json:"blocks_per_day"`
}

// IncrementModifier Parameter store keys
var (
	BlocksPerDay = []byte("BlocksPerDay")
)

// ParamKeyTable ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	blocksPerDay sdk.Int,
) Params {
	return Params{
		BlocksPerDay: blocksPerDay,
	}
}

// DefaultParams default minting module parameters
func DefaultParams() Params {
	return Params{
		BlocksPerDay: sdk.NewInt(17280), // assuming 5 second block times
	}
}

// Validate validate params
func (p Params) Validate() error {
	if err := validateBlocksPerDay(p.BlocksPerDay); err != nil {
		return err
	}

	return nil

}

// ParamSetPairs Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(BlocksPerDay, &p.BlocksPerDay, validateBlocksPerDay),
	}
}

func validateBlocksPerDay(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.IsZero() {
		return fmt.Errorf("blocks per day must be positive: %s", v)
	}
	return nil
}

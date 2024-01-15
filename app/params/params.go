package params

import sdk "github.com/cosmos/cosmos-sdk/types"

// Simulation parameter constant
const (
	StakePerAccount           = "stake_per_account"
	InitiallyBondedValidators = "initially_bonded_validators"
	BondDenom                 = "acudos"
	AdminTokenDenom           = "cudosAdmin"
	MinSelfDelegation         = "2000000000000000000000000"
)

var DefaultPowerReduction = sdk.NewIntFromUint64(1000000000000000000)

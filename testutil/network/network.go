package network

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/network"
)

type (
	Network = network.Network
	Config  = network.Config
)

// New creates instance with fully configured cosmos network.
// Accepts optional config, that will be used in place of the DefaultConfig() if provided.
func New(t *testing.T, configs ...network.Config) *network.Network {
	if len(configs) > 1 {
		panic("at most one config should be provided")
	}
	var cfg network.Config
	if len(configs) == 0 {
		cfg = network.DefaultConfig(func() network.TestFixture {
			return network.TestFixture{}
		})
	} else {
		cfg = configs[0]
	}
	net, err := network.New(t, "", cfg)
	if err != nil {
		panic(err)
	}

	t.Cleanup(net.Cleanup)
	return net
}

// DefaultConfig will initialize config for the network with custom application,
// genesis and single validator. All other parameters are inherited from cosmos-sdk/testutil/network.DefaultConfig
// func DefaultConfig() network.Config {
// 	encoding := app.MakeEncodingConfig()
// 	return network.Config{
// 		Codec:             encoding.Codec,
// 		TxConfig:          encoding.TxConfig,
// 		LegacyAmino:       encoding.Amino,
// 		InterfaceRegistry: encoding.InterfaceRegistry,
// 		AccountRetriever:  authtypes.AccountRetriever{},
// 		AppConstructor: func(val network.Validator) servertypes.Application {
// 			return app.New(
// 				val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0,
// 				encoding,
// 				simapp.EmptyAppOptions{},
// 				baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
// 				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
// 			)
// 		},
// 		GenesisState:    app.ModuleBasics.DefaultGenesis(encoding.Codec),
// 		TimeoutCommit:   2 * time.Second,
// 		ChainID:         "chain-" + tmrand.NewRand().Str(6),
// 		NumValidators:   1,
// 		BondDenom:       sdk.DefaultBondDenom,
// 		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
// 		AccountTokens:   sdk.TokensFromConsensusPower(1000, sdk.DefaultPowerReduction),
// 		StakingTokens:   sdk.TokensFromConsensusPower(500, sdk.DefaultPowerReduction),
// 		BondedTokens:    sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction),
// 		PruningStrategy: storetypes.PruningOptionNothing,
// 		CleanupDir:      true,
// 		SigningAlgo:     string(hd.Secp256k1Type),
// 		KeyringOptions:  []keyring.Option{},
// 	}
// }

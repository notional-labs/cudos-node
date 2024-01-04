package decorators_test

import (
	"github.com/CudoVentures/cudos-node/app/decorators"
	cudoMinttypes "github.com/CudoVentures/cudos-node/x/cudoMint/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (suite *AnteTestSuite) TestStakingMin() {
	suite.SetupTest(true) // setup
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	adminAmount := sdk.NewCoins(sdk.NewCoin("cudosAdmin", sdk.NewInt(100000000000000)))

	priv1, _, addr1 := testdata.KeyTestPubAddr()
	priv2, _, addr2 := testdata.KeyTestPubAddr()
	suite.Require().NoError(suite.app.BankKeeper.MintCoins(suite.ctx, cudoMinttypes.ModuleName, adminAmount))

	suite.Require().NoError(
		suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, cudoMinttypes.ModuleName, addr1, adminAmount),
	)
	decorator := decorators.NewMinSelfDelegationDecorator()
	antehandler := sdk.ChainAnteDecorators(decorator)

	// Set IsCheckTx to true
	suite.ctx = suite.ctx.WithIsCheckTx(true)

	// normal so ok
	suite.ctx = suite.ctx.WithBlockHeight(100)
	minSelfDelegator, _ := sdk.NewIntFromString(decorators.MinSelfDelegation)
	suite.Require().NoError(suite.txBuilder.SetMsgs(
		&stakingtypes.MsgCreateValidator{
			Description:       stakingtypes.Description{},
			Commission:        stakingtypes.CommissionRates{},
			MinSelfDelegation: minSelfDelegator,
			DelegatorAddress:  string(addr1),
			ValidatorAddress:  string(addr2),
			Pubkey:            nil,
			Value:             sdk.Coin{},
		},
	))

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1, priv2}, []uint64{0, 1}, []uint64{0, 0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().NoError(err)

	// Not enough self delegation
	suite.ctx = suite.ctx.WithBlockHeight(100)
	suite.Require().NoError(suite.txBuilder.SetMsgs(
		&stakingtypes.MsgCreateValidator{
			Description:       stakingtypes.Description{},
			Commission:        stakingtypes.CommissionRates{},
			MinSelfDelegation: sdk.NewInt(1),
			DelegatorAddress:  string(addr1),
			ValidatorAddress:  string(addr2),
			Pubkey:            nil,
			Value:             sdk.Coin{},
		},
	))

	tx, err = suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	_, err = antehandler(suite.ctx, tx, false)
	suite.Require().Error(err)
}

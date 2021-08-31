package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/tendermint/budget/app"
	"github.com/tendermint/budget/x/budget/keeper"
	"github.com/tendermint/budget/x/budget/types"
)

//// createTestApp returns a budget app with custom BudgetKeeper.
//func createTestApp(isCheckTx bool) (*app.BudgetApp, sdk.Context) {
//	app := app.Setup(isCheckTx)
//	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
//	app.BudgetKeeper.SetParams(ctx, types.DefaultParams())
//
//	return app, ctx
//}

const (
	denom1 = "denom1"
	denom2 = "denom2"
	denom3 = "denom3"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000),
		sdk.NewInt64Coin(denom1, 1_000_000_000),
		sdk.NewInt64Coin(denom2, 1_000_000_000),
		sdk.NewInt64Coin(denom3, 1_000_000_000))
	smallBalances = mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake")
)

type KeeperTestSuite struct {
	suite.Suite

	app               *simapp.BudgetApp
	ctx               sdk.Context
	keeper            keeper.Keeper
	addrs             []sdk.AccAddress
	budgetSourceAddrs []sdk.AccAddress
	collectionAddrs   []sdk.AccAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.BudgetKeeper.SetParams(ctx, types.DefaultParams())

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.BudgetKeeper
	suite.addrs = simapp.AddTestAddrs(suite.app, suite.ctx, 10, sdk.ZeroInt())
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	cAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr3")))
	cAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr4")))
	cAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr5")))
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr2")))
	tAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr3")))
	tAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr4")))
	tAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr5")))
	suite.budgetSourceAddrs = []sdk.AccAddress{tAddr1, tAddr2, tAddr3, tAddr4, tAddr5}
	suite.collectionAddrs = []sdk.AccAddress{cAddr1, cAddr2, cAddr3, cAddr4, cAddr5}
	for _, addr := range append(suite.addrs, suite.budgetSourceAddrs[:3]...) {
		err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, suite.budgetSourceAddrs[3], smallBalances)
	suite.Require().NoError(err)
}

func coinsEq(exp, got sdk.Coins) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func mustParseRFC3339(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func mustParseCoinsNormalized(coinStr string) (coins sdk.Coins) {
	coins, err := sdk.ParseCoinsNormalized(coinStr)
	if err != nil {
		panic(err)
	}
	return coins
}

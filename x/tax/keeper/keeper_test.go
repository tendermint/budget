package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/tendermint/tax/app"
	"github.com/tendermint/tax/x/tax/keeper"
	"github.com/tendermint/tax/x/tax/types"
)

//// createTestApp returns a tax app with custom TaxKeeper.
//func createTestApp(isCheckTx bool) (*app.TaxApp, sdk.Context) {
//	app := app.Setup(isCheckTx)
//	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
//	app.TaxKeeper.SetParams(ctx, types.DefaultParams())
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

	app             *simapp.TaxApp
	ctx             sdk.Context
	keeper          keeper.Keeper
	addrs           []sdk.AccAddress
	taxSourceAddrs  []sdk.AccAddress
	collectionAddrs []sdk.AccAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.TaxKeeper.SetParams(ctx, types.DefaultParams())

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.TaxKeeper
	suite.addrs = simapp.AddTestAddrs(suite.app, suite.ctx, 10, sdk.ZeroInt())
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	cAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr3")))
	cAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr4")))
	cAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr5")))
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr2")))
	tAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr3")))
	tAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr4")))
	tAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr5")))
	suite.taxSourceAddrs = []sdk.AccAddress{tAddr1, tAddr2, tAddr3, tAddr4, tAddr5}
	suite.collectionAddrs = []sdk.AccAddress{cAddr1, cAddr2, cAddr3, cAddr4, cAddr5}
	for _, addr := range append(suite.addrs, suite.taxSourceAddrs[:3]...) {
		err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, suite.taxSourceAddrs[3], smallBalances)
	suite.Require().NoError(err)
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

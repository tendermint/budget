package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/tendermint/budget/app"
	"github.com/tendermint/budget/x/budget/keeper"
	"github.com/tendermint/budget/x/budget/types"
)

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
	querier           keeper.Querier
	govHandler        govtypes.Handler
	addrs             []sdk.AccAddress
	budgetSourceAddrs []sdk.AccAddress
	collectionAddrs   []sdk.AccAddress
	budgets           []types.Budget
}

func testProposal(changes ...proposal.ParamChange) *proposal.ParameterChangeProposal {
	return proposal.NewParameterChangeProposal("title", "description", changes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.govHandler = params.NewParamChangeProposalHandler(suite.app.ParamsKeeper)

	suite.keeper = suite.app.BudgetKeeper
	suite.querier = keeper.Querier{Keeper: suite.keeper}
	suite.addrs = simapp.AddTestAddrs(suite.app, suite.ctx, 10, sdk.ZeroInt())
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	cAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr3")))
	cAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr4")))
	cAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr5")))
	cAddr6 := suite.app.AccountKeeper.GetModuleAccount(suite.ctx, authtypes.FeeCollectorName).GetAddress()
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr2")))
	tAddr3 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr3")))
	tAddr4 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr4")))
	tAddr5 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr5")))
	tAddr6 := sdk.AccAddress(address.Module("farming", []byte("GravityDEXFarmingBudget")))
	suite.collectionAddrs = []sdk.AccAddress{cAddr1, cAddr2, cAddr3, cAddr4, cAddr5, cAddr6}
	suite.budgetSourceAddrs = []sdk.AccAddress{tAddr1, tAddr2, tAddr3, tAddr4, tAddr5, tAddr6}
	for _, addr := range append(suite.addrs, suite.budgetSourceAddrs[:3]...) {
		err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, suite.budgetSourceAddrs[3], smallBalances)
	suite.Require().NoError(err)

	suite.budgets = []types.Budget{
		{
			Name:                "budget1",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
			CollectionAddress:   suite.collectionAddrs[0].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget2",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
			CollectionAddress:   suite.collectionAddrs[1].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget3",
			Rate:                sdk.MustNewDecFromStr("1.0"),
			BudgetSourceAddress: suite.budgetSourceAddrs[1].String(),
			CollectionAddress:   suite.collectionAddrs[2].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget4",
			Rate:                sdk.MustNewDecFromStr("1"),
			BudgetSourceAddress: suite.budgetSourceAddrs[2].String(),
			CollectionAddress:   suite.collectionAddrs[3].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("0000-01-02T00:00:00Z"),
		},
		{
			Name:                "budget5",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: suite.budgetSourceAddrs[3].String(),
			CollectionAddress:   suite.collectionAddrs[0].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget6",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: suite.budgetSourceAddrs[3].String(),
			CollectionAddress:   suite.collectionAddrs[1].String(),
			StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "gravity-dex-farming-20213Q-20313Q",
			Rate:                sdk.MustNewDecFromStr("0.5"),
			BudgetSourceAddress: suite.budgetSourceAddrs[5].String(),
			CollectionAddress:   suite.collectionAddrs[5].String(),
			StartTime:           types.MustParseRFC3339("2021-09-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2031-09-30T00:00:00Z"),
		},
	}
}

func coinsEq(exp, got sdk.Coins) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func mustParseCoinsNormalized(coinStr string) (coins sdk.Coins) {
	coins, err := sdk.ParseCoinsNormalized(coinStr)
	if err != nil {
		panic(err)
	}
	return coins
}

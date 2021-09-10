package keeper_test

import (
	_ "github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/types"
)

func (suite *KeeperTestSuite) TestGRPCParams() {
	resp, err := suite.querier.Params(sdk.WrapSDKContext(suite.ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(suite.keeper.GetParams(suite.ctx), resp.Params)
}

func (suite *KeeperTestSuite) TestGRPCBudgets() {
	budgets := []types.Budget{
		{
			Name:                "budget1",
			Rate:                sdk.NewDecWithPrec(5, 2),
			BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
			CollectionAddress:   suite.collectionAddrs[0].String(),
			StartTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget2",
			Rate:                sdk.NewDecWithPrec(5, 2),
			BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
			CollectionAddress:   suite.collectionAddrs[1].String(),
			StartTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget3",
			Rate:                sdk.NewDecWithPrec(5, 2),
			BudgetSourceAddress: suite.budgetSourceAddrs[1].String(),
			CollectionAddress:   suite.collectionAddrs[0].String(),
			StartTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:                "budget4",
			Rate:                sdk.NewDecWithPrec(5, 2),
			BudgetSourceAddress: suite.budgetSourceAddrs[1].String(),
			CollectionAddress:   suite.collectionAddrs[1].String(),
			StartTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:             mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
	}

	params := suite.keeper.GetParams(suite.ctx)
	params.Budgets = budgets
	suite.keeper.SetParams(suite.ctx, params)

	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.budgetSourceAddrs[0])
	expectedCoins, _ := sdk.NewDecCoinsFromCoins(balance...).MulDec(sdk.NewDecWithPrec(5, 2)).TruncateDecimal()

	suite.ctx = suite.ctx.WithBlockTime(mustParseRFC3339("2021-08-31T00:00:00Z"))
	err := suite.keeper.BudgetCollection(suite.ctx)
	suite.Require().NoError(err)

	for _, tc := range []struct {
		name      string
		req       *types.QueryBudgetsRequest
		expectErr bool
		postRun   func(response *types.QueryBudgetsResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query all",
			&types.QueryBudgetsRequest{},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 4)
			},
		},
		{
			"query by not existing name",
			&types.QueryBudgetsRequest{Name: "notfound"},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 0)
			},
		},
		{
			"query by name",
			&types.QueryBudgetsRequest{Name: "budget1"},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 1)
				suite.Require().Equal("budget1", resp.Budgets[0].Budget.Name)
			},
		},
		{
			"invalid budget source addr",
			&types.QueryBudgetsRequest{BudgetSourceAddress: "invalid"},
			true,
			nil,
		},
		{
			"query by budget source addr",
			&types.QueryBudgetsRequest{BudgetSourceAddress: suite.budgetSourceAddrs[0].String()},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 2)
				for _, b := range resp.Budgets {
					suite.Require().Equal(suite.budgetSourceAddrs[0].String(), b.Budget.BudgetSourceAddress)
				}
			},
		},
		{
			"invalid collection addr",
			&types.QueryBudgetsRequest{CollectionAddress: "invalid"},
			true,
			nil,
		},
		{
			"query by collection addr",
			&types.QueryBudgetsRequest{CollectionAddress: suite.collectionAddrs[0].String()},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 2)
				for _, b := range resp.Budgets {
					suite.Require().Equal(suite.collectionAddrs[0].String(), b.Budget.CollectionAddress)
				}
			},
		},
		{
			"query with multiple filters",
			&types.QueryBudgetsRequest{
				BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
				CollectionAddress:   suite.collectionAddrs[1].String(),
			},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 1)
				suite.Require().Equal(suite.budgetSourceAddrs[0].String(), resp.Budgets[0].Budget.BudgetSourceAddress)
				suite.Require().Equal(suite.collectionAddrs[1].String(), resp.Budgets[0].Budget.CollectionAddress)
			},
		},
		{
			"correct total collected coins",
			&types.QueryBudgetsRequest{Name: "budget1"},
			false,
			func(resp *types.QueryBudgetsResponse) {
				suite.Require().Len(resp.Budgets, 1)
				suite.Require().True(coinsEq(expectedCoins, resp.Budgets[0].TotalCollectedCoins))
			},
		},
	} {
		suite.Run(tc.name, func() {
			resp, err := suite.querier.Budgets(sdk.WrapSDKContext(suite.ctx), tc.req)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

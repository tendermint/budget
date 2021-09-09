package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/types"
)

func (suite *KeeperTestSuite) TestBudgetCollection() {
	for _, tc := range []struct {
		name           string
		budgets        []types.Budget
		epochBlocks    uint32
		accAsserts     []sdk.AccAddress
		balanceAsserts []sdk.Coins
		expectErr      bool
	}{
		{
			"basic budgets case",
			suite.budgets[:4],
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.budgetSourceAddrs[0],
				suite.budgetSourceAddrs[1],
				suite.budgetSourceAddrs[2],
			},
			[]sdk.Coins{
				mustParseCoinsNormalized("500000000denom1,500000000denom2,500000000denom3,500000000stake"),
				mustParseCoinsNormalized("500000000denom1,500000000denom2,500000000denom3,500000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				{},
				{},
				{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
			},
			false,
		},
		{
			"only expired budget case",
			[]types.Budget{suite.budgets[3]},
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[3],
				suite.budgetSourceAddrs[2],
			},
			[]sdk.Coins{
				{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
			},
			false,
		},
		{
			"budget source has small balances case",
			suite.budgets[4:6],
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.budgetSourceAddrs[3],
			},
			[]sdk.Coins{
				mustParseCoinsNormalized("1denom2,1denom3,500000000stake"),
				mustParseCoinsNormalized("1denom2,1denom3,500000000stake"),
				mustParseCoinsNormalized("1denom1,1denom3"),
			},
			false,
		},
		{
			"none budgets case",
			nil,
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.budgetSourceAddrs[0],
				suite.budgetSourceAddrs[1],
				suite.budgetSourceAddrs[2],
				suite.budgetSourceAddrs[3],
			},
			[]sdk.Coins{
				{},
				{},
				{},
				{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
		{
			"disabled budget epoch",
			nil,
			0,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.budgetSourceAddrs[0],
				suite.budgetSourceAddrs[1],
				suite.budgetSourceAddrs[2],
				suite.budgetSourceAddrs[3],
			},
			[]sdk.Coins{
				{},
				{},
				{},
				{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
		{
			"disabled budget epoch with budgets",
			suite.budgets[:4],
			0,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.budgetSourceAddrs[0],
				suite.budgetSourceAddrs[1],
				suite.budgetSourceAddrs[2],
				suite.budgetSourceAddrs[3],
			},
			[]sdk.Coins{
				{},
				{},
				{},
				{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
	} {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			params := suite.keeper.GetParams(suite.ctx)
			params.Budgets = tc.budgets
			params.EpochBlocks = tc.epochBlocks
			suite.keeper.SetParams(suite.ctx, params)

			err := suite.keeper.BudgetCollection(suite.ctx)
			if tc.expectErr {
				suite.Error(err)
			} else {
				suite.NoError(err)

				for i, acc := range tc.accAsserts {
					suite.True(suite.app.BankKeeper.GetAllBalances(suite.ctx, acc).IsEqual(tc.balanceAsserts[i]))
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestBudgetExpiration() {
	// TODO: not implemented
}

func (suite *KeeperTestSuite) TestGetSetTotalCollectedCoins() {
	collectedCoins := suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().Nil(collectedCoins)

	suite.keeper.SetTotalCollectedCoins(suite.ctx, "budget1", sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	collectedCoins = suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)), collectedCoins))

	suite.keeper.AddTotalCollectedCoins(suite.ctx, "budget1", sdk.NewCoins(sdk.NewInt64Coin(denom2, 1000000)))
	collectedCoins = suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000), sdk.NewInt64Coin(denom2, 1000000)), collectedCoins))

	suite.keeper.AddTotalCollectedCoins(suite.ctx, "budget2", sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))
	collectedCoins = suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget2")
	suite.Require().True(coinsEq(sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)), collectedCoins))
}

func (suite *KeeperTestSuite) TestTotalCollectedCoins() {
	budget := types.Budget{
		Name:                "budget1",
		Rate:                sdk.NewDecWithPrec(5, 2), // 5%
		BudgetSourceAddress: suite.budgetSourceAddrs[0].String(),
		CollectionAddress:   suite.collectionAddrs[0].String(),
		StartTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
		EndTime:             mustParseRFC3339("9999-12-31T00:00:00Z"),
	}

	params := suite.keeper.GetParams(suite.ctx)
	params.Budgets = []types.Budget{budget}
	suite.keeper.SetParams(suite.ctx, params)

	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.budgetSourceAddrs[0])
	expectedCoins, _ := sdk.NewDecCoinsFromCoins(balance...).MulDec(sdk.NewDecWithPrec(5, 2)).TruncateDecimal()

	collectedCoins := suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().Equal(sdk.Coins(nil), collectedCoins)

	suite.ctx = suite.ctx.WithBlockTime(mustParseRFC3339("2021-08-31T00:00:00Z"))
	err := suite.keeper.BudgetCollection(suite.ctx)
	suite.Require().NoError(err)

	collectedCoins = suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().True(coinsEq(expectedCoins, collectedCoins))
}

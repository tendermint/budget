package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramscutils "github.com/cosmos/cosmos-sdk/x/params/client/utils"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"

	"github.com/tendermint/budget/app"
	"github.com/tendermint/budget/x/budget/types"
)

func (suite *KeeperTestSuite) TestCollectBudgets() {
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

			err := suite.keeper.CollectBudgets(suite.ctx)
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

func (suite *KeeperTestSuite) TestBudgetChangeSituation() {
	encCfg := app.MakeTestEncodingConfig()
	params := suite.keeper.GetParams(suite.ctx)
	suite.keeper.SetParams(suite.ctx, params)
	height := 1
	suite.ctx = suite.ctx.WithBlockTime(types.MustParseRFC3339("2021-08-01T00:00:00Z"))
	suite.ctx = suite.ctx.WithBlockHeight(int64(height))

	for _, tc := range []struct {
		name                   string
		proposal               *proposal.ParameterChangeProposal
		budgetCount            int
		collectibleBudgetCount int
		govTime                time.Time
		nextBlockTime          time.Time
		expErr                 error
	}{
		{
			"add budget 1",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					}
				]`,
			}),
			1,
			0,
			types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			nil,
		},
		{
			"add budget 2",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-2",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1m63436cdxnu9ymyj02e7k3xljkn8klyf5ahqa75degq748xxkmksvtlp8n",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2021-09-30T00:00:00Z"
					}
				]`,
			}),
			2,
			2,
			types.MustParseRFC3339("2021-09-03T00:00:00Z"),
			types.MustParseRFC3339("2021-09-03T00:00:00Z"),
			nil,
		},
		{
			"add budget 3 with invalid total rate case 1",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-2",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1m63436cdxnu9ymyj02e7k3xljkn8klyf5ahqa75degq748xxkmksvtlp8n",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2021-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-3",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos17avp6xs5c8ycqzy20yv99ccxwunu32e507kpm8ql5nfg47pzj9qqxhujxr",
					"start_time": "2021-09-30T00:00:00Z",
					"end_time": "2021-10-10T00:00:00Z"
					}
				]`,
			}),
			0,
			0,
			types.MustParseRFC3339("2021-09-29T00:00:00Z"),
			types.MustParseRFC3339("2021-09-30T00:00:00Z"),
			types.ErrInvalidTotalBudgetRate,
		},
		{
			"add budget 3 with invalid total rate case 2",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-2",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1m63436cdxnu9ymyj02e7k3xljkn8klyf5ahqa75degq748xxkmksvtlp8n",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2021-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-3",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos17avp6xs5c8ycqzy20yv99ccxwunu32e507kpm8ql5nfg47pzj9qqxhujxr",
					"start_time": "2021-09-30T00:00:00Z",
					"end_time": "2021-10-10T00:00:00Z"
					}
				]`,
			}),
			0,
			0,
			types.MustParseRFC3339("2021-10-01T00:00:00Z"),
			types.MustParseRFC3339("2021-10-01T00:00:00Z"),
			types.ErrInvalidTotalBudgetRate,
		},
		{
			"add budget 3",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-3",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos17avp6xs5c8ycqzy20yv99ccxwunu32e507kpm8ql5nfg47pzj9qqxhujxr",
					"start_time": "2021-09-30T00:00:00Z",
					"end_time": "2021-10-10T00:00:00Z"
					}
				]`,
			}),
			2,
			2,
			types.MustParseRFC3339("2021-10-01T00:00:00Z"),
			types.MustParseRFC3339("2021-10-01T00:00:00Z"),
			nil,
		},
		{
			"add budget 4 without date range overlap",
			testProposal(proposal.ParamChange{
				Subspace: types.ModuleName,
				Key:      string(types.KeyBudgets),
				Value: `[
					{
					"name": "gravity-dex-farming-20213Q-20313Q",
					"rate": "0.500000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
					"start_time": "2021-09-01T00:00:00Z",
					"end_time": "2031-09-30T00:00:00Z"
					},
					{
					"name": "gravity-dex-farming-4",
					"rate": "1.000000000000000000",
					"budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
					"collection_address": "cosmos17avp6xs5c8ycqzy20yv99ccxwunu32e507kpm8ql5nfg47pzj9qqxhujxr",
					"start_time": "2031-09-30T00:00:01Z",
					"end_time": "2031-12-10T00:00:00Z"
					}
				]`,
			}),
			2,
			1,
			types.MustParseRFC3339("2021-09-29T00:00:00Z"),
			types.MustParseRFC3339("2021-09-30T00:00:00Z"),
			nil,
		},
	} {
		suite.Run(tc.name, func() {
			proposalJson := paramscutils.ParamChangeProposalJSON{}
			bz, err := tc.proposal.Marshal()
			suite.Require().NoError(err)
			err = encCfg.Amino.Unmarshal(bz, &proposalJson)
			suite.Require().NoError(err)
			proposal := paramproposal.NewParameterChangeProposal(
				proposalJson.Title, proposalJson.Description, proposalJson.Changes.ToParamChanges(),
			)
			suite.Require().NoError(err)

			// endblock gov paramchange ->(new block)-> beginblock budget -> mempool -> endblock gov paramchange ->(new block)-> ...
			suite.ctx = suite.ctx.WithBlockTime(tc.govTime)
			err = suite.govHandler(suite.ctx, proposal)
			if tc.expErr != nil {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				params := suite.keeper.GetParams(suite.ctx)
				suite.Require().Len(params.Budgets, tc.budgetCount)
				for _, budget := range params.Budgets {
					err := budget.Validate()
					suite.Require().NoError(err)
				}
				// (new block)
				height += 1
				suite.ctx = suite.ctx.WithBlockHeight(int64(height))
				suite.ctx = suite.ctx.WithBlockTime(tc.nextBlockTime)
				budgets := suite.keeper.CollectibleBudgets(params.Budgets, suite.ctx.BlockTime())
				suite.Require().Len(budgets, tc.collectibleBudgetCount)

				// BeginBlocker
				err := suite.keeper.CollectBudgets(suite.ctx)
				suite.Require().NoError(err)
			}
		})
	}
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
		StartTime:           types.MustParseRFC3339("0000-01-01T00:00:00Z"),
		EndTime:             types.MustParseRFC3339("9999-12-31T00:00:00Z"),
	}

	params := suite.keeper.GetParams(suite.ctx)
	params.Budgets = []types.Budget{budget}
	suite.keeper.SetParams(suite.ctx, params)

	balance := suite.app.BankKeeper.GetAllBalances(suite.ctx, suite.budgetSourceAddrs[0])
	expectedCoins, _ := sdk.NewDecCoinsFromCoins(balance...).MulDec(sdk.NewDecWithPrec(5, 2)).TruncateDecimal()

	collectedCoins := suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().Equal(sdk.Coins(nil), collectedCoins)

	suite.ctx = suite.ctx.WithBlockTime(types.MustParseRFC3339("2021-08-31T00:00:00Z"))
	err := suite.keeper.CollectBudgets(suite.ctx)
	suite.Require().NoError(err)

	collectedCoins = suite.keeper.GetTotalCollectedCoins(suite.ctx, "budget1")
	suite.Require().True(coinsEq(expectedCoins, collectedCoins))
}

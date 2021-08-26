package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/budget/x/budget/types"
)

func (k Keeper) BudgetCollection(ctx sdk.Context) error {
	// Get all the Budgets registered in params.Budgets and select only the valid budgets. If there is no valid budget, exit.
	budgets := k.CollectibleBudgets(ctx)
	if len(budgets) == 0 {
		return nil
	}

	var inputs []banktypes.Input
	var outputs []banktypes.Output
	sendCoins := func(from, to sdk.AccAddress, coins sdk.Coins) {
		if !coins.Empty() && coins.IsValid() {
			inputs = append(inputs, banktypes.NewInput(from, coins))
			outputs = append(outputs, banktypes.NewOutput(to, coins))
		}
	}
	// Create a map by BudgetSourceAddress to handle the budgets for the same BudgetSourceAddress together based on the
	// same balance when calculating rates for the same BudgetSourceAddress.
	budgetsBySourceMap := types.GetBudgetsBySourceMap(budgets)
	for budgetSource, budgetsBySource := range budgetsBySourceMap {
		budgetSourceAcc, err := sdk.AccAddressFromBech32(budgetSource)
		if err != nil {
			continue
		}
		budgetSourceBalances := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, budgetSourceAcc)...)
		if budgetSourceBalances.IsZero() {
			continue
		}
		expectedCollectionCoins, _ := budgetSourceBalances.MulDecTruncate(budgetsBySource.TotalRate).TruncateDecimal()
		var validatedExpectedCollectionCoins sdk.Coins
		for _, coin := range expectedCollectionCoins {
			if coin.IsValid() && coin.IsZero() {
				validatedExpectedCollectionCoins = append(validatedExpectedCollectionCoins, coin)
			}
		}
		expectedDiffCoins := expectedCollectionCoins.Sub(validatedExpectedCollectionCoins)

		var totalCollectionCoins sdk.Coins
		var totalChangeCoins sdk.DecCoins
		for _, budget := range budgetsBySource.Budgets {
			collectionAcc, err := sdk.AccAddressFromBech32(budget.CollectionAddress)
			if err != nil {
				continue
			}
			collectionCoins, changeCoins := budgetSourceBalances.MulDecTruncate(budget.Rate).TruncateDecimal()
			totalCollectionCoins = totalCollectionCoins.Add(collectionCoins...)
			totalChangeCoins = totalChangeCoins.Add(changeCoins...)
			// TODO: sendcoins after validation
			sendCoins(budgetSourceAcc, collectionAcc, collectionCoins)
		}
		// temporary validation logic
		if totalCollectionCoins.IsAnyGT(validatedExpectedCollectionCoins) {
			panic("totalCollectionCoins.IsAnyGT(expectedCollectionCoins)")
		}
		if _, neg := sdk.NewDecCoinsFromCoins(expectedDiffCoins...).SafeSub(totalChangeCoins); neg {
			panic("expectedChangeCoins.Sub(totalChangeCoins).IsAnyNegative()")
		}
	}
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}
	// TODO: add metric or record total collection coins each budget
	return nil
}

// Get all the Budgets registered in params.Budgets and return only the valid and not expired budgets
func (k Keeper) CollectibleBudgets(ctx sdk.Context) (budgets []types.Budget) {
	params := k.GetParams(ctx)
	if params.EpochBlocks > 0 && ctx.BlockHeight()%int64(params.EpochBlocks) == 0 {
		for _, budget := range params.Budgets {
			err := budget.Validate()
			if err == nil && !budget.Expired(ctx.BlockTime()) {
				budgets = append(budgets, budget)
			}
		}
	}
	return
}

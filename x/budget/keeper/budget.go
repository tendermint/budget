package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/budget/x/budget/types"
)

// CollectBudgets collects all the valid budgets registered in params.Budgets and
// distributes the total collected coins to collection address.
func (k Keeper) CollectBudgets(ctx sdk.Context) error {
	budgets := k.CollectibleBudgets(ctx)
	if len(budgets) == 0 {
		return nil
	}

	var inputs []banktypes.Input
	var outputs []banktypes.Output

	// Get a map GetBudgetsBySourceMap that has a list of budgets and their total rate, which
	// contain the same BudgetSourceAddress
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

			if !collectionCoins.Empty() && collectionCoins.IsValid() {
				inputs = append(inputs, banktypes.NewInput(budgetSourceAcc, collectionCoins))
				outputs = append(outputs, banktypes.NewOutput(collectionAcc, collectionCoins))
			}

			k.AddTotalCollectedCoins(ctx, budget.Name, collectionCoins)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeBudgetCollected,
					sdk.NewAttribute(types.AttributeValueName, budget.Name),
					sdk.NewAttribute(types.AttributeValueCollectionAddress, budget.CollectionAddress),
					sdk.NewAttribute(types.AttributeValueBudgetSourceAddress, budget.BudgetSourceAddress),
					sdk.NewAttribute(types.AttributeValueRate, budget.Rate.String()),
					sdk.NewAttribute(types.AttributeValueAmount, collectionCoins.String()),
				),
			})
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
	// TODO: add metric
	return nil
}

// CollectibleBudgets returns scan through the budgets registered in params.Budgets
// and returns only the valid and not expired budgets.
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

// GetTotalCollectedCoins returns total collected coins for a budget.
func (k Keeper) GetTotalCollectedCoins(ctx sdk.Context, budgetName string) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTotalCollectedCoinsKey(budgetName))
	if bz == nil {
		return nil
	}
	var collectedCoins types.TotalCollectedCoins
	k.cdc.MustUnmarshal(bz, &collectedCoins)
	return collectedCoins.TotalCollectedCoins
}

// IterateAllTotalCollectedCoins iterates over all the stored TotalCollectedCoins and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllTotalCollectedCoins(ctx sdk.Context, cb func(record types.BudgetRecord) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TotalCollectedCoinsKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var record types.BudgetRecord
		var collectedCoins types.TotalCollectedCoins
		k.cdc.MustUnmarshal(iterator.Value(), &collectedCoins)
		record.Name = types.ParseTotalCollectedCoinsKey(iterator.Key())
		record.TotalCollectedCoins = collectedCoins.TotalCollectedCoins
		if cb(record) {
			break
		}
	}
}

// SetTotalCollectedCoins sets total collected coins for a budget.
func (k Keeper) SetTotalCollectedCoins(ctx sdk.Context, budgetName string, amount sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	collectedCoins := types.TotalCollectedCoins{TotalCollectedCoins: amount}
	bz := k.cdc.MustMarshal(&collectedCoins)
	store.Set(types.GetTotalCollectedCoinsKey(budgetName), bz)
}

// AddTotalCollectedCoins increases total collected coins for a budget.
func (k Keeper) AddTotalCollectedCoins(ctx sdk.Context, budgetName string, amount sdk.Coins) {
	collectedCoins := k.GetTotalCollectedCoins(ctx, budgetName)
	collectedCoins = collectedCoins.Add(amount...)
	k.SetTotalCollectedCoins(ctx, budgetName, collectedCoins)
}

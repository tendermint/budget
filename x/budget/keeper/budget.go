package keeper

import (
	"github.com/armon/go-metrics"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/budget/x/budget/types"
)

// CollectBudgets collects all the valid budgets registered in params.Budgets and
// distributes the total collected coins to collection address.
func (k Keeper) CollectBudgets(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	var budgets []types.Budget
	if params.EpochBlocks > 0 && ctx.BlockHeight()%int64(params.EpochBlocks) == 0 {
		budgets = types.CollectibleBudgets(params.Budgets, ctx.BlockTime())
	}
	if len(budgets) == 0 {
		return nil
	}

	// Get a map GetBudgetsBySourceMap that has a list of budgets and their total rate, which
	// contain the same BudgetSourceAddress
	budgetsBySourceMap := types.GetBudgetsBySourceMap(budgets)
	for budgetSource, budgetsBySource := range budgetsBySourceMap {
		budgetSourceAcc, err := sdk.AccAddressFromBech32(budgetSource)
		if err != nil {
			return err
		}
		budgetSourceBalances := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, budgetSourceAcc)...)
		if budgetSourceBalances.IsZero() {
			continue
		}

		var inputs []banktypes.Input
		var outputs []banktypes.Output
		budgetsBySource.CollectionCoins = make([]sdk.Coins, len(budgetsBySource.Budgets))
		for i, budget := range budgetsBySource.Budgets {
			collectionAcc, err := sdk.AccAddressFromBech32(budget.CollectionAddress)
			if err != nil {
				return err
			}

			collectionCoins, _ := budgetSourceBalances.MulDecTruncate(budget.Rate).TruncateDecimal()
			if collectionCoins.Empty() || !collectionCoins.IsValid() {
				continue
			}

			inputs = append(inputs, banktypes.NewInput(budgetSourceAcc, collectionCoins))
			outputs = append(outputs, banktypes.NewOutput(collectionAcc, collectionCoins))
			budgetsBySource.CollectionCoins[i] = collectionCoins
		}

		if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
			return err
		}

		for i, budget := range budgetsBySource.Budgets {
			// Capture the variables in a loop for the deferred func
			i := i
			collectionCoins := budgetsBySource.CollectionCoins
			defer func() {
				for _, coin := range collectionCoins[i] {
					if coin.Amount.IsInt64() {
						telemetry.SetGaugeWithLabels(
							[]string{types.ModuleName},
							float32(coin.Amount.Int64()),
							[]metrics.Label{
								telemetry.NewLabel("collection_address", budget.CollectionAddress),
								telemetry.NewLabel("denom", coin.Denom),
							},
						)
					}
				}
			}()

			k.AddTotalCollectedCoins(ctx, budget.Name, collectionCoins[i])

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeBudgetCollected,
					sdk.NewAttribute(types.AttributeValueName, budget.Name),
					sdk.NewAttribute(types.AttributeValueCollectionAddress, budget.CollectionAddress),
					sdk.NewAttribute(types.AttributeValueBudgetSourceAddress, budget.BudgetSourceAddress),
					sdk.NewAttribute(types.AttributeValueRate, budget.Rate.String()),
					sdk.NewAttribute(types.AttributeValueAmount, collectionCoins[i].String()),
				),
			})
		}
	}
	return nil
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

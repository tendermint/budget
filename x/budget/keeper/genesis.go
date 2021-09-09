package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/types"
)

// InitGenesis initializes the budget module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err)
	}

	k.SetParams(ctx, genState.Params)
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	k.accountKeeper.SetModuleAccount(ctx, moduleAcc)

	for _, record := range genState.BudgetRecords {
		k.SetTotalCollectedCoins(ctx, record.Name, record.TotalCollectedCoins)
	}
}

// ExportGenesis returns the budget module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	var budgetRecords []types.BudgetRecord
	for _, budget := range params.Budgets {
		budgetRecords = append(budgetRecords, types.BudgetRecord{
			Name:                budget.Name,
			TotalCollectedCoins: k.GetTotalCollectedCoins(ctx, budget.Name),
		})
	}
	return types.NewGenesisState(params, budgetRecords)
}

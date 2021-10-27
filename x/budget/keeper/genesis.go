package keeper

import (
	"fmt"

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
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	for _, record := range genState.BudgetRecords {
		k.SetTotalCollectedCoins(ctx, record.Name, record.TotalCollectedCoins)
	}
}

// ExportGenesis returns the budget module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	var budgetRecords []types.BudgetRecord

	k.IterateAllTotalCollectedCoins(ctx, func(record types.BudgetRecord) (stop bool) {
		budgetRecords = append(budgetRecords, record)
		return false
	})

	return types.NewGenesisState(params, budgetRecords)
}

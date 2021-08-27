package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers all budget invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	//ir.RegisterRoute(types.ModuleName, "budget_pool_accounts_invariant",
	//	BudgetPoolAccountsInvariant(k))
}

// AllInvariants runs all invariants of the budget module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := BudgetPoolAccountsInvariant(k)(ctx)
		return res, stop
	}
}

// BudgetPoolAccountsInvariant checks that invariants of budget source, collecting account with each module account name
func BudgetPoolAccountsInvariant(k Keeper) sdk.Invariant {
	// TODO: not implemented
	return nil
	//return func(ctx sdk.Context) (string, bool) {
	//
	//	//return sdk.FormatInvariant(types.ModuleName, ""), true
	//}
}

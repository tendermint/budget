package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tax/x/tax/types"
)

// RegisterInvariants registers all tax invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "tax_pool_accounts_invariant",
		TaxPoolAccountsInvariant(k))
}

// AllInvariants runs all invariants of the tax module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := TaxPoolAccountsInvariant(k)(ctx)
		return res, stop
	}
}

// TaxPoolAccountsInvariant checks that invariants of tax source, collecting account with each module account name
func TaxPoolAccountsInvariant(k Keeper) sdk.Invariant {
	// TODO: not implemented
	return nil
	//return func(ctx sdk.Context) (string, bool) {
	//
	//	//return sdk.FormatInvariant(types.ModuleName, ""), true
	//}
}

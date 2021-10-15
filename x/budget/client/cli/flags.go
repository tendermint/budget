package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagName                = "name"
	FlagBudgetSourceAddress = "budget-source-address"
	FlagCollectionAddress   = "collection-address"
)

// flagSetBudgets returns the FlagSet used for budgets.
func flagSetBudgets() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagName, "", "The budget name")
	fs.String(FlagBudgetSourceAddress, "", "The bech32 address of the budget source account")
	fs.String(FlagCollectionAddress, "", "The bech32 address of the collection account")

	return fs
}

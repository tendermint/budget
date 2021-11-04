package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagName               = "name"
	FlagSourceAddress      = "source-address"
	FlagDestinationAddress = "destination-address"
)

// flagSetBudgets returns the FlagSet used for budgets.
func flagSetBudgets() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagName, "", "The budget name")
	fs.String(FlagSourceAddress, "", "The bech32 address of the source account")
	fs.String(FlagDestinationAddress, "", "The bech32 address of the destination account")

	return fs
}

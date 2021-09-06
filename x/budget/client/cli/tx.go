package cli

// Modification of Budgets of Params proceeds to governance proposition, not to Tx.
//import (
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/spf13/cobra"
//
//	"github.com/tendermint/budget/x/budget/types"
//)
//
//// GetTxCmd returns a root CLI command handler for all x/budget transaction commands.
//func GetTxCmd() *cobra.Command {
//	budgetTxCmd := &cobra.Command{
//		Use:                        types.ModuleName,
//		Short:                      "Budget transaction subcommands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	budgetTxCmd.AddCommand(
//	)
//
//	return budgetTxCmd
//}

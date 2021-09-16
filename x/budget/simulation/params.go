package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/budget/x/budget/types"
)

// DONTCOVER

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEpochBlocks),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%d", GenEpochBlocks(r))
			},
		),
		// TODO: Randomize structured ParamChange
		// ref. https://github.com/cosmos/cosmos-sdk/blob/49102b1d988f542c4293af7a85e403d858f348a8/x/gov/simulation/params.go#L39
		//simulation.NewSimParamChange(types.ModuleName, string(types.KeyBudgets),
		//	func(r *rand.Rand) string {
		//		return fmt.Sprintf("%s", GenBudgets(r))
		//	},
		//),
	}
}

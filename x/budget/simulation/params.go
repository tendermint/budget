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
		// TODO: Randomize
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEpochBlocks),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenEpochBlocks(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyBudgets),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenBudgets(r))
			},
		),
	}
}

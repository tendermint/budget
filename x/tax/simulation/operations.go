package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/tax/x/tax/keeper"
	"github.com/tendermint/tax/x/tax/types"
)

// Simulation operation weights constants
const (
	OpWeightUpdateTaxProposal = "op_weight_update_tax_proposal"
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec, ak types.AccountKeeper,
	bk types.BankKeeper, k keeper.Keeper,
) simulation.WeightedOperations {

	// TODO: Not implemented, it should be weight of UpdateTaxProposal or ParamChanges, not msg
	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			100,
			SimulateUpdateTaxProposal(ak, bk, k),
		),
	}
}

// SimulateMsgCreateFixedAmountPlan generates a MsgCreateFixedAmountPlan with random values
// nolint: interfacer
func SimulateUpdateTaxProposal(ak types.AccountKeeper, bk types.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// TODO: not implemented yet
		return simtypes.OperationMsg{}, nil, nil
	}
}

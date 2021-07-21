package simulation_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"

	"github.com/tendermint/tax/x/tax/simulation"
)

func TestDecodeTaxStore(t *testing.T) {

	cdc := simapp.MakeTestEncodingConfig()
	_ = simulation.NewDecodeStore(cdc.Marshaler)

	// TODO: not implemented yet

}

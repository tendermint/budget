package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tax/x/tax/types"
)

// TODO: not implemented
func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name        string
		configure   func(*types.GenesisState)
		expectedErr string
	}{
		{
			"happy case",
			func(genState *types.GenesisState) {
				params := types.DefaultParams()
				params.Taxes = []types.Tax{}
				genState.Params = params
			},
			"",
		},
		// {
		// 	"invalid case",
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			genState := types.DefaultGenesisState()
			tc.configure(genState)

			err := types.ValidateGenesis(*genState)
			if tc.expectedErr == "" {
				require.Nil(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

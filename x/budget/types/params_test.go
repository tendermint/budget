package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/budget/x/budget/types"
)

func TestParams(t *testing.T) {
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `epoch_blocks: 1
budgets: []
`
	require.Equal(t, paramsStr, defaultParams.String())
}

func TestValidateBudgets(t *testing.T) {
	dAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("destinationAddr1")))
	dAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("destinationAddr2")))
	sAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("sourceAddr1")))
	sAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("sourceAddr2")))
	budgets := []types.Budget{
		{
			Name:               "test",
			Rate:               sdk.NewDec(1),
			SourceAddress:      sAddr1.String(),
			DestinationAddress: dAddr1.String(),
			StartTime:          time.Time{},
			EndTime:            time.Time{},
		},
		{
			Name:               "test2",
			Rate:               sdk.NewDec(1),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          time.Time{},
			EndTime:            time.Time{},
		},
		{
			Name:               "test3",
			Rate:               sdk.MustNewDecFromStr("0.1"),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          time.Time{},
			EndTime:            time.Time{},
		},
		{
			Name:               "test3",
			Rate:               sdk.MustNewDecFromStr("0.1"),
			SourceAddress:      sAddr2.String(),
			DestinationAddress: dAddr2.String(),
			StartTime:          time.Time{},
			EndTime:            time.Time{},
		},
	}

	err := types.ValidateBudgets(budgets[:2])
	require.NoError(t, err)

	err = types.ValidateBudgets(budgets[:3])
	require.ErrorIs(t, err, types.ErrInvalidTotalBudgetRate)

	err = types.ValidateBudgets(budgets)
	require.ErrorIs(t, err, types.ErrDuplicateBudgetName)
}

func TestValidateEpochBlocks(t *testing.T) {
	err := types.ValidateEpochBlocks(uint32(0))
	require.NoError(t, err)

	err = types.ValidateEpochBlocks(nil)
	require.EqualError(t, err, "invalid parameter type: <nil>")

	err = types.ValidateEpochBlocks(types.DefaultEpochBlocks)
	require.NoError(t, err)

	err = types.ValidateEpochBlocks(10000000000000000)
	require.EqualError(t, err, "invalid parameter type: int")
}

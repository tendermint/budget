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
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("budgetSourceAddr2")))
	budgets := []types.Budget{
		{
			Name:                "test",
			Rate:                sdk.NewDec(1),
			BudgetSourceAddress: tAddr1.String(),
			CollectionAddress:   cAddr1.String(),
			StartTime:           time.Time{},
			EndTime:             time.Time{},
		},
		{
			Name:                "test2",
			Rate:                sdk.NewDec(1),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           time.Time{},
			EndTime:             time.Time{},
		},
		{
			Name:                "test3",
			Rate:                sdk.MustNewDecFromStr("0.1"),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           time.Time{},
			EndTime:             time.Time{},
		},
		{
			Name:                "test3",
			Rate:                sdk.MustNewDecFromStr("0.1"),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           time.Time{},
			EndTime:             time.Time{},
		},
	}

	err := types.ValidateBudgets(budgets[:2])
	require.NoError(t, err)

	err = types.ValidateBudgets(budgets[:3])
	require.Error(t, err, types.ErrOverflowedBudgetRate)

	err = types.ValidateBudgets(budgets)
	require.Error(t, err, types.ErrDuplicatedBudgetName)
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

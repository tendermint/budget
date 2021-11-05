package types_test

import (
	"testing"

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
			Rate:                sdk.OneDec(),
			BudgetSourceAddress: tAddr1.String(),
			CollectionAddress:   cAddr1.String(),
			StartTime:           types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-08-03T00:00:00Z"),
		},
		{
			Name:                "test2",
			Rate:                sdk.OneDec(),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           types.MustParseRFC3339("2021-08-03T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-08-05T00:00:00Z"),
		},
		{
			Name:                "test3",
			Rate:                sdk.MustNewDecFromStr("0.1"),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           types.MustParseRFC3339("2021-08-04T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-08-06T00:00:00Z"),
		},
		{
			Name:                "test3",
			Rate:                sdk.MustNewDecFromStr("0.1"),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           types.MustParseRFC3339("2021-07-25T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-07-30T00:00:00Z"),
		},
		{
			Name:                "test4",
			Rate:                sdk.OneDec(),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           types.MustParseRFC3339("2021-08-01T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-08-20T00:00:00Z"),
		},
		{
			Name:                "test5",
			Rate:                sdk.MustNewDecFromStr("0.1"),
			BudgetSourceAddress: tAddr2.String(),
			CollectionAddress:   cAddr2.String(),
			StartTime:           types.MustParseRFC3339("2021-08-19T00:00:00Z"),
			EndTime:             types.MustParseRFC3339("2021-08-25T00:00:00Z"),
		},
	}

	err := types.ValidateBudgets(budgets[:2])
	require.NoError(t, err)

	err = types.ValidateBudgets(budgets[:3])
	require.ErrorIs(t, err, types.ErrInvalidTotalBudgetRate)

	err = types.ValidateBudgets(budgets[3:5])
	require.NoError(t, err)

	err = types.ValidateBudgets(budgets[4:6])
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

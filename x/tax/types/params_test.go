package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tax/x/tax/types"
)

func TestParams(t *testing.T) {
	require.IsType(t, paramstypes.KeyTable{}, types.ParamKeyTable())

	defaultParams := types.DefaultParams()

	paramsStr := `epoch_blocks: 1
taxes: []
`
	require.Equal(t, paramsStr, defaultParams.String())
}

func TestValidateTaxes(t *testing.T) {
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr2")))
	taxes := []types.Tax{
		{
			Name:              "test",
			Rate:              sdk.NewDec(1),
			TaxSourceAddress:  tAddr1.String(),
			CollectionAddress: cAddr1.String(),
			StartTime:         time.Time{},
			EndTime:           time.Time{},
		},
		{
			Name:              "test2",
			Rate:              sdk.NewDec(1),
			TaxSourceAddress:  tAddr2.String(),
			CollectionAddress: cAddr2.String(),
			StartTime:         time.Time{},
			EndTime:           time.Time{},
		},
		{
			Name:              "test3",
			Rate:              sdk.MustNewDecFromStr("0.1"),
			TaxSourceAddress:  tAddr2.String(),
			CollectionAddress: cAddr2.String(),
			StartTime:         time.Time{},
			EndTime:           time.Time{},
		},
		{
			Name:              "test3",
			Rate:              sdk.MustNewDecFromStr("0.1"),
			TaxSourceAddress:  tAddr2.String(),
			CollectionAddress: cAddr2.String(),
			StartTime:         time.Time{},
			EndTime:           time.Time{},
		},
	}

	err := types.ValidateTaxes(taxes[:2])
	require.NoError(t, err)

	err = types.ValidateTaxes(taxes[:3])
	require.Error(t, err, types.ErrOverflowedTaxRate)

	err = types.ValidateTaxes(taxes)
	require.Error(t, err, types.ErrDuplicatedTaxName)
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

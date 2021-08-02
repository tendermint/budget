package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/tax/x/tax/types"
)

func TestValidateTaxes(t *testing.T) {
	cAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr1")))
	cAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("collectionAddr2")))
	tAddr1 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr1")))
	tAddr2 := sdk.AccAddress(address.Module(types.ModuleName, []byte("taxSourceAddr2")))
	taxes := []types.Tax{
		{
			Name:                  "test",
			Rate:                  sdk.NewDec(1),
			CollectionAddress:     cAddr1.String(),
			CollectionAccountName: "",
			TaxSourceAddress:      tAddr1.String(),
			TaxSourceAccountName:  "",
			StartTime:             time.Time{},
			EndTime:               time.Time{},
		},
		{
			Name:                  "test2",
			Rate:                  sdk.NewDec(1),
			CollectionAddress:     cAddr2.String(),
			CollectionAccountName: "",
			TaxSourceAddress:      tAddr2.String(),
			TaxSourceAccountName:  "",
			StartTime:             time.Time{},
			EndTime:               time.Time{},
		},
		{
			Name:                  "test3",
			Rate:                  sdk.MustNewDecFromStr("0.1"),
			CollectionAddress:     cAddr2.String(),
			CollectionAccountName: "",
			TaxSourceAddress:      tAddr2.String(),
			TaxSourceAccountName:  "",
			StartTime:             time.Time{},
			EndTime:               time.Time{},
		},
		{
			Name:                  "test3",
			Rate:                  sdk.MustNewDecFromStr("0.1"),
			CollectionAddress:     cAddr2.String(),
			CollectionAccountName: "",
			TaxSourceAddress:      tAddr2.String(),
			TaxSourceAccountName:  "",
			StartTime:             time.Time{},
			EndTime:               time.Time{},
		},
	}

	err := types.ValidateTaxes(taxes[:2])
	require.NoError(t, err)

	err = types.ValidateTaxes(taxes[:3])
	require.Error(t, err, types.ErrOverflowedTaxRate)

	err = types.ValidateTaxes(taxes)
	require.Error(t, err, types.ErrDuplicatedTaxName)
}

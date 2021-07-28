package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

func (tax Tax) String() string {
	out, _ := tax.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of an Tax.
func (tax Tax) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &tax)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

func (tax Tax) Validate() error {
	//Name is connected to - without spaces and must be within 30 characters
	if len(tax.Name) > MaxTaxNameLength || tax.Name == "" {
		return ErrInvalidTaxName
	}

	// Check that the CollectionAddress is a valid address, if there is a CollectionAccountName value, it must be generated in accordance with Rule ADR-028(moudule-account) and matched to CollectionAddress.
	_, err := ValidityAddrWithName(tax.CollectionAddress, tax.CollectionAccountName)
	if err != nil {
		// TODO: return error with Wrapping
		return err
	}

	// Check that the TaxSourceAddress is a valid address, if there is a TaxSourceAccountName value, it must be generated in accordance with Rule ADR-028(moudule-account) and matched to TaxSourceAddress.
	_, err = ValidityAddrWithName(tax.TaxSourceAddress, tax.TaxSourceAccountName)
	if err != nil {
		// TODO: return error with Wrapping
		return err
	}

	if tax.EndTime.Before(tax.StartTime) {
		return ErrInvalidStartEndTime
	}

	// EndTime should not be earlier than StartTime.
	if tax.Rate == sdk.ZeroDec() {
		return ErrZeroTaxRate
	}
	return nil
}

func ValidityAddrWithName(bech32, name string) (sdk.AccAddress, error) {
	acc, err := sdk.AccAddressFromBech32(bech32)
	if err != nil {
		return nil, err
	}
	if name != "" {
		accFromName := sdk.AccAddress(address.Module(ModuleName, []byte(name)))
		if acc.String() != accFromName.String() {
			return nil, ErrInvalidNameOfAddr
		}
	}
	return acc, nil
}

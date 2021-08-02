package types

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	reTaxNameString = `[a-zA-Z][a-zA-Z0-9-]{0,49}`
	reTaxName       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reTaxNameString))
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
	//Name only allowed letters(`A-Z, a-z`), digits(`0-9`), and `-` without spaces and the maximum length is 50.
	err := ValidateName(tax.Name)
	if err != nil {
		return err
	}

	// Check that the CollectionAddress is a valid address, if there is a CollectionAccountName value, it must be generated in accordance with Rule ADR-028(moudule-account) and matched to CollectionAddress.
	_, err = ValidityAddrWithName(tax.CollectionAddress, tax.CollectionAccountName)
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

	// EndTime should not be earlier than StartTime.
	if tax.EndTime.Before(tax.StartTime) {
		return ErrInvalidStartEndTime
	}

	if tax.Rate == sdk.ZeroDec() {
		return ErrZeroTaxRate
	}
	return nil
}

func (tax Tax) Expired(now time.Time) bool {
	return tax.EndTime.After(now)
}

func ValidityAddrWithName(bech32, name string) (sdk.AccAddress, error) {
	acc, err := sdk.AccAddressFromBech32(bech32)
	if err != nil {
		return nil, err
	}
	if name != "" {
		// TODO: this rule can be fixed, TBD
		accFromName := sdk.AccAddress(address.Module(ModuleName, []byte(name)))
		if acc.String() != accFromName.String() {
			return nil, ErrInvalidNameOfAddr
		}
	}
	return acc, nil
}

// ValidateName is the default validation function for Tax.Name.
func ValidateName(name string) error {
	if !reTaxName.MatchString(name) {
		return sdkerrors.Wrap(ErrInvalidTaxName, name)
	}
	return nil
}

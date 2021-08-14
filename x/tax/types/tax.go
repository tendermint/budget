package types

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	reTaxNameString = fmt.Sprintf(`[a-zA-Z][a-zA-Z0-9-]{0,%d}`, MaxTaxNameLength-1)
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

	// Check that the CollectionAddress is a valid address.
	_, err = ValidityAddr(tax.CollectionAddress)
	if err != nil {
		// TODO: return error with Wrapping
		return err
	}

	// Check that the TaxSourceAddress is a valid address.
	_, err = ValidityAddr(tax.TaxSourceAddress)
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

func (tax Tax) Expired(blockTime time.Time) bool {
	return !tax.EndTime.After(blockTime)
}

func ValidityAddr(bech32 string) (sdk.AccAddress, error) {
	acc, err := sdk.AccAddressFromBech32(bech32)
	if err != nil {
		return nil, err
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

type TaxesByTaxSource struct {
	Taxes     []Tax
	TotalRate sdk.Dec
}

type TaxesByTaxSourceMap map[string]TaxesByTaxSource

// Create a map by TaxSourceAddress to handle the taxes for the same TaxSourceAddress together based on the
// same balance when calculating rates for the same TaxSourceAddress.
func GetTaxesByTaxSourceMap(taxes []Tax) TaxesByTaxSourceMap {
	taxesMap := make(TaxesByTaxSourceMap)
	for _, tax := range taxes {
		if taxesByTaxSource, ok := taxesMap[tax.TaxSourceAddress]; ok {
			taxesByTaxSource.TotalRate = taxesByTaxSource.TotalRate.Add(tax.Rate)
			taxesByTaxSource.Taxes = append(taxesByTaxSource.Taxes, tax)
			taxesMap[tax.TaxSourceAddress] = taxesByTaxSource
		} else {
			taxesMap[tax.TaxSourceAddress] = TaxesByTaxSource{
				Taxes:     []Tax{tax},
				TotalRate: tax.Rate,
			}
		}
	}
	return taxesMap
}

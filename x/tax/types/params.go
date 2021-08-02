package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gopkg.in/yaml.v2"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// MaxTaxNameLength is the maximum length of the name of each tax.
	MaxTaxNameLength int = 50
)

// Parameter store keys
var (
	KeyTaxes = []byte("Taxes")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default tax module parameters.
func DefaultParams() Params {
	return Params{
		Taxes: []Tax{},
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyTaxes, &p.Taxes, ValidateTaxes),
	}
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.Taxes, ValidateTaxes},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

func ValidateTaxes(i interface{}) error {
	taxes, ok := i.([]Tax)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	addrTotalRate := make(map[string]sdk.Dec)
	names := make(map[string]bool)
	for _, tax := range taxes {
		err := tax.Validate()
		if err != nil {
			return err
		}
		if rate, ok := addrTotalRate[tax.TaxSourceAddress]; ok {
			addrTotalRate[tax.TaxSourceAddress] = rate.Add(tax.Rate)
		} else {
			addrTotalRate[tax.TaxSourceAddress] = tax.Rate
		}
		if _, ok := names[tax.Name]; ok {
			return sdkerrors.Wrap(ErrDuplicatedTaxName, tax.Name)
		}
		names[tax.Name] = true
	}
	for addr, totalRate := range addrTotalRate {
		if totalRate.GT(sdk.NewDec(1)) {
			return sdkerrors.Wrap(ErrOverflowedTaxRate, addr)
		}
	}
	return nil
}

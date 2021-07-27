package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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

package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterLegacyAminoCodec registers the necessary x/tax interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

// RegisterInterfaces registers the x/tax interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	//registry.RegisterImplementations(
	//	(*sdk.Msg)(nil),
	//)

	//registry.RegisterImplementations(
	//	(*govtypes.Content)(nil),
	//	&UpdateTaxesProposal{},
	//)

	//registry.RegisterInterface(
	//	"cosmos.tax.v1beta1.TaxI",
	//)

	//msgservice.RegisterMsgServiceDesc(registry, Msg_serviceDesc())
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/tax module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/tax and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

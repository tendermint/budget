package types

import (
	"fmt"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeUpdateTaxes string = "UpdateTaxes"
)

// TODO: it could be replaced to ParamChangeProposal

// Implements Proposal Interface
var _ gov.Content = &UpdateTaxesProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeUpdateTaxes)
	gov.RegisterProposalTypeCodec(&UpdateTaxesProposal{}, "cosmos-sdk/UpdateTaxesProposal")
}

func NewUpdateTaxesProposal(title, description string, taxes []Tax) (gov.Content, error) {
	return &UpdateTaxesProposal{
		Title:       title,
		Description: description,
		Taxes:       taxes,
	}, nil
}

func (p *UpdateTaxesProposal) GetTitle() string { return p.Title }

func (p *UpdateTaxesProposal) GetDescription() string { return p.Description }

func (p *UpdateTaxesProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateTaxesProposal) ProposalType() string { return ProposalTypeUpdateTaxes }

func (p *UpdateTaxesProposal) ValidateBasic() error {
	//for _, tax := range p.Taxes {
	//}
	return gov.ValidateAbstract(p)
}

func (p UpdateTaxesProposal) String() string {
	return fmt.Sprintf(`Update Taxes Proposal:
  Title:       %s
  Description: %s
  Taxes:       %s
`, p.Title, p.Description, p.Taxes)
}

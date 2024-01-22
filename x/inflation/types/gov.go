package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeCreateInflationControlParams = "CreateInflationControlParams"
	ProposalTypeUpdateInflationControlParams = "UpdateInflationControlParams"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateInflationControlParams)
	govtypes.RegisterProposalType(ProposalTypeUpdateInflationControlParams)
}

var (
	_ govtypes.Content = &InflationCreateControlParamsProposal{}
	_ govtypes.Content = &InflationUpdateControlParamsProposal{}
)

func (p *InflationCreateControlParamsProposal) ProposalRoute() string { return RouterKey }

func (p *InflationCreateControlParamsProposal) ProposalType() string {
	return ProposalTypeCreateInflationControlParams
}

func (p *InflationCreateControlParamsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

func (p *InflationUpdateControlParamsProposal) ProposalRoute() string { return RouterKey }

func (p *InflationUpdateControlParamsProposal) ProposalType() string {
	return ProposalTypeUpdateInflationControlParams
}

func (p *InflationUpdateControlParamsProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return nil
}

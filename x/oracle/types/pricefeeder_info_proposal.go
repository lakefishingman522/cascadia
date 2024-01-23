package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	// UpdatePriceFeederInfo defines the type for a UpdatePriceFeederInfoProposal
	UpdatePriceFeederInfo = "UpdatePriceFeederInfo"
)

// Assert UpdatePriceFeederInfoProposal implements govtypes.Content at compile-time
var _ govtypes.Content = &UpdatePriceFeederInfoProposal{}

func init() {
	govtypes.RegisterProposalType(UpdatePriceFeederInfo)
}

func NewUpdatePriceFeederInfoProposal(
	title, description string, priceFeederInfo PriceFeederInfo,
) *UpdatePriceFeederInfoProposal {
	return &UpdatePriceFeederInfoProposal{title, description, priceFeederInfo}
}

// GetTitle returns the title of a update price feeder info proposal.
func (p *UpdatePriceFeederInfoProposal) GetTitle() string { return p.Title }

// GetDescription returns the description of a update price feeder info proposal.
func (p *UpdatePriceFeederInfoProposal) GetDescription() string { return p.Description }

// ProposalRoute returns the routing key of a update price feeder info proposal.
func (*UpdatePriceFeederInfoProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns the type of a update price feeder info proposal.
func (*UpdatePriceFeederInfoProposal) ProposalType() string { return UpdatePriceFeederInfo }

// ValidateBasic validates the update price feeder info proposal.
func (p *UpdatePriceFeederInfoProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	return ValidatePriceFeederInfo(p.PriceFeederInfo)
}

// ValidatePriceFeederInfo performs basic validation checks over a PriceFeederInfo. It
// returns an error if PriceFeederInfo is invalid.
func ValidatePriceFeederInfo(priceFeederInfo PriceFeederInfo) error {

	if priceFeederInfo.Name == "" {
		return ErrInvalidPricefeederName
	}

	_, err := sdk.AccAddressFromBech32(priceFeederInfo.Address)
	if err != nil {
		return ErrInvalidPricefeederAddress
	}

	return nil
}

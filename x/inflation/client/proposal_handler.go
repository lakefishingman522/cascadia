package client

import (
	"github.com/cascadiafoundation/cascadia/x/inflation/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the param change proposal handler.
var CreateInflationControlParamsProposalHandler = govclient.NewProposalHandler(cli.NewSubmitCreateInflationControlParamsProposalTxCmd)
var UpdateInflationControlParamsProposalHandler = govclient.NewProposalHandler(cli.NewSubmitUpdateInflationControlParamsProposalTxCmd)

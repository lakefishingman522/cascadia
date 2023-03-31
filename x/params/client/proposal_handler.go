package client

import (
	govclient "github.com/cascadiafoundation/cascadia/x/gov/client"
	"github.com/cascadiafoundation/cascadia/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)

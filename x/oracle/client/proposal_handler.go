package client

import (
	"github.com/cascadiafoundation/cascadia/x/oracle/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the param change proposal handler.
var PricefeedInfoProposalHandler = govclient.NewProposalHandler(cli.NewSubmitUpdatePriceFeederInfoProposalTxCmd)

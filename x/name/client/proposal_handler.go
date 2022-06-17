package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/MonikaCat/provenance/x/name/client/cli"
	"github.com/MonikaCat/provenance/x/name/client/rest"
)

// ProposalHandler is the create root name proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetRootNameProposalCmd, rest.RootNameProposalHandler)
)

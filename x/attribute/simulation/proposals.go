package simulation

import (
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/MonikaCat/provenance/x/attribute/keeper"
)

// ProposalContents defines the module weighted proposals' contents (none for attribute)
func ProposalContents(k keeper.Keeper) []simtypes.WeightedProposalContent {
	return []simtypes.WeightedProposalContent{}
}

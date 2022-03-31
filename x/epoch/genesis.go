package epoch

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/provenance-io/provenance/x/epoch/keeper"
	"github.com/provenance-io/provenance/x/epoch/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// set epoch info from genesis
	for _, epoch := range genState.Epochs {
		// if epoch has default value, set Start height to current blockheight
		// if set will be honored.
		if epoch.StartHeight == 0 {
			epoch.StartHeight = uint64(ctx.BlockHeight())
		}
		// same for current epoch start height
		epoch.CurrentEpochStartHeight = uint64(ctx.BlockHeight())

		k.SetEpochInfo(ctx, epoch)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Epochs = k.AllEpochInfos(ctx)
	return genesis
}

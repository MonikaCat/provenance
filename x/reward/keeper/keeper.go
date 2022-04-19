package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	"github.com/tendermint/tendermint/libs/log"

	epochkeeper "github.com/provenance-io/provenance/x/epoch/keeper"
	"github.com/provenance-io/provenance/x/reward/types"
)

const StoreKey = types.ModuleName

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryCodec
	EpochKeeper   epochkeeper.Keeper
	stakingKeeper types.StakingKeeper
	govKeeper     *govkeeper.Keeper
	bankKeeper    bankkeeper.Keeper
	authkeeper    authkeeper.AccountKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	epochKeeper epochkeeper.Keeper,
	stakingKeeper types.StakingKeeper,
	govKeeper *govkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	authKeeper authkeeper.AccountKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		EpochKeeper:   epochKeeper,
		stakingKeeper: stakingKeeper,
		govKeeper:     govKeeper,
		bankKeeper:    bankKeeper,
		authkeeper:    authKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetRewardProgram sets the reward program in the keeper
func (k Keeper) SetRewardProgram(ctx sdk.Context, rewardProgram types.RewardProgram) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardProgram)
	store.Set(types.GetRewardProgramKey(int64(rewardProgram.Id)), bz)
}

// GetRewardProgram returns a RewardProgram by id
func (k Keeper) GetRewardProgram(ctx sdk.Context, id int64) (rewardProgram types.RewardProgram, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardProgramKey(id)
	bz := store.Get(key)
	if len(bz) == 0 {
		return rewardProgram, nil
	}
	err = k.cdc.Unmarshal(bz, &rewardProgram)
	return rewardProgram, err
}

// IterateRewardPrograms iterates all reward programs with the given handler function.
func (k Keeper) IterateRewardPrograms(ctx sdk.Context, handle func(rewardProgram types.RewardProgram) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardProgramKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.RewardProgram{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetRewardClaim sets the reward program in the keeper
func (k Keeper) SetRewardClaim(ctx sdk.Context, rewardClaim types.RewardClaim) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardClaim)
	store.Set(types.GetRewardClaimsKey([]byte(rewardClaim.Address)), bz)
}

// GetRewardClaim returns a RewardClaim by id
func (k Keeper) GetRewardClaim(ctx sdk.Context, addr string) (rewardClaim types.RewardClaim, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardClaimsKey([]byte(addr))
	bz := store.Get(key)
	if len(bz) == 0 {
		return rewardClaim, err
	}
	err = k.cdc.Unmarshal(bz, &rewardClaim)
	return rewardClaim, err
}

// IterateRewardClaims  iterates all reward claims with the given handler function.
func (k Keeper) IterateRewardClaims(ctx sdk.Context, handle func(rewardClaim types.RewardClaim) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardClaimKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.RewardClaim{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetEpochRewardDistribution sets the EpochRewardDistribution in the keeper
func (k Keeper) SetEpochRewardDistribution(ctx sdk.Context, epochRewardDistribution types.EpochRewardDistribution) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&epochRewardDistribution)
	store.Set(types.GetEpochRewardDistributionKey(epochRewardDistribution.EpochId, fmt.Sprintf("%d", epochRewardDistribution.RewardProgramId)), bz)
}

// GetEpochRewardDistribution returns a EpochRewardDistribution by epoch id and reward id
func (k Keeper) GetEpochRewardDistribution(ctx sdk.Context, epochId string, rewardId uint64) (epochRewardDistribution types.EpochRewardDistribution, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetEpochRewardDistributionKey(epochId, fmt.Sprintf("%d", rewardId))
	bz := store.Get(key)
	if len(bz) == 0 {
		return epochRewardDistribution, nil
	}
	err = k.cdc.Unmarshal(bz, &epochRewardDistribution)
	return epochRewardDistribution, err
}

// IterateEpochRewardDistributions  iterates all epoch reward distributions with the given handler function.
func (k Keeper) IterateEpochRewardDistributions(ctx sdk.Context, handle func(epochRewardDistribution types.EpochRewardDistribution) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EpochRewardDistributionKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.EpochRewardDistribution{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetEligibilityCriteria sets the reward epoch reward distribution in the keeper
func (k Keeper) SetEligibilityCriteria(ctx sdk.Context, eligibilityCriteria types.EligibilityCriteria) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&eligibilityCriteria)
	store.Set(types.GetEligibilityCriteriaKey(eligibilityCriteria.Name), bz)
}

// GetEligibilityCriteria returns a reward eligibility criteria by name if it exists nil if it does not
func (k Keeper) GetEligibilityCriteria(ctx sdk.Context, name string) (eligibilityCriteria types.EligibilityCriteria, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetEligibilityCriteriaKey(name)
	bz := store.Get(key)
	if len(bz) == 0 {
		return eligibilityCriteria, err
	}
	err = k.cdc.Unmarshal(bz, &eligibilityCriteria)
	return eligibilityCriteria, err
}

// IterateEligibilityCriterias  iterates all reward eligibility criterions with the given handler function.
func (k Keeper) IterateEligibilityCriterias(ctx sdk.Context, handle func(eligibilityCriteria types.EligibilityCriteria) (stop bool)) error {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.EligibilityCriteriaKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		record := types.EligibilityCriteria{}
		if err := k.cdc.Unmarshal(iterator.Value(), &record); err != nil {
			return err
		}
		if handle(record) {
			break
		}
	}
	return nil
}

// SetActionDelegate sets the reward epoch reward distribution in the keeper
func (k Keeper) SetActionDelegate(ctx sdk.Context, actionDelegate types.ActionDelegate) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&actionDelegate)
	store.Set(types.GetActionDelegateKey(), bz)
}

// GetActionDelegate returns a action delegate
func (k Keeper) GetActionDelegate(ctx sdk.Context) (actionDelegate types.ActionDelegate, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetActionDelegateKey())
	if len(bz) == 0 {
		return actionDelegate, nil
	}
	err = k.cdc.Unmarshal(bz, &actionDelegate)
	return actionDelegate, err
}

// SetActionTransferDelegations sets the reward epoch reward distribution in the keeper
func (k Keeper) SetActionTransferDelegations(ctx sdk.Context, actionTransferDelegations types.ActionTransferDelegations) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&actionTransferDelegations)
	store.Set(types.GetActionTransferDelegationsKey(), bz)
}

// GetActionTransferDelegations returns a action transfer delegations
func (k Keeper) GetActionTransferDelegations(ctx sdk.Context) (actionTransferDelegations types.ActionTransferDelegations, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetActionTransferDelegationsKey())
	if len(bz) == 0 {
		return actionTransferDelegations, nil
	}
	err = k.cdc.Unmarshal(bz, &actionTransferDelegations)
	return actionTransferDelegations, err
}

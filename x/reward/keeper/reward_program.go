package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/provenance-io/provenance/x/reward/types"
)

// SetRewardProgram sets the reward program in the keeper
func (k Keeper) SetRewardProgram(ctx sdk.Context, rewardProgram types.RewardProgram) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardProgram)
	store.Set(types.GetRewardProgramKey(rewardProgram.Id), bz)
}

// Removes a reward program in the keeper
func (k Keeper) RemoveRewardProgram(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardProgramKey(id)
	bz := store.Get(key)
	keyExists := store.Has(bz)
	if keyExists {
		store.Delete(bz)
	}
	return keyExists
}

// GetRewardProgram returns a RewardProgram by id
func (k Keeper) GetRewardProgram(ctx sdk.Context, id uint64) (rewardProgram types.RewardProgram, err error) {
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

func (k Keeper) GetOutstandingRewardPrograms(ctx sdk.Context) ([]types.RewardProgram, error) {
	var rewardPrograms []types.RewardProgram
	err := k.IterateRewardPrograms(ctx, func(rewardProgram types.RewardProgram) (stop bool) {
		if !rewardProgram.Finished {
			rewardPrograms = append(rewardPrograms, rewardProgram)
		}
		return false
	})
	return rewardPrograms, err
}

func (k Keeper) GetAllActiveRewardPrograms(ctx sdk.Context) ([]types.RewardProgram, error) {
	var rewardPrograms []types.RewardProgram
	// get all the rewards programs
	err := k.IterateRewardPrograms(ctx, func(rewardProgram types.RewardProgram) (stop bool) {
		if !rewardProgram.Finished && rewardProgram.Started {
			rewardPrograms = append(rewardPrograms, rewardProgram)
		}
		return false
	})
	return rewardPrograms, err
}

// Gets all the RewardPrograms
func (k Keeper) GetAllRewardPrograms(ctx sdk.Context) ([]types.RewardProgram, error) {
	var rewardPrograms []types.RewardProgram
	err := k.IterateRewardPrograms(ctx, func(rewardProgram types.RewardProgram) (stop bool) {
		rewardPrograms = append(rewardPrograms, rewardProgram)
		return false
	})
	if err != nil {
		return nil, err
	}
	return rewardPrograms, nil
}

// Check if a RewardProgram is valid
func (k Keeper) RewardProgramIsValid(rewardProgram *types.RewardProgram) bool {
	return rewardProgram.Id != 0
}

// Removes all RewardPrograms that are expired
func (k Keeper) RemoveExpiredPrograms(ctx sdk.Context) error {
	rewardPrograms, err := k.GetAllRewardPrograms(ctx)
	if err != nil {
		return err
	}
	for _, rewardProgram := range rewardPrograms {
		if !rewardProgram.Finished {
			continue
		}

		k.RemoveRewardProgram(ctx, rewardProgram.GetId())
	}
	return nil
}

// GetRewardProgramID gets the highest rewardprogram ID
func (k Keeper) GetRewardProgramID(ctx sdk.Context) (rewardprogramID uint64, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RewardProgramIDKey)
	if bz == nil {
		return 0, errors.New("initial rewardprogram ID hasn't been set")
	}

	rewardprogramID = types.GetRewardProgramIDFromBytes(bz)
	return rewardprogramID, nil
}

// SetRewardProgramID sets the new rewardprogram ID to the store
func (k Keeper) SetRewardProgramID(ctx sdk.Context, rewardprogramID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RewardProgramIDKey, types.GetRewardProgramIDBytes(rewardprogramID))
}

// SetRewardProgram sets the reward program in the keeper
func (k Keeper) SetRewardProgramBalance(ctx sdk.Context, rewardProgramBalance types.RewardProgramBalance) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewardProgramBalance)
	store.Set(types.GetRewardProgramBalanceKey(rewardProgramBalance.RewardProgramId), bz)
}

// Removes a reward program in the keeper
func (k Keeper) RemoveRewardProgramBalance(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardProgramBalanceKey(id)
	bz := store.Get(key)
	keyExists := store.Has(bz)
	if keyExists {
		store.Delete(bz)
	}
	return keyExists
}

// GetRewardProgram returns a RewardProgram by id
func (k Keeper) GetRewardProgramBalance(ctx sdk.Context, id uint64) (rewardProgramBalance types.RewardProgram, err error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRewardProgramBalanceKey(id)
	bz := store.Get(key)
	if len(bz) == 0 {
		return rewardProgramBalance, nil
	}
	err = k.cdc.Unmarshal(bz, &rewardProgramBalance)
	return rewardProgramBalance, err
}

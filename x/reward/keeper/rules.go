package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/provenance-io/provenance/x/reward/types"
)

type EvaluationResult struct {
	eventTypeToSearch string
	attributeKey      string
	shares            int64
	address           sdk.AccAddress // shares to attribute to this address
}

// EvaluateRules takes in a Eligibility criteria and measure it against the events in the context
func (k Keeper) EvaluateRules(ctx sdk.Context, epochNumber uint64, program types.RewardProgram, distribution types.EpochRewardDistribution) error {
	// get the events from the context history
	switch program.EligibilityCriteria.Action.TypeUrl {
	case proto.MessageName(&types.ActionTransferDelegations{}):
		{
			ctx.Logger().Info(fmt.Sprintf("The Action type is %s", proto.MessageName(&types.ActionTransferDelegations{})))
			// check the event history
			// for transfers event and make sure there is a sender
			evaluateRes, err := k.EvaluateTransferAndCheckDelegation(ctx)
			if err != nil {
				return err
			}
			errorRecordsClaim := k.RecordRewardClaims(ctx, epochNumber, program, distribution, evaluateRes)
			if errorRecordsClaim != nil {
				return errorRecordsClaim
			}
			//set total rewards
			k.SetEpochRewardDistribution(ctx, distribution)

		}
	case proto.MessageName(&types.ActionDelegate{}):
		{
			ctx.Logger().Info(fmt.Sprintf("The Action type is %s", proto.MessageName(&types.ActionDelegate{})))
			// check the event history
			// for transfers event and make sure there is a sender
			evaluateRes, err := k.EvaluateDelegation(ctx)
			if err != nil {
				return err
			}

			errorRecordsClaim := k.RecordRewardClaims(ctx, epochNumber, program, distribution, evaluateRes)
			if errorRecordsClaim != nil {
				return errorRecordsClaim
			}
			//set total rewards
			k.SetEpochRewardDistribution(ctx, distribution)
		}
	default:
		// TODO throw an error or just log it? Leaning towards just logging it for now
		ctx.Logger().Error(fmt.Sprintf("The Action type %s, cannot be evaluated", program.EligibilityCriteria.Action.TypeUrl))
	}
	return nil
}

func (k Keeper) RecordRewardClaims(ctx sdk.Context, epochNumber uint64, program types.RewardProgram, distribution types.EpochRewardDistribution, evaluateRes []EvaluationResult) error {
	// get the address from the eval and check if it has delegation
	// it's an array so should be deterministic
	for _, res := range evaluateRes {
		// add a share to the final total
		// we know the rewards it so update the epoch reward
		distribution.TotalShares = distribution.TotalShares + res.shares
		// add it to the claims
		claim, err := k.GetRewardClaim(ctx, res.address.Bytes())
		if err != nil {
			return err
		}

		if claim.Address == "" {
			found := false
			var mutatedSharesPerEpochRewards []types.SharesPerEpochPerRewardsProgram
			// set a new claim or add to a claim
			for _, rewardClaimForAddress := range claim.SharesPerEpochPerReward {
				if rewardClaimForAddress.RewardProgramId == program.Id && (rewardClaimForAddress.EphemeralActionCount < program.Minimum || rewardClaimForAddress.EphemeralActionCount >= program.Maximum) {
					rewardClaimForAddress.EphemeralActionCount = rewardClaimForAddress.EphemeralActionCount + 1
					mutatedSharesPerEpochRewards = append(mutatedSharesPerEpochRewards, rewardClaimForAddress)
					found = true
				} else if rewardClaimForAddress.RewardProgramId == program.Id && rewardClaimForAddress.EphemeralActionCount < program.Maximum {
					rewardClaimForAddress.EphemeralActionCount = rewardClaimForAddress.EphemeralActionCount + 1
					rewardClaimForAddress.TotalShares = rewardClaimForAddress.TotalShares + 1
					rewardClaimForAddress.LatestRecordedEpoch = epochNumber
					mutatedSharesPerEpochRewards = append(mutatedSharesPerEpochRewards, rewardClaimForAddress)
					found = true
				} else {
					mutatedSharesPerEpochRewards = append(mutatedSharesPerEpochRewards, rewardClaimForAddress)
				}
			}
			if found {
				claim.SharesPerEpochPerReward = mutatedSharesPerEpochRewards
			} else {
				mutatedSharesPerEpochRewards = append(mutatedSharesPerEpochRewards, types.SharesPerEpochPerRewardsProgram{
					RewardProgramId:     program.Id,
					TotalShares:         1,
					LatestRecordedEpoch: epochNumber,
					Claimed:             false,
					Expired:             false,
					TotalRewardClaimed:  sdk.Coin{},
				})
			}
			claim.SharesPerEpochPerReward = mutatedSharesPerEpochRewards
			k.SetRewardClaim(ctx, claim)
		} else {
			//set a brand new claim
			var mutatedSharesPerEpochRewards []types.SharesPerEpochPerRewardsProgram
			k.SetRewardClaim(ctx, types.RewardClaim{
				Address: res.address.String(),
				SharesPerEpochPerReward: append(mutatedSharesPerEpochRewards, types.SharesPerEpochPerRewardsProgram{
					RewardProgramId:     program.Id,
					TotalShares:         1,
					LatestRecordedEpoch: epochNumber,
					Claimed:             false,
					Expired:             false,
					TotalRewardClaimed:  sdk.Coin{},
				}),
			})

		}
	}
	return nil
}

func (k Keeper) EvaluateTransferAndCheckDelegation(ctx sdk.Context) ([]EvaluationResult, error) {
	result := ([]EvaluationResult)(nil)
	evaluateRes, err := k.EvaluateSearchEvents(ctx, "transfer", "sender")
	if err != nil {
		return nil, err
	}
	for _, s := range evaluateRes {
		if len(k.CheckActiveDelegations(ctx, s.address)) > 0 {
			result = append(result, s)
		}
	}
	return result, nil
}

func (k Keeper) EvaluateDelegation(ctx sdk.Context) ([]EvaluationResult, error) {
	evaluateRes, err := k.EvaluateSearchEvents(ctx, "staking", "sender")
	return evaluateRes, err
}

func (k Keeper) EvaluateSearchEvents(ctx sdk.Context, eventTypeToSearch string, attributeKey string) ([]EvaluationResult, error) {
	result := ([]EvaluationResult)(nil)
	for _, s := range ctx.EventManager().GetABCIEventHistory() {
		ctx.Logger().Info(fmt.Sprintf("events type is %s", s.Type))
		if s.Type == eventTypeToSearch {
			// now look for the attribute
			for _, y := range s.Attributes {
				ctx.Logger().Info(fmt.Sprintf("event attribute is %s attribute_key:attribute_value  %s:%s", s.Type, y.Key, y.Value))
				//4:24PM INF events type is transfer
				//4:24PM INF event attribute is transfer attribute_key:attribute_value  recipient:tp17xpfvakm2amg962yls6f84z3kell8c5l2udfyt
				//4:24PM INF event attribute is transfer attribute_key:attribute_value  sender:tp1sha7e07l5knw4vdw2vgc3k06gd0fscz9r32yv6
				//4:24PM INF event attribute is transfer attribute_key:attribute_value  amount:76200000000000nhash
				if attributeKey == string(y.Key) {

					// really not possible to get an error but could happen i guess
					address, err := sdk.AccAddressFromBech32(string(y.Value))

					//TODO check this address has a delegation
					if err != nil {
						return nil, err
					}
					result = append(result, EvaluationResult{
						eventTypeToSearch: eventTypeToSearch,
						attributeKey:      string(y.Key),
						shares:            1,
						address:           address,
					})
				}
			}
		}
	}

	return result, nil

}

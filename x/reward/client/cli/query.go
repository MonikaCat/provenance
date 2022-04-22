package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/provenance-io/provenance/x/reward/types"
)

var cmdStart = fmt.Sprintf("%s query reward", version.AppName)

func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the rewards module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	queryCmd.AddCommand(
		GetRewardProgramCmd(),
		GetRewardClaimCmd(),
	)
	return queryCmd
}

func GetRewardProgramCmd() *cobra.Command {
	const all = "all"
	const active = "active"

	cmd := &cobra.Command{
		Use:     "reward-program {reward_program_id|\"all\"|\"active\"}",
		Aliases: []string{"rp", "rewardprogram"},
		Short:   "Query the current reward programs",
		Long: fmt.Sprintf(`%[1]s reward-program {reward_program_id} - gets the reward program for a given id.
%[1]s reward-program all - gets all the reward programs
%[1]s reward-program active - gets all active the reward programs`, cmdStart),
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(`%[1]s reward-program 1
%[1]s reward-program all
%[1]s reward-program active`, cmdStart),
		RunE: func(cmd *cobra.Command, args []string) error {
			arg0 := strings.TrimSpace(args[0])
			if arg0 == all {
				return outputRewardProgramsAll(cmd)
			} else if arg0 == active {
				return outputRewardProgramsActive(cmd)
			}

			return outputRewardProgramById(cmd, arg0)
		},
	}

	return cmd
}

func GetRewardClaimCmd() *cobra.Command {
	const all = "all"

	cmd := &cobra.Command{
		Use:     "reward-claim {address|\"all\"}",
		Aliases: []string{"rc", "rewardclaim"},
		Short:   "Query the current reward claims",
		Long: fmt.Sprintf(`%[1]s reward-claim {address} - gets the reward claim for the address.
%[1]s reward-claim all - gets all the reward claims`, cmdStart),
		Args: cobra.ExactArgs(1),
		Example: fmt.Sprintf(`%[1]s reward-claim "pb1skjwj5whet0lpe65qaq4rpq03hjxlwd9nf39lk"
%[1]s reward-claim all`, cmdStart),
		RunE: func(cmd *cobra.Command, args []string) error {
			arg0 := strings.TrimSpace(args[0])
			if arg0 == all {
				return outputRewardClaimsAll(cmd)
			}

			return outputRewardClaimById(cmd, arg0)
		},
	}

	return cmd
}

// Query for all Reward Programs
func outputRewardProgramsAll(cmd *cobra.Command) error {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}

	queryClient := types.NewQueryClient(clientCtx)

	var response *types.RewardProgramsResponse
	if response, err = queryClient.RewardPrograms(
		context.Background(),
		&types.RewardProgramsRequest{},
	); err != nil {
		return fmt.Errorf("failed to query reward programs: %s", err.Error())
	}

	return clientCtx.PrintProto(response)
}

// Query for all active Reward Programs
func outputRewardProgramsActive(cmd *cobra.Command) error {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	queryClient := types.NewQueryClient(clientCtx)

	var response *types.ActiveRewardProgramsResponse
	if response, err = queryClient.ActiveRewardPrograms(
		context.Background(),
		&types.ActiveRewardProgramsRequest{},
	); err != nil {
		return fmt.Errorf("failed to query active reward programs: %s", err.Error())
	}
	return clientCtx.PrintProto(response)
}

// Query for a RewardProgram by Id
func outputRewardProgramById(cmd *cobra.Command, arg string) error {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	queryClient := types.NewQueryClient(clientCtx)
	programId, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}

	var response *types.RewardProgramByIDResponse
	if response, err = queryClient.RewardProgramByID(
		context.Background(),
		&types.RewardProgramByIDRequest{Id: uint64(programId)},
	); err != nil {
		return fmt.Errorf("failed to query reward program %d: %s", programId, err.Error())
	}

	if response.GetRewardProgram() == nil {
		return fmt.Errorf("reward program %d does not exist", programId)
	}

	return clientCtx.PrintProto(response)
}

func outputRewardClaimsAll(cmd *cobra.Command) error {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}

	queryClient := types.NewQueryClient(clientCtx)

	var response *types.RewardProgramsResponse
	if response, err = queryClient.RewardPrograms(
		context.Background(),
		&types.RewardProgramsRequest{},
	); err != nil {
		return fmt.Errorf("failed to query reward programs: %s", err.Error())
	}

	return clientCtx.PrintProto(response)
}

func outputRewardClaimById(cmd *cobra.Command, arg string) error {
	return fmt.Errorf("not implemented yet")

	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return err
	}
	queryClient := types.NewQueryClient(clientCtx)
	id, err := strconv.Atoi(arg)
	if err != nil {
		return err
	}

	var response *types.RewardProgramByIDResponse
	if response, err = queryClient.RewardProgramByID(
		context.Background(),
		&types.RewardProgramByIDRequest{Id: uint64(id)},
	); err != nil {
		return fmt.Errorf("failed to query reward claim %d: %s", id, err.Error())
	}

	if response.GetRewardProgram() == nil {
		return fmt.Errorf("reward claim %d does not exist", id)
	}

	return clientCtx.PrintProto(response)
}

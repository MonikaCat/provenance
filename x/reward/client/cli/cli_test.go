package cli_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"

	"github.com/provenance-io/provenance/internal/antewrapper"
	"github.com/provenance-io/provenance/testutil"

	epochtypes "github.com/provenance-io/provenance/x/epoch/types"
	"github.com/provenance-io/provenance/x/reward/types"
	rewardtypes "github.com/provenance-io/provenance/x/reward/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network

	accountAddr sdk.AccAddress
	accountKey  *secp256k1.PrivKey
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")
	s.accountKey = secp256k1.GenPrivKeyFromSecret([]byte("acc2"))
	addr, err := sdk.AccAddressFromHex(s.accountKey.PubKey().Address().String())
	s.Require().NoError(err)
	s.accountAddr = addr

	s.cfg = testutil.DefaultTestNetworkConfig()

	genesisState := s.cfg.GenesisState
	s.cfg.NumValidators = 1

	epochData := epochtypes.NewGenesisState([]epochtypes.EpochInfo{
		{Identifier: "minute",
			StartHeight:             0,
			Duration:                uint64((60) / 5),
			CurrentEpoch:            0,
			CurrentEpochStartHeight: 0,
		},
	})
	epochDataBz, err := s.cfg.Codec.MarshalJSON(epochData)
	s.Require().NoError(err)
	genesisState[epochtypes.ModuleName] = epochDataBz

	rewardData := rewardtypes.NewGenesisState(
		[]rewardtypes.RewardProgram{
			types.NewRewardProgram(
				1,
				s.accountAddr.String(),
				sdk.NewInt64Coin("jackthecat", 1),
				sdk.NewInt64Coin("jackthecat", 2),
				"minute",
				1,
				1,
				rewardtypes.NewEligibilityCriteria("action-name", &rewardtypes.ActionDelegate{}),
				false,
				1,
				2,
			),
		},
		[]rewardtypes.RewardClaim{},
		[]rewardtypes.EpochRewardDistribution{},
		[]rewardtypes.EligibilityCriteria{},
		rewardtypes.ActionDelegate{},
		rewardtypes.ActionTransferDelegations{},
	)

	rewardDataBz, err := s.cfg.Codec.MarshalJSON(rewardData)
	s.Require().NoError(err)
	genesisState[rewardtypes.ModuleName] = rewardDataBz

	s.cfg.GenesisState = genesisState

	s.cfg.ChainID = antewrapper.SimAppChainID

	s.network = network.New(s.T(), s.cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestQueryRewardPrograms() {
	s.Assert().FailNow("not implemented")
	//cli.QueryRewardProgramsCmd()
	// TODO Need a way to create a reward program before these can be implemented
}

func (s *IntegrationTestSuite) TestQueryRewardProgramsById() {
	s.Assert().FailNow("not implemented")
	//cli.RewardProgramByIdCmd()
	// TODO Need a way to create a reward program before these can be implemented
}

func (s *IntegrationTestSuite) TestQueryRewardClaims() {
	s.Assert().FailNow("not implemented")
	//cli.QueryRewardProgramsCmd()
	// TODO Need a way to create a reward claim before these can be implemented
}

func (s *IntegrationTestSuite) TestQueryRewardClaimsById() {
	s.Assert().FailNow("not implemented")
	//cli.QueryRewardProgramsCmd()
	// TODO Need a way to create a reward claim before these can be implemented
}

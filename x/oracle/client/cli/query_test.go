package cli_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
	network "github.com/encichain/enci/testutil/network"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     testnet.Config
	network *testnet.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	cfg.NumValidators = 1

	s.cfg = cfg
	s.network = testnet.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

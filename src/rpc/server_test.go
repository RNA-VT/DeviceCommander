package rpc

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/utilities"
)

type RPCTestSuite struct {
	suite.Suite
}

func (s *RPCTestSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *RPCTestSuite) TestServer() {
	// testServer := NewDCServer()

}

func TestRPCTestSuite(t *testing.T) {
	suite.Run(t, new(RPCTestSuite))
}

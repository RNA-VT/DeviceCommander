package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/utilities"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *TestSuite) TestGenerateRandomNewDevice() {
	testLength := 3
	testNewDevices := GenerateRandomNewDevices(testLength)

	for _, v := range testNewDevices {
		s.Require().NotNil(v, "all of the devices should not be nil")
	}

	s.Equal(len(testNewDevices), testLength, "there should be the correct number of devices")
}

func (s *TestSuite) TestGenerateRandomNewEndpoints() {
	testLength := 3
	tmpDeviceID := uuid.New().String()
	testNewEndpoints := GenerateRandomNewEndpoints(tmpDeviceID, testLength)

	s.Equal(len(testNewEndpoints), testLength, "there should be the correct number of endpoints")

	testNewEndpoint := testNewEndpoints[0]

	s.Equal(testNewEndpoint.DeviceID, tmpDeviceID, "the NewEndpoint should have the correct DeviceID")
}

func (s *TestSuite) GenerateRandomNewParameter() {
	testLength := 3
	testNewParameters := GenerateRandomNewParameter(testLength)

	s.Equal(len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	s.Nil(testNewParameter.EndpointID, "the NewParameter's EndpointID should not be initialized")
}

func (s *TestSuite) GenerateRandomNewParameterForEndpoint() {
	testLength := 3
	tmpEndpointID := uuid.New().String()
	testNewParameters := GenerateRandomNewParameterForEndpoint(tmpEndpointID, testLength)

	s.Equal(len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	s.Equal(testNewParameter.EndpointID, tmpEndpointID, "the NewParameter should have the correct EndpointID")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(s.T(), len(testNewDevices), testLength, "there should be the correct number of devices")
}

func (s *TestSuite) TestGenerateRandomNewEndpoints() {
	testLength := 3
	tmpDeviceID := uuid.New().String()
	testNewEndpoints := GenerateRandomNewEndpoints(tmpDeviceID, testLength)

	assert.Equal(s.T(), len(testNewEndpoints), testLength, "there should be the correct number of endpoints")

	testNewEndpoint := testNewEndpoints[0]

	assert.Equal(s.T(), testNewEndpoint.DeviceID, tmpDeviceID, "the NewEndpoint should have the correct DeviceID")
}

func (s *TestSuite) GenerateRandomNewParameter() {
	testLength := 3
	testNewParameters := GenerateRandomNewParameter(testLength)

	assert.Equal(s.T(), len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	assert.Nil(s.T(), testNewParameter.EndpointID, "the NewParameter's EndpointID should not be initialized")
}

func (s *TestSuite) GenerateRandomNewParameterForEndpoint() {
	testLength := 3
	tmpEndpointID := uuid.New().String()
	testNewParameters := GenerateRandomNewParameterForEndpoint(tmpEndpointID, testLength)

	assert.Equal(s.T(), len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	assert.Equal(s.T(), testNewParameter.EndpointID, tmpEndpointID, "the NewParameter should have the correct EndpointID")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

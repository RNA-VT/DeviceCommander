package endpoint

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/pkg/utilities"
)

type EndpointSuite struct {
	suite.Suite
}

func (s *EndpointSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *EndpointSuite) CreateTestNewEndpoint() NewEndpointParams {
	return GenerateRandomNewEndpointParams(uuid.New().String(), 1)[0]
}

func (s *EndpointSuite) TestNewEndpoint() {
	testNewEndpoint := s.CreateTestNewEndpoint()
	testEndpoint, err := FromNewEndpoint(testNewEndpoint)
	s.Nil(err, "creating a new Endpoint from a NewEndpoint should not throw an error")

	s.NotNil(testEndpoint.Parameters, "the Parameters field should be initialized")

	s.NotNil(testEndpoint.ID, "the endpoint ID should be initialized")

	s.Equal(testNewEndpoint.Description, testEndpoint.Description, "the description should carry through to the NewEndpoint")
}

func (s *EndpointSuite) TestGenerateRandomNewEndpoints() {
	testLength := 3
	tmpDeviceID := uuid.New().String()
	testNewEndpoints := GenerateRandomNewEndpointParams(tmpDeviceID, testLength)

	s.Equal(len(testNewEndpoints), testLength, "there should be the correct number of endpoints")

	testNewEndpoint := testNewEndpoints[0]

	s.Equal(testNewEndpoint.DeviceID, tmpDeviceID, "the NewEndpoint should have the correct DeviceID")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestEndpointSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

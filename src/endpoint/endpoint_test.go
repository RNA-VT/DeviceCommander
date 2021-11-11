package endpoint

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/graph/model"
	"github.com/rna-vt/devicecommander/src/test"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type EndpointSuite struct {
	suite.Suite
}

func (s *EndpointSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *EndpointSuite) CreateTestNewEndpoint() model.NewEndpoint {
	return test.GenerateRandomNewEndpoints(uuid.New().String(), 1)[0]
}

func (s *EndpointSuite) TestNewEndpoint() {
	testNewEndpoint := s.CreateTestNewEndpoint()
	testEndpoint := EndpointFromNewEndpoint(testNewEndpoint)

	assert.NotNil(s.T(), testEndpoint.ID, "the endpoint ID should be initialized")

	assert.Equal(s.T(), len(testNewEndpoint.Parameters), len(testEndpoint.Parameters), "there should be the same amount of parameters in the NewEndpoint as there are in the Endpoint")
}

func (s *EndpointSuite) TestNewParameter() {
	testNewEndpoint := s.CreateTestNewEndpoint()

	testEndpoint := EndpointFromNewEndpoint(testNewEndpoint)

	testNewParameter := test.GenerateRandomNewParameterForEndpoint(testEndpoint.ID.String(), 1)[0]
	newParameter := NewParameterFromNewParameter(testNewParameter)

	assert.NotNil(s.T(), newParameter.ID, "the parameter ID should be initialized")

	assert.Equal(s.T(), testEndpoint.ID, newParameter.EndpointID, "the new parameter should have the correct EndpointID relation")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestEndpointSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

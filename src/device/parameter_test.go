package device

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/utilities"
)

type ParameterSuite struct {
	suite.Suite
}

func (s *ParameterSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *ParameterSuite) TestNewParameter() {
	tmpUUID := uuid.New()
	testNewParameter := GenerateRandomNewParameterForEndpoint(tmpUUID.String(), 1)[0]
	newParameter, err := FromNewParameter(testNewParameter)

	assert.Nil(s.T(), err, "creating a new parameter from a NewParameter should not throw an error")

	assert.NotNil(s.T(), newParameter.ID, "the parameter ID should be initialized")

	assert.Equal(s.T(), tmpUUID, newParameter.EndpointID, "the new parameter should have the correct EndpointID relation")
}

func (s *ParameterSuite) TestNewParameterInvalid() {
	tmpUUID := uuid.New()
	testNewParameter := GenerateRandomNewParameterForEndpoint(tmpUUID.String(), 1)[0]
	testNewParameter.EndpointID = ""

	_, err := FromNewParameter(testNewParameter)
	assert.NotNil(s.T(), err, "creating a new Parameter with a NewParameter with an invalid EndpointID should throw an error")
}

func (s *ParameterSuite) GenerateRandomNewParameter() {
	testLength := 3
	testNewParameters := GenerateRandomNewParameter(testLength)

	s.Equal(len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	s.Nil(testNewParameter.EndpointID, "the NewParameter's EndpointID should not be initialized")
}

func (s *ParameterSuite) GenerateRandomNewParameterForEndpoint() {
	testLength := 3
	tmpEndpointID := uuid.New().String()
	testNewParameters := GenerateRandomNewParameterForEndpoint(tmpEndpointID, testLength)

	s.Equal(len(testNewParameters), testLength, "there should be the correct number of parameters")

	testNewParameter := testNewParameters[0]

	s.Equal(testNewParameter.EndpointID, tmpEndpointID, "the NewParameter should have the correct EndpointID")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestParameterSuite(t *testing.T) {
	suite.Run(t, new(ParameterSuite))
}

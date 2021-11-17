package postgres

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/endpoint"
	"github.com/rna-vt/devicecommander/src/graph/model"
	"github.com/rna-vt/devicecommander/src/parameter"
	"github.com/rna-vt/devicecommander/src/test"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type PostgresParameterRepositorySuite struct {
	suite.Suite
	testDevices         []model.Device
	testEndpoints       []model.Endpoint
	testParameters      []model.Parameter
	deviceRepository    DeviceRepository
	endpointRepository  endpoint.IEndpointCRUDRepository
	parameterRepository parameter.IParameterCRUDRepository
}

func (s *PostgresParameterRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	dbConfig := GetDBConfigFromEnv()

	deviceRepository, err := NewDeviceRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	endpointRepository, err := NewEndpointRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	parameterRepository, err := NewParameterRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.deviceRepository = deviceRepository
	s.endpointRepository = endpointRepository
	s.parameterRepository = parameterRepository

	newDevs := test.GenerateRandomNewDevices(1)
	dev, err := s.deviceRepository.Create(newDevs[0])
	assert.Nil(s.T(), err)

	testEndpoint := test.GenerateRandomNewEndpoints(dev.ID.String(), 1)

	end, err := s.endpointRepository.Create(testEndpoint[0])
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *dev)
	s.testEndpoints = append(s.testEndpoints, *end)
}

func (s *PostgresParameterRepositorySuite) CreateTestParameter() model.Parameter {
	currentTestEndpoint := s.testEndpoints[0]
	testParameters := test.GenerateRandomNewParameterForEndpoint(currentTestEndpoint.ID.String(), 1)

	param, err := s.parameterRepository.Create(testParameters[0])
	assert.Nil(s.T(), err, "creating a test parameter should not throw an error")

	s.testParameters = append(s.testParameters, *param)

	return *param
}

func (s *PostgresParameterRepositorySuite) TestGet() {
	testParameter := s.CreateTestParameter()

	results, err := s.parameterRepository.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, len(results), "there should only be a single return when searching by id")

	assert.Equal(s.T(), &testParameter, results[0], "the return from create should be equal to the return from get")
}

func (s *PostgresParameterRepositorySuite) TestDelete() {
	testParameter := s.CreateTestParameter()

	deleteResult, err := s.parameterRepository.Delete(testParameter.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, testParameter.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.parameterRepository.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 0, len(getResults), "there should be 0 parameters with the ID of the deleted device")
}

func (s *PostgresParameterRepositorySuite) TestUpdate() {
	testParameter := s.CreateTestParameter()

	tmpDesc := "Radom test update"
	err := s.parameterRepository.Update(model.UpdateParameter{
		ID:          testParameter.ID.String(),
		Description: &tmpDesc,
	})
	assert.Nil(s.T(), err)

	getResults, err := s.parameterRepository.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated parameter should have the new description")
}

func (s *PostgresParameterRepositorySuite) TestParamUpdate() {
	testParameter := s.CreateTestParameter()

	tmpDesc := "Radom test update 710"
	paramUpdate := model.UpdateParameter{
		ID:          testParameter.ID.String(),
		Description: &tmpDesc,
	}

	err := s.parameterRepository.Update(paramUpdate)
	assert.Nil(s.T(), err)

	getResults, err := s.parameterRepository.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated parameter should have the new description")
}

func (s *PostgresParameterRepositorySuite) TearDownSuite() {
	for _, p := range s.testParameters {
		_, err := s.parameterRepository.Delete(p.ID.String())
		log.Warn(err)
	}

	for _, e := range s.testEndpoints {
		_, err := s.parameterRepository.Delete(e.ID.String())
		log.Warn(err)
	}

	for _, d := range s.testDevices {
		_, err := s.deviceRepository.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresParameterRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresParameterRepositorySuite))
}

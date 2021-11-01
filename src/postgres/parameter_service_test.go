package postgres

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/test"
	"github.com/rna-vt/devicecommander/utilities"
)

type PostgresParameterServiceSuite struct {
	suite.Suite
	testDevices      []model.Device
	testEndpoints    []model.Endpoint
	testParameters   []model.Parameter
	deviceService    DeviceService
	endpointService  EndpointCRUDService
	parameterService ParameterCRUDService
}

func (s *PostgresParameterServiceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()

	dbConfig := DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}

	deviceService, err := NewDeviceService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	endpointService, err := NewEndpointService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	parameterService, err := NewParameterService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.deviceService = deviceService
	s.endpointService = endpointService
	s.parameterService = parameterService

	newDevs := test.GenerateRandomNewDevices(1)
	dev, err := s.deviceService.Create(newDevs[0])
	assert.Nil(s.T(), err)

	testEndpoint := test.GenerateRandomNewEndpoints(dev.ID.String(), 1)

	end, err := s.endpointService.Create(testEndpoint[0])
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *dev)
	s.testEndpoints = append(s.testEndpoints, *end)
}

func (s *PostgresParameterServiceSuite) CreateTestParameter() model.Parameter {
	currentTestEndpoint := s.testEndpoints[0]
	testParameters := test.GenerateRandomNewParameterForEndpoint(currentTestEndpoint.ID.String(), 1)

	param, err := s.parameterService.Create(testParameters[0])
	assert.Nil(s.T(), err)

	s.testParameters = append(s.testParameters, *param)

	return *param
}

func (s *PostgresParameterServiceSuite) TestGet() {
	testParameter := s.CreateTestParameter()

	results, err := s.parameterService.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, len(results), "there should only be a single return when searching by id")

	assert.Equal(s.T(), &testParameter, results[0], "the return from create should be equal to the return from get")
}

func (s *PostgresParameterServiceSuite) TestDelete() {
	testParameter := s.CreateTestParameter()

	deleteResult, err := s.parameterService.Delete(testParameter.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, testParameter.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.parameterService.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 0, len(getResults), "there should be 0 parameters with the ID of the deleted device")
}

func (s *PostgresParameterServiceSuite) TestUpdate() {
	testParameter := s.CreateTestParameter()

	tmpDesc := "Radom test update"
	err := s.parameterService.Update(model.UpdateParameter{
		ID:          testParameter.ID.String(),
		Description: &tmpDesc,
	})
	assert.Nil(s.T(), err)

	getResults, err := s.parameterService.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated parameter should have the new description")
}

func (s *PostgresParameterServiceSuite) TestParamUpdate() {
	testParameter := s.CreateTestParameter()

	tmpDesc := "Radom test update 710"
	paramUpdate := model.UpdateParameter{
		ID:          testParameter.ID.String(),
		Description: &tmpDesc,
	}

	err := s.parameterService.Update(paramUpdate)
	assert.Nil(s.T(), err)

	getResults, err := s.parameterService.Get(model.Parameter{
		ID: testParameter.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated parameter should have the new description")
}

func (s *PostgresParameterServiceSuite) TearDownSuite() {
	for _, p := range s.testParameters {
		_, err := s.parameterService.Delete(p.ID.String())
		log.Warn(err)
	}

	for _, e := range s.testEndpoints {
		_, err := s.parameterService.Delete(e.ID.String())
		log.Warn(err)
	}

	for _, d := range s.testDevices {
		_, err := s.deviceService.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresParameterServiceSuite(t *testing.T) {
	suite.Run(t, new(PostgresParameterServiceSuite))
}

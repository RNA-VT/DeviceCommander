package parameter

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/device/endpoint"
	"github.com/rna-vt/devicecommander/pkg/device/endpoint/parameter"
	"github.com/rna-vt/devicecommander/pkg/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/pkg/postgres/device"
	postgresEndpoint "github.com/rna-vt/devicecommander/pkg/postgres/endpoint"
	"github.com/rna-vt/devicecommander/pkg/utilities"
)

type PostgresParameterRepositorySuite struct {
	suite.Suite
	testDevices         []device.Device
	testEndpoints       []endpoint.Endpoint
	testParameters      []parameter.Parameter
	deviceRepository    postgresDevice.Repository
	endpointRepository  endpoint.Repository
	parameterRepository parameter.Repository
}

func (s *PostgresParameterRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	dbConfig := postgres.GetDBConfigFromEnv()

	deviceRepository, err := postgresDevice.NewRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	endpointRepository, err := postgresEndpoint.NewRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	parameterRepository, err := NewParameterRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	s.deviceRepository = deviceRepository
	s.endpointRepository = endpointRepository
	s.parameterRepository = parameterRepository

	newDevices := device.GenerateRandomNewDeviceParams(1)
	dev, err := s.deviceRepository.Create(newDevices[0])
	s.Nil(err)

	testEndpoint := endpoint.GenerateRandomNewEndpointParams(dev.ID.String(), 1)

	end, err := s.endpointRepository.Create(testEndpoint[0])
	s.Nil(err)

	s.testDevices = append(s.testDevices, *dev)
	s.testEndpoints = append(s.testEndpoints, *end)
}

func (s *PostgresParameterRepositorySuite) CreateTestParameter() parameter.Parameter {
	currentTestEndpoint := s.testEndpoints[0]
	testParameters := parameter.GenerateRandomNewParameterForEndpoint(currentTestEndpoint.ID.String(), 1)

	param, err := s.parameterRepository.Create(testParameters[0])
	s.Nil(err, "creating a test parameter should not throw an error")

	s.testParameters = append(s.testParameters, *param)

	return *param
}

func (s *PostgresParameterRepositorySuite) TestGet() {
	testParameter := s.CreateTestParameter()

	results, err := s.parameterRepository.Get(parameter.Parameter{
		ID: testParameter.ID,
	})
	s.Nil(err)

	s.Equal(1, len(results), "there should only be a single return when searching by id")

	s.Equal(&testParameter, results[0], "the return from create should be equal to the return from get")
}

func (s *PostgresParameterRepositorySuite) TestDelete() {
	testParameter := s.CreateTestParameter()

	deleteResult, err := s.parameterRepository.Delete(testParameter.ID.String())
	s.Nil(err)

	s.Equal(deleteResult.ID, testParameter.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.parameterRepository.Get(parameter.Parameter{
		ID: testParameter.ID,
	})
	s.Nil(err)

	s.Equal(0, len(getResults), "there should be 0 parameters with the ID of the deleted device")
}

func (s *PostgresParameterRepositorySuite) TestUpdate() {
	testParameter := s.CreateTestParameter()

	tmpDesc := "Radom test update"
	err := s.parameterRepository.Update(parameter.UpdateParameterParams{
		ID:          testParameter.ID.String(),
		Description: &tmpDesc,
	})
	s.Nil(err)

	getResults, err := s.parameterRepository.Get(parameter.Parameter{
		ID: testParameter.ID,
	})
	s.Nil(err)

	s.Equal(tmpDesc, *getResults[0].Description, "the updated parameter should have the new description")
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
		s.Nil(err)
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestPostgresParameterRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresParameterRepositorySuite))
}

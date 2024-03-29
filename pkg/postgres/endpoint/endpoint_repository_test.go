package endpoint

import (
	"testing"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/device/endpoint"
	"github.com/rna-vt/devicecommander/pkg/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/pkg/postgres/device"
	"github.com/rna-vt/devicecommander/pkg/utilities"
)

type PostgresEndpointRepositorySuite struct {
	suite.Suite
	testDevices        []device.Device
	testEndpoints      []endpoint.Endpoint
	endpointRepository endpoint.Repository
	deviceRepository   device.Repository
}

func (s *PostgresEndpointRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	dbConfig := postgres.GetDBConfigFromEnv()

	endpointRepository, err := NewRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	deviceRepository, err := postgresDevice.NewRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	s.endpointRepository = endpointRepository
	s.deviceRepository = deviceRepository

	newDevices := device.GenerateRandomNewDeviceParams(1)
	dev, err := s.deviceRepository.Create(newDevices[0])
	s.Require().Nil(err, "creating a test device should not throw an error")

	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresEndpointRepositorySuite) CreateTestEndpoint() endpoint.Endpoint {
	testEndpoints := endpoint.GenerateRandomNewEndpointParams(s.testDevices[0].ID.String(), 1)
	testEndpoint := testEndpoints[0]

	end, err := s.endpointRepository.Create(testEndpoint)
	s.Nil(err)
	s.testDevices[0].Endpoints = nil

	s.testEndpoints = append(s.testEndpoints, *end)

	return *end
}

func (s *PostgresEndpointRepositorySuite) TestGet() {
	testEndpoint := s.CreateTestEndpoint()

	results, err := s.endpointRepository.Get(endpoint.Endpoint{
		ID: testEndpoint.ID,
	})
	s.Nil(err)

	s.Equal(1, len(results), "there should only be a single return when searching by id")
	s.Equal(testEndpoint, *results[0], "the return from create should be equal to the return from get")

	for _, p := range results[0].Parameters {
		s.Equal(testEndpoint.ID, p.EndpointID, "the new param should have the correct endpoint id")
	}

	s.Equal(len(testEndpoint.Parameters), len(results[0].Parameters), "the endpoint should have the same number of parameters as the new obj")
}

func (s *PostgresEndpointRepositorySuite) TestDelete() {
	testEndpoint := s.CreateTestEndpoint()

	deleteResult, err := s.endpointRepository.Delete(testEndpoint.ID.String())
	s.Nil(err)

	s.Equal(deleteResult.ID, testEndpoint.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.endpointRepository.Get(endpoint.Endpoint{
		ID: testEndpoint.ID,
	})
	s.Nil(err)

	s.Nil(err)

	s.Equal(0, len(getResults), "there should be 0 endpoints with the ID of the deleted endpoint")
}

func (s *PostgresEndpointRepositorySuite) TestUpdate() {
	testEndpoint := s.CreateTestEndpoint()

	tmpDesc := "update random test"
	err := s.endpointRepository.Update(endpoint.UpdateEndpointParams{
		ID:          testEndpoint.ID.String(),
		Description: &tmpDesc,
	})
	s.Nil(err)

	getResults, err := s.endpointRepository.Get(endpoint.Endpoint{
		ID: testEndpoint.ID,
	})
	s.Nil(err)

	s.Equal(tmpDesc, *getResults[0].Description, "the updated endpoint should have the new description")

	s.Equal(testEndpoint.Method, getResults[0].Method, "the updated endpoint's method should remain unchanged")
}

func (s *PostgresEndpointRepositorySuite) TestUpdateNonExistent() {
	tmpDesc := "non existent random test"
	tmpUUID := uuid.New()
	err := s.endpointRepository.Update(endpoint.UpdateEndpointParams{
		ID:          tmpUUID.String(),
		Description: &tmpDesc,
	})

	s.NotNil(err, "updating an endpoint that does not exist should throw an error")
}

func (s *PostgresEndpointRepositorySuite) TearDownSuite() {
	for _, e := range s.testEndpoints {
		_, err := s.endpointRepository.Delete(e.ID.String())
		log.Warn(err)
	}

	for _, d := range s.testDevices {
		_, err := s.deviceRepository.Delete(d.ID.String())
		s.Nil(err)
	}

	s.testDevices = []device.Device{}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestPostgresEndpointRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresEndpointRepositorySuite))
}

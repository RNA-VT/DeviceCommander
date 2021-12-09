package endpoint

import (
	"testing"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/endpoint"
	"github.com/rna-vt/devicecommander/src/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/src/postgres/device"
	"github.com/rna-vt/devicecommander/src/test"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type PostgresEndpointRepositorySuite struct {
	suite.Suite
	testDevices        []model.Device
	testEndpoints      []model.Endpoint
	endpointRepository endpoint.Repository
	deviceRepository   device.Repository
}

func (s *PostgresEndpointRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()

	dbConfig := postgres.GetDBConfigFromEnv()
	endpointRepository, err := NewRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	deviceRepository, err := postgresDevice.NewRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.endpointRepository = endpointRepository
	s.deviceRepository = deviceRepository

	newDevs := test.GenerateRandomNewDevices(1)
	dev, err := s.deviceRepository.Create(newDevs[0])
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresEndpointRepositorySuite) CreateTestEndpoint() model.Endpoint {
	testEndpoints := test.GenerateRandomNewEndpoints(s.testDevices[0].ID.String(), 1)
	testEndpoint := testEndpoints[0]

	end, err := s.endpointRepository.Create(testEndpoint)
	assert.Nil(s.T(), err)

	s.testEndpoints = append(s.testEndpoints, *end)

	return *end
}

func (s *PostgresEndpointRepositorySuite) TestGet() {
	testEndpoint := s.CreateTestEndpoint()

	results, err := s.endpointRepository.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, len(results), "there should only be a single return when searching by id")

	assert.Equal(s.T(), testEndpoint, *results[0], "the return from create should be equal to the return from get")

	for _, p := range results[0].Parameters {
		assert.Equal(s.T(), testEndpoint.ID, p.EndpointID, "the new param should have the correct endpoint id")
	}

	assert.Equal(s.T(), len(testEndpoint.Parameters), len(results[0].Parameters), "the endpoint should have the same number of parameters as the new obj")
}

func (s *PostgresEndpointRepositorySuite) TestDelete() {
	testEndpoint := s.CreateTestEndpoint()

	deleteResult, err := s.endpointRepository.Delete(testEndpoint.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, testEndpoint.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.endpointRepository.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 0, len(getResults), "there should be 0 endpoints with the ID of the deleted endpoint")
}

func (s *PostgresEndpointRepositorySuite) TestUpdate() {
	testEndpoint := s.CreateTestEndpoint()

	tmpDesc := "update random test"
	err := s.endpointRepository.Update(model.UpdateEndpoint{
		ID:          testEndpoint.ID.String(),
		Description: &tmpDesc,
	})
	assert.Nil(s.T(), err)

	getResults, err := s.endpointRepository.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated endpoint should have the new description")

	assert.Equal(s.T(), testEndpoint.Method, getResults[0].Method, "the updated endpoint's method should remain unchanged")
}

func (s *PostgresEndpointRepositorySuite) TestUpdateNonExistent() {
	tmpDesc := "non existent random test"
	tmpUUID := uuid.New()
	err := s.endpointRepository.Update(model.UpdateEndpoint{
		ID:          tmpUUID.String(),
		Description: &tmpDesc,
	})

	assert.NotNil(s.T(), err, "updating an endpoint that does not exist should throw an error")
}

func (s *PostgresEndpointRepositorySuite) TearDownSuite() {
	for _, e := range s.testEndpoints {
		_, err := s.endpointRepository.Delete(e.ID.String())
		log.Warn(err)
	}

	for _, d := range s.testDevices {
		_, err := s.deviceRepository.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}

	s.testDevices = []model.Device{}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresEndpointRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresEndpointRepositorySuite))
}

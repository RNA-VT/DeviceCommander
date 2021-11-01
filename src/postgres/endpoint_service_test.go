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

type PostgresEndpointServiceSuite struct {
	suite.Suite
	testDevices     []model.Device
	testEndpoints   []model.Endpoint
	endpointService EndpointCRUDService
	deviceService   DeviceCRUDService
}

func (s *PostgresEndpointServiceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()

	dbConfig := DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}
	endpointService, err := NewEndpointService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	deviceService, err := NewDeviceService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.endpointService = endpointService
	s.deviceService = deviceService

	newDevs := test.GenerateRandomNewDevices(1)
	dev, err := s.deviceService.Create(newDevs[0])
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresEndpointServiceSuite) CreateTestEndpoint() model.Endpoint {
	testEndpoints := test.GenerateRandomNewEndpoints(s.testDevices[0].ID.String(), 5)
	testEndpoint := testEndpoints[0]

	end, err := s.endpointService.Create(testEndpoint)
	assert.Nil(s.T(), err)

	s.testEndpoints = append(s.testEndpoints, *end)

	return *end
}

func (s *PostgresEndpointServiceSuite) TestGet() {
	testEndpoint := s.CreateTestEndpoint()

	results, err := s.endpointService.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, len(results), "there should only be a single return when searching by id")

	assert.Equal(s.T(), testEndpoint, *results[0], "the return from create should be equal to the return from get")

	for _, p := range results[0].Parameters {
		assert.Equal(s.T(), testEndpoint.ID.String(), p.EndpointID, "the new param should have the correct endpoint id")
	}

	assert.Equal(s.T(), len(testEndpoint.Parameters), len(results[0].Parameters), "the endpoint should have the same number of parameters as the new obj")
}

func (s *PostgresEndpointServiceSuite) TestDelete() {
	testEndpoint := s.CreateTestEndpoint()

	deleteResult, err := s.endpointService.Delete(testEndpoint.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, testEndpoint.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.endpointService.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 0, len(getResults), "there should be 0 endpoints with the ID of the deleted endpoint")
}

func (s *PostgresEndpointServiceSuite) TestUpdate() {
	testEndpoint := s.CreateTestEndpoint()

	tmpDesc := "Radom test update"
	err := s.endpointService.Update(model.UpdateEndpoint{
		ID:          testEndpoint.ID.String(),
		Description: &tmpDesc,
	})
	assert.Nil(s.T(), err)

	getResults, err := s.endpointService.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated device should have the new description")
}

func (s *PostgresEndpointServiceSuite) TestParamUpdate() {
	testEndpoint := s.CreateTestEndpoint()

	tmpDesc := "Radom test update 710"
	paramUpdate := model.UpdateEndpoint{
		ID:          testEndpoint.ID.String(),
		Description: &tmpDesc,
	}

	err := s.endpointService.Update(paramUpdate)
	assert.Nil(s.T(), err)

	getResults, err := s.endpointService.Get(model.Endpoint{
		ID: testEndpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated device should have the new description")
}

func (s *PostgresEndpointServiceSuite) TearDownSuite() {
	for _, e := range s.testEndpoints {
		_, err := s.endpointService.Delete(e.ID.String())
		log.Warn(err)
	}

	for _, d := range s.testDevices {
		_, err := s.deviceService.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}

	s.testDevices = []model.Device{}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresEndpointServiceSuite(t *testing.T) {
	suite.Run(t, new(PostgresEndpointServiceSuite))
}

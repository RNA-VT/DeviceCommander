package postgres

import (
	"fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/test"
	"github.com/rna-vt/devicecommander/utilities"
)

type PostgresEndpointServiceSuite struct {
	suite.Suite
	testDevices   []model.Device
	testEndpoints []model.Endpoint
	service       EndpointCRUDService
	deviceService DeviceCRUDService
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

	s.service = endpointService
	s.deviceService = deviceService

	newDevs := test.GenerateRandomNewDevices(1)
	dev, err := s.deviceService.Create(newDevs[0])
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresEndpointServiceSuite) TestGet() {
	newEndpoints := test.GenerateRandomNewEndpoints(s.testDevices[0].ID.String(), 1)

	endpoint, err := s.service.Create(newEndpoints[0])
	assert.Nil(s.T(), err)

	// add endpoint to test list for deletion after
	s.testEndpoints = append(s.testEndpoints, *endpoint)

	results, err := s.service.Get(model.Endpoint{
		ID: endpoint.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 1, len(results), "there should only be a single return when searching by id")

	assert.Equal(s.T(), endpoint, results[0], "the return from create should be equal to the return from get")
}

func (s *PostgresEndpointServiceSuite) TestDelete() {
	randomEndpoints := test.GenerateRandomNewEndpoints(s.testDevices[0].ID.String(), 1)
	randomEndpoint := randomEndpoints[0]

	end, err := s.service.Create(randomEndpoint)
	assert.Nil(s.T(), err)

	deleteResult, err := s.service.Delete(end.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, end.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.service.Get(model.Endpoint{
		ID: end.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 0, len(getResults), "there should be 0 devices with the ID of the deleted device")
}

func (s *PostgresEndpointServiceSuite) TestUpdate() {
	randomEndpoints := test.GenerateRandomNewEndpoints(s.testDevices[0].ID.String(), 1)
	randomEndpoint := randomEndpoints[0]

	end, err := s.service.Create(randomEndpoint)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testEndpoints = append(s.testEndpoints, *end)

	tmpDesc := "Radom test update"
	err = s.service.Update(model.UpdateEndpoint{
		ID:          end.ID.String(),
		Description: &tmpDesc,
	})
	assert.Nil(s.T(), err)

	// assert.Equal(s.T(), deleteResult.ID, dev.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.service.Get(model.Endpoint{
		ID: end.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), tmpDesc, *getResults[0].Description, "the updated device should have the new description")
}

func (s *PostgresEndpointServiceSuite) AfterTest(_, _ string) {
	for _, e := range s.testEndpoints {
		fmt.Println(e)
		// _, err := s.service.Delete(e.ID.String())
		// assert.Nil(s.T(), err)
	}

	s.testEndpoints = []model.Endpoint{}

	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *PostgresEndpointServiceSuite) TearDownSuite() {
	for _, d := range s.testDevices {
		_, err := s.deviceService.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}

	// s.testDevices = []model.Device{}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresEndpointServiceSuite(t *testing.T) {
	suite.Run(t, new(PostgresEndpointServiceSuite))
}

package postgres

import (
	"log"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/test"
	"github.com/rna-vt/devicecommander/utilities"
)

type PostgresDeviceServiceSuite struct {
	suite.Suite
	testDevices []model.Device
	service     DeviceCRUDService
}

func (s *PostgresDeviceServiceSuite) SetupSuite() {
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

	s.service = deviceService
}

func (s *PostgresDeviceServiceSuite) TestGet() {
	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.service.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)

	results, err := s.service.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(results), 0, "there should only be a single return when searching by id")

	assert.Equal(s.T(), results[0], dev, "the return from create should be equal to the return from get")
}

func (s *PostgresDeviceServiceSuite) TestDelete() {
	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.service.Create(newDev)
	assert.Nil(s.T(), err)

	deleteResult, err := s.service.Delete(dev.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, dev.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.service.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(getResults), 0, "there should be 0 devices with the ID of the deleted device")
}

func (s *PostgresDeviceServiceSuite) TestUpdate() {
	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.service.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)

	tmpMAC := test.GenerateRandomMacAddress()

	err = s.service.Update(model.UpdateDevice{
		ID:  dev.ID.String(),
		Mac: &tmpMAC,
	})
	assert.Nil(s.T(), err)

	// assert.Equal(s.T(), deleteResult.ID, dev.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.service.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), getResults[0].MAC, tmpMAC, "the updated device should have the new MAC address")
}

func (s *PostgresDeviceServiceSuite) AfterTest(_, _ string) {
	for _, d := range s.testDevices {
		_, err := s.service.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}

	s.testDevices = []model.Device{}

	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresDeviceServiceSuite(t *testing.T) {
	suite.Run(t, new(PostgresDeviceServiceSuite))
}

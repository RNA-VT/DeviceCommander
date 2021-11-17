package postgres

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/graph/model"
	"github.com/rna-vt/devicecommander/src/test"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type PostgresDeviceRepositorySuite struct {
	suite.Suite
	testDevices []model.Device
	repository  device.IDeviceCRUDRepository
}

func (s *PostgresDeviceRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	dbConfig := GetDBConfigFromEnv()

	deviceRepository, err := NewDeviceRepository(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.repository = deviceRepository

	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.repository.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresDeviceRepositorySuite) CreateTestDevice() model.Device {
	testDevices := test.GenerateRandomNewDevices(1)
	testDevice := testDevices[0]

	newDevice, err := s.repository.Create(testDevice)
	assert.Nil(s.T(), err)

	s.testDevices = append(s.testDevices, *newDevice)

	return *newDevice
}

func (s *PostgresDeviceRepositorySuite) TestDeviceRepositoryImplementsCRUDInterface() {
	// val := MyType("hello")
	// testDevice := s.CreateTestDevice()
	// _, ok := interface{}(testDevice).(DeviceCRUDRepository)

	// assert.Equal(s.T(), true, ok, "the device repository must implement the DeviceCRUDRepository interface")
}

func (s *PostgresDeviceRepositorySuite) TestGet() {
	testDevice := s.CreateTestDevice()

	results, err := s.repository.Get(model.Device{
		ID: testDevice.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(results), 1, "there should only be a single return when searching by id")

	assert.Equal(s.T(), len(results[0].Endpoints), len(testDevice.Endpoints), "the return from get should have the same number of endpoints as the original object")

	assert.Equal(s.T(), *results[0], testDevice, "the return from create should be equal to the return from get")
}

func (s *PostgresDeviceRepositorySuite) TestDelete() {
	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.repository.Create(newDev)
	assert.Nil(s.T(), err)

	deleteResult, err := s.repository.Delete(dev.ID.String())
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), deleteResult.ID, dev.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.repository.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(getResults), 0, "there should be 0 devices with the ID of the deleted device")
}

func (s *PostgresDeviceRepositorySuite) TestUpdate() {
	newDevs := test.GenerateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.repository.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)

	tmpMAC := test.GenerateRandomMacAddress()

	err = s.repository.Update(model.UpdateDevice{
		ID:  dev.ID.String(),
		Mac: &tmpMAC,
	})
	assert.Nil(s.T(), err)

	getResults, err := s.repository.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), getResults[0].MAC, tmpMAC, "the updated device should have the new MAC address")
}

func (s *PostgresDeviceRepositorySuite) AfterTest(_, _ string) {
	for _, d := range s.testDevices {
		_, err := s.repository.Delete(d.ID.String())
		assert.Nil(s.T(), err)
	}

	s.testDevices = []model.Device{}

	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPostgresDeviceRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresDeviceRepositorySuite))
}

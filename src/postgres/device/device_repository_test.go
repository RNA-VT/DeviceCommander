package device

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
	"github.com/rna-vt/devicecommander/src/utilities"
)

type PostgresDeviceRepositorySuite struct {
	suite.Suite
	testDevices []device.Device
	repository  device.Repository
}

func (s *PostgresDeviceRepositorySuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	dbConfig := postgres.GetDBConfigFromEnv()

	deviceRepository, err := NewRepository(dbConfig)
	s.Require().Nil(err, "connecting to the DB should not throw an error")

	s.repository = deviceRepository

	dev, err := s.repository.Create(device.GenerateRandomNewDeviceParams(1)[0])
	s.Nil(err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)
}

func (s *PostgresDeviceRepositorySuite) CreateTestDevice() device.Device {
	testDevices := device.GenerateRandomNewDeviceParams(1)
	testDevice := testDevices[0]

	newDevice, err := s.repository.Create(testDevice)
	s.Nil(err)

	s.testDevices = append(s.testDevices, *newDevice)

	return *newDevice
}

func (s *PostgresDeviceRepositorySuite) TestGet() {
	testDevice := s.CreateTestDevice()

	results, err := s.repository.Get(device.Device{
		ID: testDevice.ID,
	})
	s.Nil(err)

	s.Equal(len(results), 1, "there should only be a single return when searching by id")

	s.Equal(len(results[0].Endpoints), len(testDevice.Endpoints), "the return from get should have the same number of endpoints as the original object")

	s.Equal(*results[0], testDevice, "the return from create should be equal to the return from get")
}

func (s *PostgresDeviceRepositorySuite) TestDelete() {
	testDevice := s.CreateTestDevice()

	deleteResult, err := s.repository.Delete(testDevice.ID.String())
	s.Nil(err, "there should be no error when deleting a newly created device")

	s.Equal(deleteResult.ID, testDevice.ID, "the return from a delete should contain the deleted object")

	getResults, err := s.repository.Get(device.Device{
		ID: testDevice.ID,
	})
	s.Nil(err)

	s.Equal(len(getResults), 0, "there should be 0 devices with the ID of the deleted device")
}

func (s *PostgresDeviceRepositorySuite) TestUpdate() {
	testDevice := s.CreateTestDevice()

	tmpMAC := faker.MacAddress()
	err := s.repository.Update(device.UpdateDeviceParams{
		ID:  testDevice.ID.String(),
		MAC: &tmpMAC,
	})
	s.Nil(err)

	getResults, err := s.repository.Get(device.Device{
		ID: testDevice.ID,
	})
	s.Nil(err)

	s.Equal(getResults[0].MAC, tmpMAC, "the updated device should have the new MAC address")
}

func (s *PostgresDeviceRepositorySuite) TestUpdateNonExistent() {
	tmpMAC := faker.MacAddress()
	tmpUUID := uuid.New()
	err := s.repository.Update(device.UpdateDeviceParams{
		ID:  tmpUUID.String(),
		MAC: &tmpMAC,
	})

	s.NotNil(err, "updating a device that does not exist should throw an error")
}

func (s *PostgresDeviceRepositorySuite) AfterTest(_, _ string) {
	for _, d := range s.testDevices {
		_, err := s.repository.Delete(d.ID.String())
		if err != nil {
			log.Warn(err)
		}
	}

	s.testDevices = []device.Device{}

	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestPostgresDeviceRepositorySuite(t *testing.T) {
	suite.Run(t, new(PostgresDeviceRepositorySuite))
}

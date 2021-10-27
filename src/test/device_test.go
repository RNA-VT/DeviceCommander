package test

import (
	"crypto/rand"
	"fmt"
	"log"
	"testing"

	"github.com/rna-vt/devicecommander/graph/model"
	p "github.com/rna-vt/devicecommander/postgres"
	"github.com/rna-vt/devicecommander/utilities"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeviceSuite struct {
	suite.Suite
	testDevices []model.Device
	service     p.DeviceCRUDService
}

func generateRandomMacAddress() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	// Set the local bit
	buf[0] |= 2
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

func generateRandomNewDevices(count int) []model.NewDevice {
	collection := []model.NewDevice{}
	for i := 0; i < count; i++ {
		tmpMac := generateRandomMacAddress()
		tmpDev := model.NewDevice{
			Mac: &tmpMac,
		}
		collection = append(collection, tmpDev)
	}
	return collection
}

func (s *DeviceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()

	dbConfig := p.DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}
	deviceService, err := p.NewDeviceService(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	s.service = deviceService
}

func (s *DeviceSuite) TestGet() {
	newDevs := generateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.service.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)

	results, err := s.service.Get(model.Device{
		ID: dev.ID,
	})
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), len(results), 1, "there should only be a single return when searching by id")

	assert.Equal(s.T(), results[0], dev, "the return from create should be equal to the return from get")
}

func (s *DeviceSuite) TestDelete() {
	newDevs := generateRandomNewDevices(1)
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

func (s *DeviceSuite) TestUpdate() {
	newDevs := generateRandomNewDevices(1)
	newDev := newDevs[0]

	dev, err := s.service.Create(newDev)
	assert.Nil(s.T(), err)

	// add device to test list for deletion after
	s.testDevices = append(s.testDevices, *dev)

	tmpMAC := generateRandomMacAddress()

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

func (s *DeviceSuite) AfterTest(_, _ string) {
	for _, d := range s.testDevices {
		s.service.Delete(d.ID.String())
	}

	s.testDevices = []model.Device{}

	// require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeviceTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceSuite))
}

package device

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/test"
	"github.com/rna-vt/devicecommander/utilities"
)

type DeviceServiceSuite struct {
	suite.Suite
}

func (s *DeviceServiceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *DeviceServiceSuite) TestNewDeviceFromNewDeviceWith() {
	testNewDevice := test.GenerateRandomNewDevices(1)[0]

	newDeviceResult := NewDeviceFromNewDevice(testNewDevice)

	assert.Equal(s.T(), newDeviceResult.MAC, *testNewDevice.Mac, "the MAC address is properly assigned")

	assert.Equal(s.T(), newDeviceResult.Name, *testNewDevice.Name, "the Name is properly assigned")
	assert.Equal(s.T(), newDeviceResult.Description, *testNewDevice.Description, "the Description is properly assigned")

	assert.NotNil(s.T(), newDeviceResult.Endpoints, "the Endpoints array is initialized")
}

func (s *DeviceServiceSuite) TestNewDeviceWrapper() {
	testNewDevice := test.GenerateRandomNewDevices(1)[0]

	newDeviceResult := NewDeviceFromNewDevice(testNewDevice)

	wrapper := NewDeviceWrapper(&newDeviceResult)

	assert.Equal(s.T(), wrapper.Device.ID, newDeviceResult.ID, "the wrapper should have the same ID as the Device")
}

func (s *DeviceServiceSuite) TestNewDeviceFromRequestBody() {
	testNewDevice := test.GenerateRandomNewDevices(1)[0]

	b, err := json.Marshal(testNewDevice)
	assert.Nil(s.T(), err)

	r := ioutil.NopCloser(strings.NewReader(string(b))) // r type is io.ReadCloser

	newDevice, err := NewDeviceFromRequestBody(r)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), testNewDevice, newDevice, "the device should remain unchanged")
}

func (s *DeviceServiceSuite) TestDeviceURL() {
	testNewDevice := test.GenerateRandomNewDevices(1)[0]

	newDeviceResult := NewDeviceFromNewDevice(testNewDevice)

	wrapper := NewDeviceWrapper(&newDeviceResult)

	// validate the URL
	_, err := url.ParseRequestURI(wrapper.URL())
	assert.Nil(s.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeviceServiceSuite(t *testing.T) {
	suite.Run(t, new(DeviceServiceSuite))
}

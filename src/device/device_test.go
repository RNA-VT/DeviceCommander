package device

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/utilities"
)

const testDeviceResponse = `{
	"Name":"something or other",
	"Endpoints": [
		{
			"Method": "attack",
			"Parameters": [
				{
					"Name": "weapon"
				},
				{
					"Name": "inTheNameOf"
				}
			]
		},
		{
			"Method": "defend",
			"Parameters": [
				{
					"Name": "bodyPart"
				}
			]
		}
	]
}`

type DeviceServiceSuite struct {
	suite.Suite
}

func (s *DeviceServiceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *DeviceServiceSuite) TestDeviceFromNewDeviceWith() {
	testNewDevice := GenerateRandomNewDeviceParams(1)[0]

	newDeviceResult := FromNewDevice(testNewDevice)

	assert.Equal(s.T(), newDeviceResult.MAC, *testNewDevice.Mac, "the MAC address is properly assigned")

	assert.NotNil(s.T(), newDeviceResult.Endpoints, "the Endpoints array is initialized")
}

func (s *DeviceServiceSuite) TestNewDeviceFromRequestBody() {
	testNewDevice := GenerateRandomNewDeviceParams(1)[0]

	b, err := json.Marshal(testNewDevice)
	assert.Nil(s.T(), err)

	r := ioutil.NopCloser(strings.NewReader(string(b))) // r type is io.ReadCloser

	newDevice, err := NewDeviceFromRequestBody(r)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), testNewDevice, newDevice, "the device should remain unchanged")
}

func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func (s *DeviceServiceSuite) TestEvaluateSpecificationResponse() {
	assert.Equal(s.T(), true, isJSON(testDeviceResponse), "the test json should be valid json")

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testDeviceResponse)
	}))
	mockServerURL, err := url.Parse(mockServer.URL)
	assert.Nil(s.T(), err, "the mock server should have a valid URL")

	host, port, err := net.SplitHostPort(mockServerURL.Host)
	assert.Nil(s.T(), err, "splitting the mock server Host should not throw errors")

	mockServerPort, err := strconv.Atoi(port)
	assert.Nil(s.T(), err, "the mock server should have an int port")

	client := HTTPDeviceClient{}

	testDevice := Device{
		Host: host,
		Port: mockServerPort,
	}

	resp, err := client.Specification(testDevice)
	assert.Nil(s.T(), err, "requesting a mock spec should not throw an error")

	dev, err := client.EvaluateSpecificationResponse(resp)
	assert.Nil(s.T(), err, "evaluating a json string response should not throw an error")

	assert.Equal(s.T(), "something or other", dev.Name, "the Name in the json string should be applied to the Device")

	assert.Equal(s.T(), 2, len(dev.Endpoints), "there should be 2 Endpoints in the return")

	assert.Equal(s.T(), 2, len(dev.Endpoints[0].Parameters), "there should be 2 Parameters in the first Endpoint")
}

func (s *DeviceServiceSuite) TestDeviceURL() {
	testNewDeviceParams := GenerateRandomNewDeviceParams(1)[0]

	testDevice := FromNewDevice(testNewDeviceParams)

	devURL := testDevice.URL()
	// validate the URL
	_, err := url.ParseRequestURI(devURL)
	assert.Nil(s.T(), err)

	lastChar := devURL[len(devURL)-1:]

	assert.NotEqual(s.T(), lastChar, "/", "the URL should not end in a \"/\"")
}

func (s *DeviceServiceSuite) TestGenerateRandomNewDevice() {
	testLength := 3
	testNewDevices := GenerateRandomNewDeviceParams(testLength)

	for _, v := range testNewDevices {
		s.Require().NotNil(v, "all of the devices should not be nil")
	}

	s.Equal(len(testNewDevices), testLength, "there should be the correct number of devices")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestDeviceServiceSuite(t *testing.T) {
	suite.Run(t, new(DeviceServiceSuite))
}

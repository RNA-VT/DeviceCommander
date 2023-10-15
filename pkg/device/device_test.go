package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/pkg/utilities"
	"github.com/rna-vt/devicecommander/pkg/utils"
)

type DeviceServiceSuite struct {
	suite.Suite
}

func (s *DeviceServiceSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
}

func (s *DeviceServiceSuite) TestDeviceFromNewDeviceWith() {
	testNewDevice := GenerateRandomNewDeviceParams(1)[0]

	newDeviceResult := FromNewDevice(testNewDevice)

	s.Equal(newDeviceResult.MAC, *testNewDevice.MAC, "the MAC address is properly assigned")

	s.NotNil(newDeviceResult.Endpoints, "the Endpoints array is initialized")
}

func (s *DeviceServiceSuite) TestNewDeviceFromRequestBody() {
	testNewDevice := GenerateRandomNewDeviceParams(1)[0]

	b, err := json.Marshal(testNewDevice)
	s.Nil(err)

	r := io.NopCloser(strings.NewReader(string(b))) // r type is io.ReadCloser

	newDevice, err := NewDeviceFromRequestBody(r)
	s.Nil(err)

	s.Equal(testNewDevice, newDevice, "the device should remain unchanged")
}

func (s *DeviceServiceSuite) TestURLGeneration() {
	testDevice := Device{
		Host: "192.168.1.43",
		Port: 80,
	}
	url := testDevice.URL()

	s.Equal("http://192.168.1.43:80", url, "the URL should be properly generated")
}

// A compliant MCU must be running at this IP in order for this test to pass.
// func (s *DeviceServiceSuite) TestDeviceSpecificationRealMCU() {
// 	deviceClient := NewHTTPDeviceClient()

// 	testDevice := Device{
// 		Host: "192.168.1.43",
// 		Port: 80,
// 	}

// 	spec, err := deviceClient.Specification(testDevice)
// 	s.Require().Nil(err, "requesting a specification from a running device should not throw an error")

// 	fmt.Println(utils.PrettyPrintJSON(spec))
// }

func (s *DeviceServiceSuite) TestDeviceSpecification() {
	deviceClient := NewHTTPDeviceClient()

	testSpec := Specification{
		DeviceType: "relay",
		ID:         "1234",
		MAC:        "00:00:00:00:00:00",
	}

	_, host, port := s.NewHTTPMockServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, utils.PrintJSON(testSpec))
	}))

	testDevice := Device{
		Host: host,
		Port: port,
	}

	spec, err := deviceClient.Specification(testDevice)
	s.Require().Nil(err, "requesting a specification from a running device should not throw an error")

	s.Equal(testSpec, spec, "the specification should be properly returned")
}

func (s *DeviceServiceSuite) TestDeviceURL() {
	testNewDeviceParams := GenerateRandomNewDeviceParams(1)[0]

	testDevice := FromNewDevice(testNewDeviceParams)

	devURL := testDevice.URL()
	// validate the URL
	_, err := url.ParseRequestURI(devURL)
	s.Nil(err)

	lastChar := devURL[len(devURL)-1:]

	s.NotEqual(lastChar, "/", "the URL should not end in a \"/\"")
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

// Helpers

func (s *DeviceServiceSuite) NewHTTPMockServer(handler http.Handler) (mockServer *httptest.Server, host string, port int) {
	mockServer = httptest.NewServer(handler)
	mockServerURL, err := url.Parse(mockServer.URL)
	s.Nil(err, "the mock server should have a valid URL")

	host, portString, err := net.SplitHostPort(mockServerURL.Host)
	s.Nil(err, "splitting the mock server Host should not throw errors")

	portInt, err := strconv.Atoi(portString)
	s.Nil(err, "the mock server should have an int port")

	return mockServer, host, portInt
}

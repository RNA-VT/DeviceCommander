package registration

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	mockdevice "github.com/rna-vt/devicecommander/mocks/device"
	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/scanner"
	"github.com/rna-vt/devicecommander/pkg/utilities"
)

type RegistrarSuite struct {
	suite.Suite
	mockDeviceRepository mockdevice.Repository
	mockDeviceClient     mockdevice.Client
	registrar            Registrar
}

func (s *RegistrarSuite) SetupSuite() {
	utilities.ConfigureEnvironment()
	s.mockDeviceRepository = mockdevice.Repository{}
	s.mockDeviceClient = mockdevice.Client{}

	s.registrar = NewDeviceRegistrar(
		&s.mockDeviceClient,
		&s.mockDeviceRepository,
	)
}

func GenerateDevices(count int) []*device.Device {
	devices := device.GenerateRandomNewDeviceParams(count)
	collection := []*device.Device{}
	for _, d := range devices {
		tmpDev := device.FromNewDevice(d)
		collection = append(collection, &tmpDev)
	}
	return collection
}

func (s *RegistrarSuite) TestHandleProspectUnknown() {
	testDevice := GenerateDevices(1)[0]
	foundDevice := scanner.FoundDevice{
		MAC:  testDevice.MAC,
		IP:   testDevice.Host,
		Port: testDevice.Port,
	}

	testSpecification := device.Specification{
		DeviceType: "relay",
	}

	s.mockDeviceClient.On("Specification", mock.AnythingOfType("device.Device")).Return(testSpecification, nil)

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", device.Device{
		MAC: foundDevice.MAC,
	}).Return([]*device.Device{}, nil).Once()

	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(testDevice, nil).Once()

	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	d, err := s.registrar.HandleProspects([]scanner.FoundDevice{foundDevice})
	s.Equal(err, nil)
	s.Equal(foundDevice.IP, d[0].Host)
	s.Equal(foundDevice.MAC, d[0].MAC)
	s.Equal(foundDevice.Port, d[0].Port)
}

func (s *RegistrarSuite) TestHandleProspectsAlreadyExists() {
	testDevice := GenerateDevices(1)[0]
	foundDevice := scanner.FoundDevice{
		MAC:  testDevice.MAC,
		IP:   testDevice.Host,
		Port: testDevice.Port,
	}

	testSpecification := device.Specification{
		DeviceType: "relay",
	}

	s.mockDeviceClient.On("Specification", mock.AnythingOfType("device.Device")).Return(testSpecification, nil)

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{
		testDevice,
	}, nil).Twice()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("device.UpdateDeviceParams")).Return(nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	d, err := s.registrar.HandleProspects([]scanner.FoundDevice{foundDevice})
	s.Equal(err, nil)
	s.Equal(foundDevice.IP, d[0].Host)
	s.Equal(foundDevice.MAC, d[0].MAC)
	s.Equal(foundDevice.Port, d[0].Port)
}

func TestRegistrarSuite(t *testing.T) {
	suite.Run(t, new(RegistrarSuite))
}

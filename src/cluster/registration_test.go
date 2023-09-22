package cluster

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/scanner"
)

func (s *ClusterSuite) TestHandleDiscoveredDevice() {
	devices := GenerateDevices(2)
	foundDevices := make([]scanner.FoundDevice, len(devices))
	for i, d := range devices {
		foundDevices[i] = scanner.FoundDevice{
			MAC:  d.MAC,
			IP:   d.Host,
			Port: d.Port,
		}
	}

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	d, err := s.cluster.HandleDiscoveredDevice(foundDevices[0])
	s.Equal(err, nil)
	s.Equal(foundDevices[0].IP, d.Host)
	s.Equal(foundDevices[0].MAC, d.MAC)
	s.Equal(foundDevices[0].Port, d.Port)
}

func (s *ClusterSuite) TestHandleDiscoveredDeviceAlreadyExists() {
	devices := GenerateDevices(2)
	foundDevices := make([]scanner.FoundDevice, len(devices))
	for i, d := range devices {
		foundDevices[i] = scanner.FoundDevice{
			MAC:  d.MAC,
			IP:   d.Host,
			Port: d.Port,
		}
	}

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{
		devices[0],
	}, nil).Twice()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("device.UpdateDeviceParams")).Return(nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	d, err := s.cluster.HandleDiscoveredDevice(foundDevices[0])
	s.Equal(err, nil)
	s.Equal(foundDevices[0].IP, d.Host)
	s.Equal(foundDevices[0].MAC, d.MAC)
	s.Equal(foundDevices[0].Port, d.Port)
}

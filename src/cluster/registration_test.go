package cluster

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/rna-vt/devicecommander/src/device"
)

func (s *ClusterSuite) TestHandleDiscoveredDevice() {
	devices := GenerateDevices(2)
	newDevices := make([]device.NewDeviceParams, len(devices))
	for i, d := range devices {
		newDevices[i] = device.NewDeviceParams{
			Mac:  &d.MAC,
			Host: d.Host,
			Port: d.Port,
		}
	}

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	d, err := s.cluster.HandleDiscoveredDevice(newDevices[0])
	s.Equal(err, nil)
	s.Equal(newDevices[0].Host, d.Host)
	s.Equal(*newDevices[0].Mac, d.MAC)
	s.Equal(newDevices[0].Port, d.Port)
}

func (s *ClusterSuite) TestVerifyDeviceAPI() {
	devices := GenerateDevices(2)

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{
		devices[0],
	}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("device.UpdateDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("Specification", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	s.mockDeviceClient.On("EvaluateSpecificationResponse", (*http.Response)(nil)).Return(*devices[0], nil)
	_, verified := s.cluster.VerifyDeviceAPI(*devices[0])
	s.Equal(true, verified)
}

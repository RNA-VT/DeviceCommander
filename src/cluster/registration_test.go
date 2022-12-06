package cluster

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/rna-vt/devicecommander/src/device"
)

func (s *ClusterSuite) TestHandleDiscoveredHealthyDevice() {
	devices, newDevices := getNewDevices(2)

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

func (s *ClusterSuite) TestHandleDiscoveredUnhealthyDevice() {
	devices, newDevices := getNewDevices(2)

	// New, Unhealthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(false)
	d, err := s.cluster.HandleDiscoveredDevice(newDevices[0])
	s.Equal(err, nil)
	s.Equal(newDevices[0].Host, d.Host)
	s.Equal(*newDevices[0].Mac, d.MAC)
	s.Equal(newDevices[0].Port, d.Port)
}

func (s *ClusterSuite) TestVerifyDeviceAPISuccess() {
	devices := GenerateDevices(2)

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{
		devices[0],
	}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("device.UpdateDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	s.mockDeviceClient.On("GetSpecificationFromDevice", mock.AnythingOfType("device.Device")).Return(devices[0], nil)
	_, err := s.cluster.VerifyDeviceAPI(*devices[0])
	s.Equal(nil, err)
}

func (s *ClusterSuite) TestVerifyDeviceAPIFailure() {
	devices := GenerateDevices(2)

	// New, Healthy Device that returns an imparseable spec
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("device.Device")).Return([]*device.Device{
		devices[0],
	}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("device.NewDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("device.UpdateDeviceParams")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.Device")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.Device")).Return(true)
	s.mockDeviceClient.On("GetSpecificationFromDevice", mock.AnythingOfType("device.Device")).Return(devices[0], device.ErrSpecificationFailedToDecode)
	_, err := s.cluster.VerifyDeviceAPI(*devices[0])
	s.NotEqual(nil, err)
}

func getNewDevices(count int) ([]*device.Device, []device.NewDeviceParams) {
	devices := GenerateDevices(count)
	newDevices := make([]device.NewDeviceParams, len(devices))
	for i, d := range devices {
		newDevices[i] = device.NewDeviceParams{
			Mac:  &d.MAC,
			Host: d.Host,
			Port: d.Port,
		}
	}
	return devices, newDevices
}

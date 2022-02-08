package cluster

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/rna-vt/devicecommander/graph/model"
)

func (s *ClusterSuite) TestHandleDiscoveredDevice() {
	devices := GenerateDevices(2)
	newDevices := make([]model.NewDevice, len(devices))
	for i, d := range devices {
		newDevices[i] = model.NewDevice{
			Mac:  &d.MAC,
			Host: d.Host,
			Port: d.Port,
		}
	}

	// New, Healthy Device
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("model.Device")).Return([]*model.Device{}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("model.NewDevice")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.BasicDevice")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.BasicDevice")).Return(true)
	d, err := s.cluster.HandleDiscoveredDevice(newDevices[0])
	s.Equal(err, nil)
	s.Equal(newDevices[0].Host, d.Host)
	s.Equal(*newDevices[0].Mac, d.MAC)
	s.Equal(newDevices[0].Port, d.Port)
}

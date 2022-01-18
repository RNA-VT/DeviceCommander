package cluster

import (
	"net/http"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/stretchr/testify/mock"
)

func (s *ClusterSuite) TestHandleDiscoveredDevice() {
	devices := GenerateDevices(2)
	newDevices := make([]model.NewDevice, len(devices))
	for i, d := range devices {
		newDevices[i] = model.NewDevice{
			Mac:         &d.MAC,
			Name:        &d.Name,
			Description: &d.Description,
			Host:        d.Host,
			Port:        d.Port,
		}
	}

	// New, Healthy Device
	original := newDevices[0]
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("model.Device")).Return([]*model.Device{}, nil).Once()
	s.mockDeviceRepository.On("Create", mock.AnythingOfType("model.NewDevice")).Return(devices[0], nil).Once()
	s.mockDeviceClient.On("Health", mock.AnythingOfType("device.BasicDevice")).Return(nil, nil)
	s.mockDeviceClient.On("EvaluateHealthCheckResponse", (*http.Response)(nil), mock.AnythingOfType("device.BasicDevice")).Return(true)
	err := s.cluster.HandleDiscoveredDevice(newDevices[0])
	s.Assertions.Equal(err, nil)

	//Update Device
	update := newDevices[1]
	update.Mac = original.Mac
	updatedDevice := devices[0]
	updatedDevice.MAC = *original.Mac
	s.mockDeviceRepository.On("Get", mock.AnythingOfType("model.Device")).Return([]*model.Device{updatedDevice}, nil).Once()
	s.mockDeviceRepository.On("Update", mock.AnythingOfType("model.UpdateDevice")).Return(nil)
	err = s.cluster.HandleDiscoveredDevice(newDevices[1])
	s.Assertions.Equal(err, nil)

}

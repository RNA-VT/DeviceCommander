package memory

import (
	"github.com/google/uuid"
	"github.com/rna-vt/devicecommander/graph/model"
)

type InMemoryDeviceService struct {
	devices []*model.Device
}

func (s InMemoryDeviceService) Initialise() (InMemoryDeviceService, error) {
	s.devices = []*model.Device{}

	// s.Initialized = true

	return s, nil
}

func (s InMemoryDeviceService) Create(newDeviceArgs model.NewDevice) (*model.Device, error) {
	logger := getMemoryLogger()
	newDevice := model.Device{
		ID:   uuid.New(),
		Host: newDeviceArgs.Host,
		Port: newDeviceArgs.Port,
	}

	if newDeviceArgs.Mac != nil {
		newDevice.MAC = *newDeviceArgs.Mac
	}

	if newDeviceArgs.Name != nil {
		newDevice.Name = *newDeviceArgs.Name
	}

	if newDeviceArgs.Description != nil {
		newDevice.Description = *newDeviceArgs.Description
	}

	s.devices = append(s.devices, &newDevice)

	logger.Debug("Created device " + newDevice.ID.String())
	return &newDevice, nil
}

func (s InMemoryDeviceService) Update(input model.UpdateDevice) (*model.Device, error) {
	return &model.Device{}, nil
}

func (s InMemoryDeviceService) Delete(id string) (*model.Device, error) {
	return &model.Device{}, nil
}

func (s InMemoryDeviceService) Get(devQuery model.Device) ([]*model.Device, error) {
	devices := []*model.Device{}

	return devices, nil
}

func (s InMemoryDeviceService) GetAll() ([]*model.Device, error) {
	return s.devices, nil
}

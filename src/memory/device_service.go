package memory

import (
	"github.com/google/uuid"

	"github.com/rna-vt/devicecommander/graph/model"
)

var deviceCollection []*model.Device

func init() {
	deviceCollection = []*model.Device{}
}

type InMemoryDeviceService struct {
	devices []*model.Device
}

func (s InMemoryDeviceService) Initialise() (InMemoryDeviceService, error) {
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

	deviceCollection = append(deviceCollection, &newDevice)

	logger.Debug("Created device " + newDevice.ID.String())
	return &newDevice, nil
}

func (s InMemoryDeviceService) Update(input model.UpdateDevice) error {
	return nil
}

func remove(s []*model.Device, i int) []*model.Device {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s InMemoryDeviceService) Delete(id string) (*model.Device, error) {
	for index, device := range deviceCollection {
		if device.ID.String() == id {
			deviceCollection = remove(deviceCollection, index)
			return device, nil
		}
	}
	return &model.Device{}, nil
}

func (s InMemoryDeviceService) Get(devQuery model.Device) ([]*model.Device, error) {
	devices := []*model.Device{}

	for _, device := range deviceCollection {
		if device.ID == devQuery.ID {
			devices = append(devices, device)
		}
	}

	return devices, nil
}

func (s InMemoryDeviceService) GetAll() ([]*model.Device, error) {
	return s.devices, nil
}

package endpoint

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
)

type Endpoint interface {
	Execute(map[string]interface{}) error
}

// DeviceEndpoint implements the Endpoint interface. It provides a functional
// layer for interacting with a specific Device's endpoint.
type DeviceEndpoint struct {
	model.Endpoint
	Device device.Device
}

// FromNewEndpoint generates an Endpoint from a NewEndpoint with the correctly
// instantiated fields. This should be the primary way in which an Endpoint is generated.
func FromNewEndpoint(input model.NewEndpoint) (model.Endpoint, error) {
	deviceUUID, err := uuid.Parse(input.DeviceID)
	if err != nil {
		log.Error(err)
	}

	end := model.Endpoint{
		ID:         uuid.New(),
		DeviceID:   deviceUUID,
		Type:       input.Type,
		Method:     input.Method,
		Parameters: []model.Parameter{},
	}

	if input.Description != nil {
		end.Description = input.Description
	}

	return end, nil
}

// NewDeviceEndpoint generates a DeviceEndpoint
func NewDeviceEndpoint() *DeviceEndpoint {
	return &DeviceEndpoint{}
}

// Execute carries out the action associated with the Device's endpoint
// by communicating with the device.
func (e DeviceEndpoint) Execute(map[string]interface{}) error {
	log.Println(e.Method)
	log.Println(e.Device.URL())
	return nil
}

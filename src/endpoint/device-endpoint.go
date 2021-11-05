package endpoint

import (
	"log"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/graph/model"
)

type Endpoint interface {
	Execute(map[string]interface{}) error
}

type Parameter struct {
	model.Parameter
}

// DeviceEndpoint implements the Endpoint interface. It provides a functional
// layer for interacting with a specific Device's endpoint.
type DeviceEndpoint struct {
	model.Endpoint
	Device device.Device
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

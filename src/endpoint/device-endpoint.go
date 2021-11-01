package endpoint

import (
	"log"

	"github.com/rna-vt/devicecommander/device"
	"github.com/rna-vt/devicecommander/graph/model"
)

type Endpoint interface {
	Execute(map[string]interface{}) error
}

type Parameter struct {
	model.Parameter
}

type DeviceEndpoint struct {
	model.Endpoint
	Device device.Device
}

func NewDeviceEndpoint() *DeviceEndpoint {
	return &DeviceEndpoint{}
}

func (e DeviceEndpoint) Execute(map[string]interface{}) error {
	log.Println(e.Method)
	log.Println(e.Device.URL())
	return nil
}

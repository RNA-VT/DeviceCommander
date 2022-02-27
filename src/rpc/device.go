package rpc

import (
	"errors"
	"log"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
)

// The rpc.Device struct encapsulates all of the procedures available for a device.
// In other words this is the Device RPC handler.
type Device struct{}

type AddDeviceResponse struct {
	Data model.Device
}

func (d *Device) Add(payload model.NewDevice, response *AddDeviceResponse) error {
	log.Println("Device.Add")
	newDevice := device.FromNewDevice(payload)
	response.Data = newDevice

	return errors.New("big time error")
}

func (d *Device) Remove(payload model.NewDevice, reply *model.Device) error {
	return nil
}

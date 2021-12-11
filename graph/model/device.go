package model

import (
	"github.com/google/uuid"
)

// Device represents a compliant physical component & its web address.
//
// ID -- the serial nummber of the connecting device
// Name - Optional Device Nickname
// Description - Optional text describing this device
// Host - Device Api Host
// Port - Device Api Port. Set to 443 for https
//
// swagger:model
type Device struct {
	// the UUID for the device.
	//
	// required: true
	ID uuid.UUID `json:"ID"`

	// the MAC address for this device.
	// required: true
	MAC string `json:"MAC" gorm:"unique"`

	// the human readable name of the device.
	// required: false
	Name string `json:"Name"`

	// the description of the device.
	// required: false
	Description string `json:"Description"`

	// the host address of the device.
	// required: true
	Host string `json:"Host"`

	// the active port of the device.
	// required: true
	Port int `json:"Port"`

	// the count of failed actions by the device.
	// required: false
	Failures int `json:"Failures"`

	// a flag representing the responsiveness of the device.
	// required: false
	Active bool `json:"Active"`

	// a list of endpoints available for quering on a device.
	// required: false
	Endpoints []Endpoint `json:"Endpoints"`
}

package model

import (
	"github.com/google/uuid"
)

// Device represents a compliant physical component & its web address.
//
// ID -- the serial number of the connecting device
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
	ID uuid.UUID `json:"ID" faker:"uuid_hyphenated"`

	// the MAC address for this device.
	// required: true
	MAC string `json:"MAC" gorm:"unique" faker:"mac_address"`

	// the human readable name of the device.
	// required: false
	Name string `json:"Name"`

	// the description of the device.
	// required: false
	Description string `json:"Description"`

	// the host address of the device.
	// required: true
	Host string `json:"Host" faker:"ipv4"`

	// the active port of the device.
	// required: true
	Port int `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`

	// the count of failed actions by the device.
	// required: false
	Failures int `json:"Failures" faker:"boundary_start=0, boundary_end=5"`

	// a flag representing the responsiveness of the device.
	// required: false
	Active bool `json:"Active"`

	// a list of endpoints available for quering on a device.
	// required: false
	Endpoints []Endpoint `json:"Endpoints" faker:"-"`
}

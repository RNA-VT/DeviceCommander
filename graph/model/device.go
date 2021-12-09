package model

import (
	"github.com/google/uuid"
)

// Device represents a compliant physical component & its web address.
// ID -- the serial nummber of the connecting device
// Name - Optional Device Nickname
// Description - Optional text describing this device
// Host - Device Api Host
// Port - Device Api Port. Set to 443 for https
type Device struct {
	ID          uuid.UUID  `json:"ID"`
	MAC         string     `json:"MAC" gorm:"unique"`
	Name        string     `json:"Name"`
	Description string     `json:"Description"`
	Host        string     `json:"Host"`
	Port        int        `json:"Port"`
	Failures    int        `json:"Failures"`
	Active      bool       `json:"Active"`
	Endpoints   []Endpoint `json:"Endpoints"`
}

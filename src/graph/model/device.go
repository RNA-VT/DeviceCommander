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
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	Failures    int       `json:"failures"`
}

package device

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type DeviceEndpoint interface {
	Execute(map[string]interface{}) error
}

// DeviceEndpoint implements the Endpoint interface. It provides a functional
// layer for interacting with a specific Device's endpoint.
type Endpoint struct {
	// model.Endpoint
	ID          uuid.UUID   `json:"ID" faker:"uuid_hyphenated"`
	DeviceID    uuid.UUID   `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      string      `json:"Method"`
	Type        string      `json:"Type" faker:"oneof: get, set"`
	Description *string     `json:"Description"`
	Path        *string     `json:"Path"`
	Parameters  []Parameter `json:"Parameters,omitempty"`
	Device      Device
}

// FromNewEndpoint generates an Endpoint from a NewEndpoint with the correctly
// instantiated fields. This should be the primary way in which an Endpoint is generated.
func FromNewEndpoint(input NewEndpointParams) (Endpoint, error) {
	deviceUUID, err := uuid.Parse(input.DeviceID)
	if err != nil {
		log.Error(err)
	}

	end := Endpoint{
		ID:         uuid.New(),
		DeviceID:   deviceUUID,
		Type:       input.Type,
		Method:     input.Method,
		Parameters: []Parameter{},
	}

	if input.Description != nil {
		end.Description = input.Description
	}

	return end, nil
}

// NewDeviceEndpoint generates a DeviceEndpoint.
func NewDeviceEndpoint() *Endpoint {
	return &Endpoint{}
}

// Execute carries out the action associated with the Device's endpoint
// by communicating with the device.
func (e Endpoint) Execute(map[string]interface{}) error {
	log.Println(e.Method)
	log.Println(e.Device.URL())
	return nil
}

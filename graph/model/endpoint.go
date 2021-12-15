package model

import "github.com/google/uuid"

type Endpoint struct {
	ID          uuid.UUID   `json:"ID" faker:"uuid_hyphenated"`
	DeviceID    uuid.UUID   `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      string      `json:"Method"`
	Type        string      `json:"Type" faker:"oneof: get, set"`
	Description *string     `json:"Description"`
	Path        *string     `json:"Path"`
	Parameters  []Parameter `json:"Parameters,omitempty"`
}

type Parameter struct {
	ID          uuid.UUID `json:"ID" faker:"uuid_hyphenated"`
	EndpointID  uuid.UUID `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        string    `json:"Name"`
	Description *string   `json:"Description"`
	Type        string    `json:"Type" faker:"oneof: string, int, bool"`
}

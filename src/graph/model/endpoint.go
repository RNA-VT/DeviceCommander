package model

import "github.com/google/uuid"

type Endpoint struct {
	ID          uuid.UUID   `json:"ID"`
	DeviceID    uuid.UUID   `json:"DeviceID"`
	Method      string      `json:"Method"`
	Type        string      `json:"Type"`
	Description *string     `json:"Description"`
	Path        *string     `json:"Path"`
	Parameters  []Parameter `json:"Paramaters,omitempty"`
}

type Parameter struct {
	ID          uuid.UUID `json:"ID"`
	EndpointID  uuid.UUID `json:"EndpointID"`
	Name        string    `json:"Name"`
	Description *string   `json:"Description"`
	Type        string    `json:"Type"`
}

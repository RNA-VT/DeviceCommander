// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewDevice struct {
	Mac         *string `json:"MAC"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        string  `json:"Host"`
	Port        int     `json:"Port"`
	Active      *bool   `json:"Active"`
}

type NewEndpoint struct {
	DeviceID    string          `json:"DeviceID"`
	Method      string          `json:"Method"`
	Type        string          `json:"Type"`
	Description *string         `json:"Description"`
	Parameters  []*NewParameter `json:"Parameters"`
}

type NewParameter struct {
	EndpointID  string  `json:"EndpointID"`
	Name        string  `json:"Name"`
	Description *string `json:"Description"`
	Type        string  `json:"Type"`
}

type UpdateDevice struct {
	ID          string  `json:"ID"`
	Mac         *string `json:"MAC"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        *string `json:"Host"`
	Port        *int    `json:"Port"`
	Active      *bool   `json:"Active"`
}

type UpdateEndpoint struct {
	ID          string  `json:"ID"`
	DeviceID    *string `json:"DeviceID"`
	Method      *string `json:"Method"`
	Type        *string `json:"Type"`
	Description *string `json:"Description"`
}

type UpdateParameter struct {
	ID          string  `json:"ID"`
	EndpointID  *string `json:"EndpointID"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Type        *string `json:"Type"`
}

package model

// Device models.
type NewDevice struct {
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        string  `json:"Host" faker:"ipv4"`
	Port        int     `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`
	Active      *bool   `json:"Active"`
}

type UpdateDevice struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        *string `json:"Host" faker:"ipv4"`
	Port        *int    `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`
	Active      *bool   `json:"Active"`
}

// Endpoint models.
type NewEndpoint struct {
	DeviceID    string  `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      string  `json:"Method"`
	Type        string  `json:"Type" faker:"oneof: get, set"`
	Description *string `json:"Description"`
}

type UpdateEndpoint struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	DeviceID    *string `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      *string `json:"Method"`
	Type        *string `json:"Type" faker:"oneof: get, set"`
	Description *string `json:"Description"`
}

// Parameter models.
type NewParameter struct {
	EndpointID  string  `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        string  `json:"Name"`
	Description *string `json:"Description"`
	Type        string  `json:"Type" faker:"oneof: string, int, bool"`
}

type UpdateParameter struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	EndpointID  *string `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Type        *string `json:"Type" faker:"oneof: string, int, bool"`
}

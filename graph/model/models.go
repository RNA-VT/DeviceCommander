package model

type NewDevice struct {
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        string  `json:"Host" faker:"ipv4"`
	Port        int     `json:"Port"`
	Active      *bool   `json:"Active"`
}

type NewEndpoint struct {
	DeviceID    string  `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      string  `json:"Method"`
	Type        string  `json:"Type"`
	Description *string `json:"Description"`
}

type NewParameter struct {
	EndpointID  string  `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        string  `json:"Name"`
	Description *string `json:"Description"`
	Type        string  `json:"Type"`
}

type UpdateDevice struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        *string `json:"Host" faker:"ipv4"`
	Port        *int    `json:"Port"`
	Active      *bool   `json:"Active"`
}

type UpdateEndpoint struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	DeviceID    *string `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      *string `json:"Method"`
	Type        *string `json:"Type"`
	Description *string `json:"Description"`
}

type UpdateParameter struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	EndpointID  *string `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Type        *string `json:"Type"`
}

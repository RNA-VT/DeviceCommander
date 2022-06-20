package device

// Device models.
type NewDeviceParams struct {
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        string  `json:"Host" faker:"ipv4"`
	Port        int     `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`
	Active      *bool   `json:"Active"`
}

type UpdateDeviceParams struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	Mac         *string `json:"MAC" faker:"mac_address"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Host        *string `json:"Host" faker:"ipv4"`
	Port        *int    `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`
	Active      *bool   `json:"Active"`
}

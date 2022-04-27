package device

type NewParameterParams struct {
	EndpointID  string  `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        string  `json:"Name"`
	Description *string `json:"Description"`
	Type        string  `json:"Type" faker:"oneof: string, int, bool"`
}

type UpdateParameterParams struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	EndpointID  *string `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Type        *string `json:"Type" faker:"oneof: string, int, bool"`
}

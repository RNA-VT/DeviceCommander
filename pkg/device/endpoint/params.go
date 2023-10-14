package endpoint

type NewEndpointParams struct {
	DeviceID    string  `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      string  `json:"Method"`
	Type        string  `json:"Type" faker:"oneof: get, set"`
	Description *string `json:"Description"`
}

type UpdateEndpointParams struct {
	ID          string  `json:"ID" faker:"uuid_hyphenated"`
	DeviceID    *string `json:"DeviceID" faker:"uuid_hyphenated"`
	Method      *string `json:"Method"`
	Type        *string `json:"Type" faker:"oneof: get, set"`
	Description *string `json:"Description"`
}

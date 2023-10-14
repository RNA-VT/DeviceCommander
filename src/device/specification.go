package device

type Specification struct {
	ID         string `json:"id"`
	DeviceType string `json:"type"`
	Endpoints  []struct {
		Method     string `json:"method"`
		Parameters []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"parameters"`
	} `json:"endpoints"`
}

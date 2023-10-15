package device

type Specification struct {
	ID         string `json:"id"`
	DeviceType string `json:"deviceType"`
	MAC        string `json:"mac"`
	Endpoints  []struct {
		Method     string `json:"method"`
		Parameters []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"parameters"`
	} `json:"endpoints"`
}

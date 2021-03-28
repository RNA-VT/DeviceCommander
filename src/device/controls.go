package device

import "encoding/json"

type ControlType string

const (
	// On/Off Devices
	Binary ControlType = "binary"
	// Devices whose health we care about but are not controlled through DC
	None ControlType = "none"
	// Invalid Device
	Error ControlType = "error"
)

type ControlConfiguration struct {
	Type        string
	Description string
}

func ValidateControlType(t string) bool {
	var controlType ControlType
	b, err := json.Marshal(t)
	if err != nil {
		return false
	}

	err = json.Unmarshal(b, &controlType)
	if err != nil {
		return false
	}
	return true
}

package component

import "firecontroller/utilities"

/*BaseComponent object definition */
type BaseComponent struct {
	UID       int
	Enabled   bool
	Name      string                 `yaml:"name"`
	HeaderPin int                    `yaml:"header_pin"`
	Metadata  map[string]interface{} `yaml:"metadata"`
}

func (b *BaseComponent) String() string {
	jsonString, err := utilities.StringJSON(b)
	if err != nil {
		return "FailedtoStringify:" + b.Name
	}
	return jsonString
}

//Component is an interface shared by the components we can control with a raspi
type Component interface {
	Init() error
	String() string
	State() string
	Enable(bool)
	Disable()
}

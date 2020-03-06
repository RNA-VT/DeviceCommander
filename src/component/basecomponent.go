package component

/*BaseComponent object definition */
type BaseComponent struct {
	UID       int
	Enabled   bool
	Name      string                 `yaml:"name"`
	HeaderPin int                    `yaml:"header_pin"`
	Metadata  map[string]interface{} `yaml:"metadata"`
}

//Component is an interface shared by the components we can control with a raspi
type Component interface {
	Init() error
	Enable(bool) (err error)
	Disable()
	String() string
	State() string
	Healthy() bool
	Edit(map[string]interface{}) bool
}

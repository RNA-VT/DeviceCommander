package component

/*BaseComponent object definition */
type BaseComponent struct {
	UID       int
	Enabled   bool
	Name      string `yaml:"name"`
	HeaderPin int    `yaml:"header_pin"`
}

//Component is an interface shared by the components we can control with a raspi
type Component interface {
	Init() error
	String() string
	State() string
	Enable(bool)
	Disable()
}

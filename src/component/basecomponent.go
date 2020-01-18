package component

/*BaseComponent object definition */
type BaseComponent struct {
	UID       int
	Enabled   bool
	Name      string `yaml:"name"`
	HeaderPin int    `yaml:"header_pin"`
}

/*
GetBase allows us to share the properties in this type across all components
	and forces components to obey the Component interface{}
*/
func (c *BaseComponent) GetBase() *BaseComponent {
	return c
}

//Component is an interface shared by the components we can control with a raspi
type Component interface {
	Init() error
	String() string
	State() string
	Enable(bool)
	Disable()
}

package component

import (
	// "fmt"
	// "encoding/json"
	// "strings"
	"strconv"
)

/*BaseComponent object definition */
type BaseComponent struct {
	UID     int
	Name    string
	Enabled bool
	OnState bool
}

//Enable - make this component available to command
func (c *BaseComponent) Enable(restoreState bool) {
	c.Enabled = true
}

//Disable - force this component to an off or safe state and make it unavaible to command
func (c *BaseComponent) Disable() {
	c.Enabled = false
	c.OnState = false
}

/*CurrentStateSting just for pretty printing the device info */
func (c *BaseComponent) CurrentStateSting() string {
	state := "OFF"

	if c.OnState {
		state = "ON"
	}

	message := "[" + strconv.Itoa(c.UID) + "] is " + state
	return message
}

package component

import (
	// "fmt"
	// "encoding/json"
	// "strings"
	"strconv"

	"github.com/stianeikeland/go-rpio/v4"
)

/*BaseComponent object definition */
type BaseComponent struct {
	UID     int
	Name    string
	Enabled bool
	GPIO    Gpio
}

//Enable - make this component available to command
func (c *BaseComponent) Enable(restoreState bool) {
	c.Enabled = true
	c.GPIO.InitializePin(3, false)
}

//Disable - force this component to an off or safe state and make it unavaible to command
func (c *BaseComponent) Disable() {
	c.Enabled = false
	c.GPIO.Pin.Low()
}

/*CurrentStateSting just for pretty printing the device info */
func (c *BaseComponent) CurrentStateSting() string {
	state := "OFF"

	if c.GPIO.Pin.Read() == rpio.High {
		state = "ON"
	}

	message := "[" + strconv.Itoa(c.UID) + "] is " + state
	return message
}

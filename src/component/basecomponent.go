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
	UID       int
	Enabled   bool
	Name      string `yaml:"name"`
	HeaderPin int    `yaml:"header_pin"`
	GPIO      Gpio   `yaml:"gpio"`
}

//Enable - make this component available to command
func (c *BaseComponent) Enable(init bool) {
	c.Enabled = true
	if init {
		c.GPIO.InitializePin(c.HeaderPin, false)
	}
}

//Disable - force this component to an off or safe state and make it unavaible to command
func (c *BaseComponent) Disable() {
	c.Enabled = false
	c.GPIO.Pin.Low()
}

/*CurrentStateString just for pretty printing the device info */
func (c *BaseComponent) CurrentStateString() string {
	state := "OFF"

	if c.GPIO.Pin.Read() == rpio.High {
		state = "ON"
	}

	message := "[" + strconv.Itoa(c.UID) + "] is " + state
	return message
}

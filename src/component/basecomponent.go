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
	OnState bool
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

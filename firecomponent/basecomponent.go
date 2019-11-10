package firecomponent

import (
  // "fmt"
  // "encoding/json"
  // "strings"
  "strconv"
)

/* Information/Metadata about node */
type BaseComponent struct {
  UID int
  Name string
  OnState bool

}

/* Just for pretty printing the node info */
func (c BaseComponent) CurrentStateSting() string {
  state := "OFF"

  if c.OnState {
    state = "ON"
  }

  message := "[" + strconv.Itoa(c.UID) + "] is " + state
  return message
}

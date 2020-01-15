package component

import (
	"errors"
	"strconv"
	"strings"

	"github.com/stianeikeland/go-rpio/v4"
)

//Gpio - a raspberrypi digital gpio pin
type Gpio struct {
	Pin     rpio.Pin
	PinInfo RpiPin
	Failed  bool
}

//InitializePin - create gpio pin object and set modes
func (g *Gpio) InitializePin(headerPin int, initHigh bool) error {
	if err := g.loadPinInfoByHeader(headerPin); err != nil {
		return err
	}
	//This Pin Checks Out...
	g.Pin = rpio.Pin(g.PinInfo.BcmPin)
	g.Pin.Output()
	if initHigh {
		g.Pin.High()
	} else {
		g.Pin.Low()
	}

	return nil
}

func (g *Gpio) loadPinInfoByHeader(headerPin int) error {
	pins := GetPins()
	for i := 0; i < len(pins); i++ {
		if pins[i].HeaderPin == headerPin {
			if strings.Contains(pins[i].Name, "GPIO") {
				g.Failed = false
				g.PinInfo = pins[i]
				return nil
			}
			g.Failed = true
			return errors.New("[GPIO PIN INIT]: Header Pin:" + strconv.Itoa(headerPin) + " is not a GPIO pin. Type: " + pins[i].Name)
		}
	}
	g.Failed = true
	return errors.New("[GPIO PIN INIT]: Header Pin:" + strconv.Itoa(headerPin) + "not found.")
}

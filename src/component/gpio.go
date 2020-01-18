package component

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/stianeikeland/go-rpio/v4"
)

//Gpio - a raspberrypi digital gpio pin
type Gpio struct {
	Pin     rpio.Pin
	PinInfo RpiPinMap
	Failed  bool
}

func (g *Gpio) String() string {
	var pinString string
	json, err := json.Marshal(g.Pin)
	if err != nil {
		pinString = err.Error()
	} else {
		pinString = string(json)
	}
	return labelStringLine("\tFAILED", strconv.FormatBool(g.Failed)) +
		labelStringLine("\tPin", pinString) +
		labelStringLine("\tPin Info", g.PinInfo.String())
}

//Init - create gpio pin object and set modes
func (g *Gpio) Init(headerPin int, initHigh bool) error {
	if err := g.loadPinInfoByHeader(headerPin); err != nil {
		return err
	}
	log.Println("BCM Pin:", g.PinInfo.BcmPin)
	//This Pin Checks Out...
	g.Pin = rpio.Pin(g.PinInfo.BcmPin)
	//g.Pin.Output()
	if initHigh {
		//g.Pin.High()
	} else {
		//g.Pin.Low()
	}
	log.Println("[GPIO INIT]: Init Completed!\n" + g.String())
	return nil
}

func (g *Gpio) HandleEnable() bool {

	return true
}

func (g *Gpio) HandleDisable() bool {
	g.Pin.Low()
	return true
}

/*CurrentStateString just for pretty printing the device info */
func (g *Gpio) CurrentStateString() string {
	state := "OFF"

	if g.Pin.Read() == rpio.High {
		state = "ON"
	}
	return g.PinInfo.String() + "\n\t" + "State: " + state
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

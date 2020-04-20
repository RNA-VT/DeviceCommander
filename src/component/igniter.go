package component

import (
	"firecontroller/io"
	"firecontroller/utilities"
	"log"
	"strconv"
	"time"
)

//Igniter - Thesr are hot... when we want them to be
type Igniter struct {
	GPIO          io.Gpio
	IgniterConfig `yaml:",inline"`
}

//IgniterConfig -
type IgniterConfig struct {
	Type          IgniterType `yaml:"type"`
	PinMap        io.RpiPinMap
	BaseComponent `yaml:",inline"`
}

//IgniterType -
type IgniterType string

const (
	//GlowFly -
	GlowFly IgniterType = "glowfly"
	//Induction -
	Induction = "induction"
)

//GetConfig - A transportable and marshalable version of this igniter
func (i Igniter) GetConfig() (config IgniterConfig) {

	config.UID = i.UID
	config.Enabled = i.Enabled
	config.Name = i.Name
	config.HeaderPin = i.HeaderPin
	config.Metadata = i.Metadata
	config.Type = i.Type
	config.PinMap = i.GPIO.PinInfo
	return
}

//Load -
func (i *Igniter) Load(config IgniterConfig) {
	i.UID = config.UID
	i.Enabled = config.Enabled
	i.Name = config.Name
	i.HeaderPin = config.HeaderPin
	i.Metadata = config.Metadata
	i.Type = config.Type
}

//Init - Enable, set initial value, log igniter initial state
func (i *Igniter) Init() error {
	err := i.Enable(true)
	if err != nil {
		return err
	}
	//Create UUID now that GPIO is initilized
	i.setID()
	log.Println("Enabled and Initialized Igniter:", i.String())

	return nil
}

//Enable - enable this igniter and optionally initalize its gpio
func (i *Igniter) Enable(init bool) (err error) {
	i.Enabled = true
	if init {
		err := i.GPIO.Init(i.HeaderPin, false)
		if err != nil {
			return err
		}
	}
	return nil
}

//Disable - disable this igniter and set state to 'off'
func (i *Igniter) Disable() {
	i.Enabled = false
}

func (i Igniter) String() string {
	metadata, err := utilities.StringJSON(i.Metadata)
	if err != nil {
		log.Println("failed to unmarshal metadata: ", string(metadata), err)
	}
	return "\n[Component]: Igniter" +
		utilities.LabelString("UID", strconv.Itoa(i.UID)) +
		utilities.LabelString("Name", i.Name) +
		utilities.LabelString("Header Pin", strconv.Itoa(i.HeaderPin)) +
		utilities.LabelString("Enabled", strconv.FormatBool(i.Enabled)) +
		utilities.LabelString("Gpio", i.GPIO.String()) +
		utilities.LabelString("Metadata", metadata)
}

//State returns a string represnting the current state
func (i Igniter) State() string {
	return "[GPIO PIN " + strconv.Itoa(i.HeaderPin) + "]: " + i.GPIO.CurrentStateString()
}

//Healthy - true if this component is healthy
func (i Igniter) Healthy() bool {
	return i.Enabled && !i.GPIO.Failed
}

func (i *Igniter) setID() {
	//HeaderPin is unique per micro, but this may need to be revisited for components requiring more than 1 HeaderPin
	i.UID = i.HeaderPin
}

//Command - process a command request for this solenoid
func (i *Igniter) Command(cmd string) {
	switch cmd {
	case "on":
		i.On()
	case "off":
		i.Off()
	case "enable":
		i.Enable(false)
	case "disable":
		i.Disable()
	}
}

//On - hawt
func (i *Igniter) On() {
	if i.Healthy() {
		i.GPIO.Pin.High()
	}
}

//Off - not hawt
func (i *Igniter) Off() {
	if i.Healthy() {
		i.GPIO.Pin.Low()
	}
}

//OnForDuration - light this igniter for a set period of time
func (i *Igniter) OnForDuration(duration int) {
	if i.Healthy() {
		switch i.Type {
		case Induction:
			fallthrough
		case GlowFly:
			i.On()
			i.offAfter(duration)
			break
		}
	} else {
		log.Println("Cannot light unhealthy Igniter")
	}
}
func (i *Igniter) offAfter(duration int) {
	if length, err := time.ParseDuration(strconv.Itoa(duration) + "ms"); err == nil {
		time.AfterFunc(length, i.Off)
	} else {
		//Log Failure to Close
		log.Println(
			"Failed to arrest an igniter:",
			"Invalid or malformed delay time.",
			" ~~~~Arresting now~~~~",
		)
		i.Off()
	}
}

//Edit - change things about the igniter.
func (i *Igniter) Edit(newValues map[string]interface{}) bool {

	return true
}

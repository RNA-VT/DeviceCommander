package component

import (
	"encoding/json"
	"firecontroller/io"
	"firecontroller/utilities"
	"log"
	"strconv"
	"time"
)

// Solenoid - base component + solenoid specific metadata
type Solenoid struct {
	BaseComponent `yaml:",inline"`
	Type          SolenoidType `yaml:"type"`
	Mode          SolenoidMode `yaml:"mode"`
	GPIO          io.Gpio
}

//Init - Enable, set initial value, log solenoid initial state
func (s *Solenoid) Init() error {
	s.Enable(true)
	log.Println("Enabled and Initialized Solenoid:", s.String())
	//TODO: Look into what feedback we can get on gpio init
	return nil
}

//Enable and optionally initialize this Solenoid
func (s *Solenoid) Enable(init bool) {
	s.Enabled = true
	if init {
		s.GPIO.Init(s.HeaderPin, false)
	}
}

//Disable this solenoid
func (s *Solenoid) Disable() {
	s.Enabled = false
}

func (s *Solenoid) String() string {
	metadata, err := json.Marshal(s.Metadata)
	metaString := ""
	if err != nil {
		log.Println("failed to unmarshal metadata: ", string(metadata), err)
	} else {
		metaString = utilities.LabelString("Metadata", string(metadata))
	}
	return "\nSolenoid Device:" +
		utilities.LabelString("UID", strconv.Itoa(s.UID)) +
		utilities.LabelString("Name", s.Name) +
		utilities.LabelString("Header Pin", strconv.Itoa(s.HeaderPin)) +
		utilities.LabelString("Enabled", strconv.FormatBool(s.Enabled)) +
		utilities.LabelString("Type", string(s.Type)) +
		utilities.LabelString("Mode", string(s.Mode)) +
		utilities.LabelString("Gpio", s.GPIO.String()) + metaString

}

//State returns a string of the current state of this solenoid
func (s *Solenoid) State() string {
	return "[GPIO PIN " + strconv.Itoa(s.HeaderPin)
}

func (s *Solenoid) open(duration int) {
	if s.healthy() {
		s.GPIO.Pin.High()
		s.close(duration)
	} else {
		//Log attempt to open unhealthy solenoid
	}
}

func (s *Solenoid) close(delay int) {
	if s.healthy() {
		if duration, err := time.ParseDuration(strconv.Itoa(delay) + "ms"); err == nil {
			time.AfterFunc(duration, s.GPIO.Pin.Low)
		} else {
			//Log Failure to Close
		}
	} else {
		//Log attempt to close unhealthy
	}
}

func (s *Solenoid) healthy() bool {
	return s.Enabled && !s.GPIO.Failed
}

// SolenoidType -
type SolenoidType string

const (
	// NormallyClosed represents a solenoid that does not allow flow without power
	NormallyClosed SolenoidType = "NC"
	// NormallyOpen represents a solenoid that is allows flow without power
	NormallyOpen = "NO"
)

//SolenoidMode -
type SolenoidMode string

const (
	//Supply - tank supply, pilot supply and transport solenoids
	Supply SolenoidMode = "supply"
	//Outlet - propane exhaust solenoid
	Outlet = "outlet"
)

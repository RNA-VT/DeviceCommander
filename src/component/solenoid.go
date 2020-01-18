package component

import (
	"log"
	"strconv"
	"time"
)

// Solenoid - base component + solenoid specific metadata
type Solenoid struct {
	BaseComponent `yaml:",inline"`
	Voltage       int           `yaml:"voltage"`
	Type          SolenoidTypes `yaml:"type"`
	Mode          SolenoidModes `yaml:"mode"`
	GPIO          Gpio
}

//Init - Enable, set initial value, log solenoid initial state
func (s *Solenoid) Init() error {
	log.Println("Enabled and Initialized Solenoid:", s.String())
	//TODO: Look into what feedback we can get on gpio init
	return nil
}

//Enable and optionally initialize this Solenoid
func (s *Solenoid) Enable(init bool) {
	s.GetBase().Enabled = true
	if init {
		s.GPIO.Init(s.GetBase().HeaderPin, false)
	}
}

//Disable this solenoid
func (s *Solenoid) Disable() {
	s.GetBase().Enabled = false
}

func (s *Solenoid) String() string {
	return "\nSolenoid Device:" +
		labelStringLine("UID", strconv.Itoa(s.UID)) +
		labelStringLine("Name", s.Name) +
		labelStringLine("Header Pin", strconv.Itoa(s.HeaderPin)) +
		labelStringLine("Enabled", strconv.FormatBool(s.Enabled)) +
		labelStringLine("Type", string(s.Type)) +
		labelStringLine("Mode", string(s.Mode)) +
		labelStringLine("Voltage", strconv.Itoa(s.Voltage)) +
		labelStringLine("Gpio", s.GPIO.String())
}

func labelStringLine(key string, value string) string {
	return "\n\t[" + key + "]:     " + value
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

// SolenoidTypes -
type SolenoidTypes string

const (
	// NormallyClosed represents a solenoid that does not allow flow without power
	NormallyClosed = "NC"
	// NormallyOpen represents a solenoid that is allows flow without power
	NormallyOpen = "NO"
)

//SolenoidModes -
type SolenoidModes string

const (
	//Supply - tank supply, pilot supply and transport solenoids
	Supply = "supply"
	//Igniter - glowfly or other HSI
	Igniter = "igniter"
	//Outlet - propane exhaust solenoid
	Outlet = "outlet"
)

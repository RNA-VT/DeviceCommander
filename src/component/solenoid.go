package component

import (
	"strconv"
	"time"
)

// Solenoid - base component + solenoid specific metadata
type Solenoid struct {
	BaseComponent
	Voltage int           `yaml:"voltage"`
	Type    SolenoidTypes `yaml:"type"`
	Mode    SolenoidModes `yaml:"mode"`
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
type SolenoidTypes int

const (
	// NormallyClosed represents a solenoid that does not allow flow without power
	NormallyClosed = "NC"
	// NormallyOpen represents a solenoid that is allows flow without power
	NormallyOpen = "NO"
)

//SolenoidModes -
type SolenoidModes string

const (
	//Source - tank supply, pilot supply and transport solenoids
	Source = "source"
	//Igniter - glowfly or other HSI
	Igniter = "igniter"
	//Outlet - propane exhaust solenoid
	Outlet = "outlet"
)

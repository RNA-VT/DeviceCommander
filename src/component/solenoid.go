package component

// Solenoid - base component + solenoid specific metadata
type Solenoid struct {
	BaseComponent
	Voltage int
	Type    SolenoidType
}

func (s *Solenoid) open(duration int) {

}

func (s *Solenoid) close(delay int) {

}

// SolenoidType -
type SolenoidType int

const (
	// NormallyClosed represents a solenoid that does not allow flow without power
	NormallyClosed = iota + 1
	// NormallyOpen represents a solenoid that is allows flow without power
	NormallyOpen
)

package mock

import (
	"github.com/stianeikeland/go-rpio/v4"
)

//PinState -
type PinState struct {
	RpioMode  rpio.Mode
	RpioState rpio.State
}

var localState PinState

//Pin mocks an rpio.Pin object
type Pin struct {
	Pin uint8
}

// Output - Set pin as Output
func (pin Pin) Output() {
	localState.RpioMode = rpio.Output
}

// High - Set pin High
func (pin Pin) High() {
	WritePin(pin, rpio.High)
}

// Low - Set pin Low
func (pin Pin) Low() {
	WritePin(pin, rpio.Low)
}

// Toggle pin state
func (pin Pin) Toggle() {
	TogglePin(pin)
}

// Mode - Set pin Mode
func (pin Pin) Mode(mode rpio.Mode) {
	PinMode(pin, mode)
}

// Set pin state (high/low)
func (pin Pin) Write(state rpio.State) {
	WritePin(pin, state)
}

// Read pin state (high/low)
func (pin Pin) Read() rpio.State {
	return ReadPin(pin)
}

// PinMode sets the mode of a given pin (Input, Output, Clock, Pwm or Spi)
// Only Output is properly mocked right now
func PinMode(pin Pin, mode rpio.Mode) {
	localState.RpioMode = mode
}

// WritePin sets a given pin High or Low
// by setting the clear or set registers respectively
func WritePin(pin Pin, state rpio.State) {
	localState.RpioState = state
}

// ReadPin - Read the state of a pin
func ReadPin(pin Pin) rpio.State {
	return localState.RpioState
}

// TogglePin - Toggle a pin state (high -> low -> high)
func TogglePin(pin Pin) {
	if localState.RpioState == rpio.High {
		pin.Low()
	} else {
		pin.High()
	}
}

/*

From the RPIO Package:

Package rpio provides GPIO access on the Raspberry PI without any need
for external c libraries (eg. WiringPi or BCM2835).

Supports simple operations such as:
	- Pin mode/direction (input/output/clock/pwm)
	- Pin write (high/low)
	- Pin read (high/low)
	- Pin edge detection (no/rise/fall/any)
	- Pull up/down/off
Also clock/pwm related oparations:
	- Set Clock frequency
	- Set Duty cycle
And SPI oparations:
	- SPI transmit/recieve/exchange bytes
	- Chip select
	- Set speed

Example of use:

	rpio.Open()
	defer rpio.Close()

	pin := rpio.Pin(4)
	pin.Output()

	for {
		pin.Toggle()
		time.Sleep(time.Second)
	}

The library use the raw BCM2835 pinouts, not the ports as they are mapped
on the output pins for the raspberry pi, and not the wiringPi convention.

            Rev 2 and 3 Raspberry Pi                        Rev 1 Raspberry Pi (legacy)
  +-----+---------+----------+---------+-----+      +-----+--------+----------+--------+-----+
  | BCM |   Name  | Physical | Name    | BCM |      | BCM | Name   | Physical | Name   | BCM |
  +-----+---------+----++----+---------+-----+      +-----+--------+----++----+--------+-----+
  |     |    3.3v |  1 || 2  | 5v      |     |      |     | 3.3v   |  1 ||  2 | 5v     |     |
  |   2 |   SDA 1 |  3 || 4  | 5v      |     |      |   0 | SDA    |  3 ||  4 | 5v     |     |
  |   3 |   SCL 1 |  5 || 6  | 0v      |     |      |   1 | SCL    |  5 ||  6 | 0v     |     |
  |   4 | GPIO  7 |  7 || 8  | TxD     | 14  |      |   4 | GPIO 7 |  7 ||  8 | TxD    |  14 |
  |     |      0v |  9 || 10 | RxD     | 15  |      |     | 0v     |  9 || 10 | RxD    |  15 |
  |  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |      |  17 | GPIO 0 | 11 || 12 | GPIO 1 |  18 |
  |  27 | GPIO  2 | 13 || 14 | 0v      |     |      |  21 | GPIO 2 | 13 || 14 | 0v     |     |
  |  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |      |  22 | GPIO 3 | 15 || 16 | GPIO 4 |  23 |
  |     |    3.3v | 17 || 18 | GPIO  5 | 24  |      |     | 3.3v   | 17 || 18 | GPIO 5 |  24 |
  |  10 |    MOSI | 19 || 20 | 0v      |     |      |  10 | MOSI   | 19 || 20 | 0v     |     |
  |   9 |    MISO | 21 || 22 | GPIO  6 | 25  |      |   9 | MISO   | 21 || 22 | GPIO 6 |  25 |
  |  11 |    SCLK | 23 || 24 | CE0     | 8   |      |  11 | SCLK   | 23 || 24 | CE0    |   8 |
  |     |      0v | 25 || 26 | CE1     | 7   |      |     | 0v     | 25 || 26 | CE1    |   7 |
  |   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |      +-----+--------+----++----+--------+-----+
  |   5 | GPIO 21 | 29 || 30 | 0v      |     |
  |   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
  |  13 | GPIO 23 | 33 || 34 | 0v      |     |
  |  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
  |  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
  |     |      0v | 39 || 40 | GPIO 29 | 21  |
  +-----+---------+----++----+---------+-----+

See the spec for full details of the BCM2835 controller:

https://www.raspberrypi.org/documentation/hardware/raspberrypi/bcm2835/BCM2835-ARM-Peripherals.pdf
and https://elinux.org/BCM2835_datasheet_errata - for errors in that spec

*/

package io

import (
	"firecontroller/utilities"
	"strconv"
)

/*
  +-----+---------+----------+---------+-----+
  | BCM |   Name  | Physical | Name    | BCM |
  +-----+---------+----++----+---------+-----+
  |     |    3.3v |  1 || 2  | 5v      |     |
  |   2 |   SDA 1 |  3 || 4  | 5v      |     |
  |   3 |   SCL 1 |  5 || 6  | 0v      |     |
  |   4 | GPIO  7 |  7 || 8  | TxD     | 14  |
  |     |      0v |  9 || 10 | RxD     | 15  |
  |  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |
  |  27 | GPIO  2 | 13 || 14 | 0v      |     |
  |  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |
  |     |    3.3v | 17 || 18 | GPIO  5 | 24  |
  |  10 |    MOSI | 19 || 20 | 0v      |     |
  |   9 |    MISO | 21 || 22 | GPIO  6 | 25  |
  |  11 |    SCLK | 23 || 24 | CE0     | 8   |
  |     |      0v | 25 || 26 | CE1     | 7   |
  |   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |
  |   5 | GPIO 21 | 29 || 30 | 0v      |     |
  |   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
  |  13 | GPIO 23 | 33 || 34 | 0v      |     |
  |  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
  |  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
  |     |      0v | 39 || 40 | GPIO 29 | 21  |
  +-----+---------+----++----+---------+-----+
*/

//RpiPinMap -
type RpiPinMap struct {
	//BcmPin - The pin id on the processor, used by rpio for commands
	BcmPin uint8
	//Human readable name for the Raspi Pin
	Name string
	//Connection position on the header according to the block above^^
	HeaderPin int
}

func (r RpiPinMap) String() string {
	return utilities.LabelString("\t\tName", r.Name) +
		utilities.LabelString("\t\tHeader Pin", strconv.Itoa(r.HeaderPin)) +
		utilities.LabelString("\t\tBCM Pin", strconv.Itoa(int(r.BcmPin)))
}

//GetPins - Returns Pins for Raspi 4
func GetPins() []RpiPinMap {
	var noPin uint8 = 255
	return []RpiPinMap{
		RpiPinMap{
			HeaderPin: 1,
			BcmPin:    noPin,
			Name:      "3.3v",
		},
		RpiPinMap{
			HeaderPin: 3,
			BcmPin:    2,
			Name:      "SDA 1",
		},
		RpiPinMap{
			HeaderPin: 5,
			BcmPin:    3,
			Name:      "SCL 1",
		},
		RpiPinMap{
			HeaderPin: 7,
			BcmPin:    4,
			Name:      "GPIO 7",
		},
		RpiPinMap{
			HeaderPin: 9,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 11,
			BcmPin:    17,
			Name:      "GPIO 0",
		},
		RpiPinMap{
			HeaderPin: 13,
			BcmPin:    27,
			Name:      "GPIO 2",
		},
		RpiPinMap{
			HeaderPin: 15,
			BcmPin:    22,
			Name:      "GPIO 3",
		},
		RpiPinMap{
			HeaderPin: 17,
			BcmPin:    noPin,
			Name:      "3.3v",
		},
		RpiPinMap{
			HeaderPin: 19,
			BcmPin:    10,
			Name:      "MOSI",
		},
		RpiPinMap{
			HeaderPin: 21,
			BcmPin:    9,
			Name:      "MISO",
		},
		RpiPinMap{
			HeaderPin: 23,
			BcmPin:    11,
			Name:      "SCLK",
		},
		RpiPinMap{
			HeaderPin: 25,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 27,
			BcmPin:    0,
			Name:      "SDA 0",
		},
		RpiPinMap{
			HeaderPin: 29,
			BcmPin:    5,
			Name:      "GPIO 21",
		},
		RpiPinMap{
			HeaderPin: 31,
			BcmPin:    6,
			Name:      "GPIO 22",
		},
		RpiPinMap{
			HeaderPin: 33,
			BcmPin:    13,
			Name:      "GPIO 23",
		},
		RpiPinMap{
			HeaderPin: 35,
			BcmPin:    19,
			Name:      "GPIO 24",
		},
		RpiPinMap{
			HeaderPin: 37,
			BcmPin:    26,
			Name:      "GPIO 25",
		},
		RpiPinMap{
			HeaderPin: 39,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 2,
			BcmPin:    noPin,
			Name:      "5v",
		},
		RpiPinMap{
			HeaderPin: 4,
			BcmPin:    noPin,
			Name:      "5v",
		},
		RpiPinMap{
			HeaderPin: 6,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 8,
			BcmPin:    14,
			Name:      "TxD",
		},
		RpiPinMap{
			HeaderPin: 10,
			BcmPin:    15,
			Name:      "RxD",
		},
		RpiPinMap{
			HeaderPin: 12,
			BcmPin:    18,
			Name:      "GPIO 1",
		},
		RpiPinMap{
			HeaderPin: 14,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 16,
			BcmPin:    23,
			Name:      "GPIO 4",
		},
		RpiPinMap{
			HeaderPin: 18,
			BcmPin:    24,
			Name:      "GPIO 5",
		},
		RpiPinMap{
			HeaderPin: 20,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 22,
			BcmPin:    25,
			Name:      "GPIO 6",
		},
		RpiPinMap{
			HeaderPin: 24,
			BcmPin:    8,
			Name:      "CE0",
		},
		RpiPinMap{
			HeaderPin: 26,
			BcmPin:    7,
			Name:      "CE1",
		},
		RpiPinMap{
			HeaderPin: 28,
			BcmPin:    1,
			Name:      "SCL 0",
		},
		RpiPinMap{
			HeaderPin: 30,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 32,
			BcmPin:    12,
			Name:      "GPIO 26",
		},
		RpiPinMap{
			HeaderPin: 34,
			BcmPin:    noPin,
			Name:      "0v",
		},
		RpiPinMap{
			HeaderPin: 36,
			BcmPin:    16,
			Name:      "GPIO 27",
		},
		RpiPinMap{
			HeaderPin: 38,
			BcmPin:    20,
			Name:      "GPIO 28",
		},
		RpiPinMap{
			HeaderPin: 40,
			BcmPin:    21,
			Name:      "GPIO 29",
		},
	}
}

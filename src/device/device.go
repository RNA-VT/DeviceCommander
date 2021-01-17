package device

import "strconv"

//Device represents a compliant physical component & its web address.
type Device struct {
	//ID is the serial nummber of the connecting device
	ID string `json:"id"`
	//Name - Optional Device Nickname
	Name string `json:"name"`
	//Description - Optional text describing this device
	Description string `json:"description"`
	//Host - Device Api Host
	Host string `json:"host"`
	//Port - Device Api Port. Set to 443 for https
	Port     int `json:"port"`
	failures int
}

//NewDevice -
func NewDevice(host string, port int) (Device, error) {
	dev := Device{
		Host:     host,
		Port:     port,
		failures: 0,
	}
	//TODO: Load Device Data from db
	return dev, nil
}

//URL returns a network address including the ip address and port that this micro is listening on
func (d Device) URL() string {
	return d.protocol() + "://" + d.Host + ":" + strconv.Itoa(d.Port)
}

func (d Device) protocol() string {
	var protocol string
	if d.Port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

//ProcessHealthCheckResult - updates health check failure count & returns
func (d *Device) ProcessHealthCheckResult(result bool) {
	if result { //Healthy
		d.failures = 0
	} else {
		d.failures++
	}
}

//Failed - If true, device should be deregistered
func (d Device) Failed() bool {
	failThreshold := 3
	return d.failures >= failThreshold
}

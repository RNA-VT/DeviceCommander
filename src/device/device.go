package device

import (
	"encoding/json"
	"fmt"
	"io"
)

// Device represents a compliant physical component & its web address.
type Device struct {
	// ID is the serial nummber of the connecting device
	ID string `json:"id"`
	// Name - Optional Device Nickname
	Name string `json:"name"`
	// Description - Optional text describing this device
	Description string `json:"description"`
	// Host - Device Api Host
	Host string `json:"host"`
	// Port - Device Api Port. Set to 443 for https
	Port     int `json:"port"`
	failures int
}

// NewDevice -
func NewDevice(host string, port int) (Device, error) {
	dev := Device{
		Host:     host,
		Port:     port,
		failures: 0,
	}

	return dev, nil
}

func NewDeviceFromRequestBody(body io.ReadCloser) (Device, error) {
	deviceLogger := getDeviceLogger()

	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev Device
	err := decoder.Decode(&dev)
	if err != nil {
		deviceLogger.Error("Failed to decode device config from request body", err)
		return dev, err
	}

	return dev, nil
}

// URL returns a network address including the ip address and port that this device is listening on
func (d Device) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Host, d.Port)
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

// ProcessHealthCheckResult - updates health check failure count & returns
func (d *Device) ProcessHealthCheckResult(result bool) {
	if result { // Healthy
		d.failures = 0
	} else {
		d.failures++
	}
}

// Failed - If true, device should be deregistered
func (d Device) Unresponsive() bool {
	failThreshold := 3
	return d.failures >= failThreshold
}

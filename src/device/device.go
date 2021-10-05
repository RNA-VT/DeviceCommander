package device

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/rna-vt/devicecommander/graph/model"
)

// // Device represents a compliant physical component & its web address.
// type Device struct {
// 	// ID is the serial nummber of the connecting device
// 	ID string `json:"id"`
// 	// Name - Optional Device Nickname
// 	Name string `json:"name"`
// 	// Description - Optional text describing this device
// 	Description string `json:"description"`
// 	// Host - Device Api Host
// 	Host string `json:"host"`
// 	// Port - Device Api Port. Set to 443 for https
// 	Port     int `json:"port"`
// 	failures int
// }

type DeviceObj struct {
	device *model.Device
}

// NewDevice -
func NewDeviceObj(d *model.Device) (DeviceObj, error) {
	dev := DeviceObj{
		device: d,
	}

	return dev, nil
}

// NewDevice -
func NewDevice(host string, port int) (model.Device, error) {
	dev := model.Device{
		Host:     host,
		Port:     port,
		Failures: 0,
	}

	return dev, nil
}

func NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error) {
	deviceLogger := getDeviceLogger()

	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev model.NewDevice
	err := decoder.Decode(&dev)
	if err != nil {
		deviceLogger.Error("Failed to decode device config from request body", err)
		return dev, err
	}

	return dev, nil
}

// URL returns a network address including the ip address and port that this device is listening on
func (d DeviceObj) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.device.Host, d.device.Port)
}

func (d DeviceObj) protocol() string {
	var protocol string
	if d.device.Port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

// ProcessHealthCheckResult - updates health check failure count & returns
func (d *DeviceObj) ProcessHealthCheckResult(result bool) {
	if result { // Healthy
		d.device.Failures = 0
	} else {
		d.device.Failures++
	}
}

// Failed - If true, device should be deregistered
func (d DeviceObj) Unresponsive() bool {
	failThreshold := 3
	return d.device.Failures >= failThreshold
}

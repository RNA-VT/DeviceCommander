package device

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rna-vt/devicecommander/graph/model"
)

type Interface interface {
	NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error)
	URL() string
	protocol() string
	ProcessHealthCheckResult(result bool)
	Unresponsive() bool
	CheckHealth() (*Device, error)
	EvaluateHealthCheckResponse(resp *http.Response) bool
}

// Device is a wrapper for the Device model. It aims to provide a helpful
// layer of abstraction away from the gqlgen/postgres models.
type Device struct {
	model.Device
}

// NewDeviceWrapper creates a new instance of a device.Wrapper
func NewDeviceWrapper(d *model.Device) (*Device, error) {
	dev := Device{
		// Device: d,
	}

	return &dev, nil
}

// NewDevice creates a barebones new instance of a Device with a host and port.
func NewDevice(host string, port int) (model.Device, error) {
	dev := model.Device{
		Host:     host,
		Port:     port,
		Failures: 0,
	}

	return dev, nil
}

// NewDeviceFromRequestBody creates a new instance of a NewDevice
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
func (d Device) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Host, d.Port)
}

// protocol determines the http/https protocol by Port allocation
func (d Device) protocol() string {
	var protocol string
	if d.Device.Port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

// ProcessHealthCheckResult - updates health check failure count & returns the failure count
func (d Device) ProcessHealthCheckResult(result bool) int {
	if result { // Healthy
		return 0
	}
	return d.Failures + 1
}

// Failed - If true, device should be deregistered
func (d Device) Unresponsive() bool {
	failThreshold := 3
	return d.Failures >= failThreshold
}

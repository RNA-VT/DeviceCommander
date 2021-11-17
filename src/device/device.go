package device

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/graph/model"
)

type IDevice interface {
	NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error)
	URL() string
	protocol() string
	ProcessHealthCheckResult(result bool)
	Unresponsive() bool
	RunHealthCheck(client IDeviceClient) (Device, error)
}

// Device is a wrapper for the Device model. It aims to provide a helpful
// layer of abstraction away from the gqlgen/postgres models.
type Device struct {
	Device *model.Device
	logger *log.Entry
}

// DeviceFromNewDevice generates a Device from a NewDevice with the correct instantiations.
// This should be the primary method for creationg model.Device(s).
func FromNewDevice(newDeviceArgs model.NewDevice) model.Device {
	newDevice := model.Device{
		ID:        uuid.New(),
		Host:      newDeviceArgs.Host,
		Port:      newDeviceArgs.Port,
		Endpoints: []model.Endpoint{},
	}

	if newDeviceArgs.Mac != nil {
		newDevice.MAC = *newDeviceArgs.Mac
	}

	if newDeviceArgs.Name != nil {
		newDevice.Name = *newDeviceArgs.Name
	}

	if newDeviceArgs.Description != nil {
		newDevice.Description = *newDeviceArgs.Description
	}

	return newDevice
}

// NewDeviceWrapper creates a new instance of a device.Wrapper.
func NewDeviceWrapper(d model.Device) Device {
	dev := Device{
		Device: &d,
		logger: log.WithFields(log.Fields{"module": "device"}),
	}

	return dev
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

// NewDeviceFromRequestBody creates a new instance of a NewDevice.
func NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev model.NewDevice
	err := decoder.Decode(&dev)
	if err != nil {
		return dev, err
	}

	return dev, nil
}

// URL returns a network address including the ip address and port that this device is listening on.
func (d Device) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Device.Host, d.Device.Port)
}

// protocol determines the http/https protocol by Port allocation.
func (d Device) protocol() string {
	var protocol string
	if d.Device.Port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

// ProcessHealthCheckResult updates health check failure count & returns the failure count.
func (d Device) ProcessHealthCheckResult(result bool) int {
	if result { // Healthy
		return 0
	}
	return d.Device.Failures + 1
}

// Unresponsive determines the state of the device relative to its past performance.
// If true, device should be deregistered.
func (d Device) Unresponsive() bool {
	failThreshold := 3
	return d.Device.Failures >= failThreshold
}

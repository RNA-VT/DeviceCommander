package device

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/graph/model"
)

type Device interface {
	NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error)
	ID() uuid.UUID
	URL() string
	protocol() string
	ProcessHealthCheckResult(result bool) int
	Unresponsive() bool
	RunHealthCheck(client Client) error
}

// DeviceFromNewDevice generates a Device from a NewDevice with the correct instantiations.
// This should be the primary method for creating model.Device(s).
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

	return newDevice
}

// Device is a wrapper for the Device model. It aims to provide a helpful
// layer of abstraction away from the gqlgen/postgres models.
type BasicDevice struct {
	Device *model.Device
	logger *log.Entry
}

// NewDeviceWrapper creates a new instance of a device.Wrapper.
func NewDeviceWrapper(d model.Device) BasicDevice {
	dev := BasicDevice{
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
func (d BasicDevice) NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev model.NewDevice

	if err := decoder.Decode(&dev); err != nil {
		return dev, err
	}

	return dev, nil
}

func (d BasicDevice) ID() uuid.UUID {
	return d.Device.ID
}

// URL returns a network address including the ip address and port that this device is listening on.
func (d BasicDevice) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Device.Host, d.Device.Port)
}

// protocol determines the http/https protocol by Port allocation.
func (d BasicDevice) protocol() string {
	var protocol string
	if d.Device.Port == viper.GetInt("DEFAULT_TLS_PORT") {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

// ProcessHealthCheckResult updates health check failure count & returns the failure count.
func (d BasicDevice) ProcessHealthCheckResult(result bool) int {
	if result { // Healthy
		return 0
	}
	return d.Device.Failures + 1
}

// Unresponsive determines the state of the device relative to its past performance.
// If true, device should be deregistered.
func (d BasicDevice) Unresponsive() bool {
	failThreshold := 3
	return d.Device.Failures >= failThreshold
}

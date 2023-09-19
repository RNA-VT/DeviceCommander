package device

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/device/endpoint"
)

type BasicDevice interface {
	URL() string
	protocol() string
	ProcessHealthCheckResult(result bool) int
	Unresponsive() bool
	RunHealthCheck(client Client) error
}

// NewDeviceFromRequestBody creates a new instance of a NewDevice.
func NewDeviceFromRequestBody(body io.ReadCloser) (NewDeviceParams, error) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev NewDeviceParams

	if err := decoder.Decode(&dev); err != nil {
		return dev, err
	}

	return dev, nil
}

// DeviceFromNewDevice generates a Device from a NewDevice with the correct instantiations.
// This should be the primary method for creating model.Device(s).
func FromNewDevice(newDeviceArgs NewDeviceParams) Device {
	newDevice := Device{
		ID:        uuid.New(),
		Host:      newDeviceArgs.Host,
		Port:      newDeviceArgs.Port,
		Endpoints: []endpoint.Endpoint{},
	}

	if newDeviceArgs.MAC != nil {
		newDevice.MAC = *newDeviceArgs.MAC
	}

	return newDevice
}

// Device is one of the core concepts of this application. A Device represents
// a microcontroller that complies with the DeviceCommander standard.
//
// swagger:model
type Device struct {

	// the UUID for the device.
	//
	// required: true
	// example: 705e4dcb-3ecd-24f3-3a35-3e926e4bded5
	ID uuid.UUID `json:"ID" faker:"uuid_hyphenated"`

	// the MAC address for this device.
	// required: true
	MAC string `json:"MAC" gorm:"unique" faker:"mac_address"`

	// the human readable name of the device.
	// required: false
	Name string `json:"Name"`

	// the description of the device.
	// required: false
	Description string `json:"Description"`

	// the host address of the device.
	// required: true
	Host string `json:"Host" faker:"ipv4"`

	// the active port of the device.
	// required: true
	Port int `json:"Port" faker:"boundary_start=49152, boundary_end=65535"`

	// the count of failed actions by the device.
	// required: false
	Failures int `json:"Failures" faker:"boundary_start=0, boundary_end=5"`

	// a flag representing the responsiveness of the device.
	// required: false
	Active bool `json:"Active"`

	// a list of endpoints available for quering on a device.
	// required: false
	Endpoints []endpoint.Endpoint `json:"Endpoints" faker:"-"`

	logger *log.Entry
}

// NewDevice creates a barebones new instance of a Device with a host and port.
func NewDevice(host string, port int) (Device, error) {
	dev := Device{
		Host:     host,
		Port:     port,
		Failures: 0,
	}

	return dev, nil
}

// URL returns a network address including the ip address and port that this device is listening on.
func (d Device) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Host, d.Port)
}

// protocol determines the http/https protocol by Port allocation.
func (d Device) protocol() string {
	var protocol string
	if d.Port == viper.GetInt("DEFAULT_TLS_PORT") {
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
	return d.Failures + 1
}

// Unresponsive determines the state of the device relative to its past performance.
// If true, device should be deregistered.
func (d Device) Unresponsive() bool {
	failThreshold := 3
	return d.Failures >= failThreshold
}

package device

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/rna-vt/devicecommander/graph/model"
)

type Interface interface {
	NewDeviceFromRequestBody(body io.ReadCloser) (model.NewDevice, error)
	URL() string
	protocol() string
	ProcessHealthCheckResult(result bool)
	Unresponsive() bool
	CheckHealth() (*Wrapper, error)
	EvaluateHealthCheckResponse(resp *http.Response) bool
}

// DeviceObj is a wrapper for the Device struct. It aims to provide a helpful
// layer of abstraction away from the gqlgen/postgres models.
type Wrapper struct {
	Device *model.Device
}

// NewDeviceWrapper creates a new instance of a device.Wrapper
func NewDeviceWrapper(d *model.Device) (*Wrapper, error) {
	dev := Wrapper{
		Device: d,
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
func (d Wrapper) URL() string {
	return fmt.Sprintf("%s://%s:%d", d.protocol(), d.Device.Host, d.Device.Port)
}

// protocol determines the http/https protocol by Port allocation
func (d Wrapper) protocol() string {
	var protocol string
	if d.Device.Port == 443 {
		protocol = "https"
	} else {
		protocol = "http"
	}
	return protocol
}

// ProcessHealthCheckResult - updates health check failure count & returns the failure count
func (d Wrapper) ProcessHealthCheckResult(result bool) int {
	if result { // Healthy
		return 0
	}
	return d.Device.Failures + 1
}

// Failed - If true, device should be deregistered
func (d Wrapper) Unresponsive() bool {
	failThreshold := 3
	return d.Device.Failures >= failThreshold
}

func (d Wrapper) GetValueHash(keys []string) string {
	r := reflect.ValueOf(d.Device)
	s := "hash:"
	for _, key := range keys {
		tmp := reflect.Indirect(r).FieldByName(key)
		log.Println(tmp.String())
		// if !tmp.IsNil() {
		// 	s = s + tmp.String()
		// }
	}

	log.Println(s)

	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	return string(bs)
}

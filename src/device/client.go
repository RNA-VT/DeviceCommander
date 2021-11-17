package device

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/graph/model"
)

// IDeviceClient implements the common http actions when interacting with a device
type IDeviceClient interface {
	Info(Device) (model.NewDevice, error)
	Health(Device) (*http.Response, error)
	EvaluateHealthCheckResponse(resp *http.Response, d Device) bool
}

// HTTPDeviceClient is an implementation of the IDeviceClient. It communicates
// with the device via http.
type HTTPDeviceClient struct {
	logger *log.Entry
}

// NewHTTPDeviceClient creates an instantiated HTTPDeviceClientl. This should be the
// primary method of generating a HTTPDeviceClient struct.
func NewHTTPDeviceClient() HTTPDeviceClient {
	return HTTPDeviceClient{
		log.WithFields(log.Fields{"module": "device-client"}),
	}
}

func (c HTTPDeviceClient) Info(d Device) (model.NewDevice, error) {
	panic("function not implemented")
}

// Health on the HTTPDeviceClient queries the device.URL `/health` endpoint to determine health.
func (c HTTPDeviceClient) Health(d Device) (*http.Response, error) {
	url := d.URL() + "/health"

	resp, err := http.Get(url)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("Error checking [%s] %s", url, err.Error()))
		return &http.Response{}, err
	}

	return resp, nil
}

// EvaluateHealthCheckResponse determines the final outcome of the HealthCheck by examining a http.Response.
// The only requirements for a positive health response is a status code of 200 and a parseable body.
func (c HTTPDeviceClient) EvaluateHealthCheckResponse(resp *http.Response, d Device) bool {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("failed to read the healthcheck response")
		return false
	}
	healthy := false
	switch resp.StatusCode {
	case 200:
		c.logger.WithFields(log.Fields{"event": "isHealthy"}).Info(d.Device.ID)
		healthy = true
	case 404:
		c.logger.Error("Registered Device Not Found: " + d.Device.ID.String())
	default:
		c.logger.Error("Unexpected Result: " + d.Device.ID.String())
		c.logger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
		c.logger.Error("Response: " + string(body))
	}
	return healthy
}
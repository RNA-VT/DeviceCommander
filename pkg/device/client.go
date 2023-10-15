package device

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/carlmjohnson/requests"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/pkg/utils"
)

// DeviceClient implements the common http actions when interacting with a device.
type Client interface {
	Health(Device) (*http.Response, error)
	EvaluateHealthCheckResponse(resp *http.Response, d Device) bool
	Specification(Device) (Specification, error)
}

// HTTPDeviceClient is an implementation of the IDeviceClient. It communicates
// with the device via http.
type HTTPDeviceClient struct {
	logger *log.Entry
}

// NewHTTPDeviceClient creates an instantiated HTTPDeviceClient. This should be the
// primary method of generating a HTTPDeviceClient struct.
func NewHTTPDeviceClient() HTTPDeviceClient {
	return HTTPDeviceClient{
		log.WithFields(log.Fields{"module": "device-client"}),
	}
}

func (c HTTPDeviceClient) dangerousHTTPGet(url string) (resp *http.Response, err error) {
	// look into validating this URL and the request/response.

	//nolint:gosec
	return http.Get(url)
}

// Health on the HTTPDeviceClient queries the device.URL `/health` endpoint to determine health.
func (c HTTPDeviceClient) Health(d Device) (*http.Response, error) {
	url := d.URL() + "/health"

	resp, err := c.dangerousHTTPGet(url)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error checking health [%s] %w", url, err)
	}

	return resp, nil
}

type HealthResponse struct{}

// EvaluateHealthCheckResponse determines the final outcome of the HealthCheck by examining a http.Response.
// The only requirements for a positive health response is a status code of 200 and a parseable body.
func (c HTTPDeviceClient) EvaluateHealthCheckResponse(resp *http.Response, d Device) bool {
	defer resp.Body.Close()

	healthResponse := HealthResponse{}
	err := json.NewDecoder(resp.Body).Decode(&healthResponse)
	if err != nil {
		c.logger.Error("failed to read the healthcheck response")
		return false
	}
	healthy := false

	switch resp.StatusCode {
	case http.StatusOK:
		c.logger.WithFields(log.Fields{"event": "isHealthy"}).Info(d.ID)
		healthy = true
	case http.StatusNotFound:
		c.logger.Error("Registered Device Not Found: " + d.ID.String())
	default:
		c.logger.Error("Unexpected Result: " + d.ID.String())
		c.logger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
		c.logger.Error("Response: " + utils.PrettyPrintJSON(healthResponse))
	}
	return healthy
}

func (c HTTPDeviceClient) Specification(d Device) (Specification, error) {
	var spec Specification
	err := requests.
		URL(d.URL()).
		Path("/specification").
		ToJSON(&spec).
		Fetch(context.Background())
	if err != nil {
		return spec, err
	}

	return spec, nil
}

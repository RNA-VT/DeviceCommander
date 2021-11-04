package device

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/graph/model"
)

type IDeviceClient interface {
	Info(Device) (model.NewDevice, error)
	Health(Device) (*http.Response, error)
	EvaluateHealthCheckResponse(resp *http.Response, d Device) bool
}

type HTTPDeviceClient struct {
	logger *log.Entry
}

func NewHTTPDeviceClient() HTTPDeviceClient {
	return HTTPDeviceClient{
		log.WithFields(log.Fields{"module": "device-client"}),
	}
}

func (c HTTPDeviceClient) Health(d Device) (*http.Response, error) {
	url := d.URL() + "/health"

	resp, err := http.Get(url)
	if err != nil {
		c.logger.Warn(fmt.Sprintf("Error checking [%s] %s", url, err.Error()))
		return &http.Response{}, err
	}

	return resp, nil
}

func (c HTTPDeviceClient) Info(d Device) (model.NewDevice, error) {
	panic("function not implemented")
}

func (c HTTPDeviceClient) EvaluateHealthCheckResponse(resp *http.Response, d Device) bool {
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	c.logger.Error("failed to read the healthcheck response")
	// 	return false
	// }
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
		// c.logger.Error("Response: " + string(body))
	}
	return healthy
}

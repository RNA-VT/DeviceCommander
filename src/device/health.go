package device

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// CheckHealth probes the health endpoint of the device in question. The health
// endpoint is currently at Device.URL()/health
func (d DeviceObj) CheckHealth() (DeviceObj, error) {
	logger := getDeviceLogger()

	url := d.URL() + "/health"

	logger.Info(fmt.Sprintf("Checking Device: [%s] at %s", d.device.ID, url))

	resp, err := http.Get(url)
	if err != nil {
		logger.Warn(fmt.Sprintf("Error checking [%s] %s", url, err.Error()))
		return DeviceObj{}, err
	}

	d.EvaluateHealthCheckResponse(resp)
	result := d.EvaluateHealthCheckResponse(resp)
	d.device.Failures = d.ProcessHealthCheckResult(result)

	// TODO: need to cleanup unresponsive nodes somewhere
	// if d.Unresponsive() {
	// 	healthDeregistrationLogger := healthLogger.WithFields(log.Fields{"event": "deregistration"})
	// 	healthDeregistrationLogger.Info("Failure Threshold Reached... Removing Device: " + d.ID)
	// 	c.RemoveDevice(d.ID)
	// }
}

// evaluateHealthCheckResponse inspects the repsponse from a device and extracts
// a few details. Firstly it will create useful logs for better understanding
// the response from the device health check. Secondly, it will return a true/false
// determining the health of the device.
func (d DeviceObj) EvaluateHealthCheckResponse(resp *http.Response) bool {
	logger := getDeviceLogger()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("reading the healtheck failed")
		return false
	}
	healthy := false
	switch resp.StatusCode {
	case 200:
		logger.WithFields(log.Fields{"event": "isHealthy"}).Info(d.device.ID)
		healthy = true
	case 404:
		logger.Error("Registered Device Not Found: " + d.device.ID.String())
	default:
		logger.Error("Unexpected Result: " + d.device.ID.String())
		logger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
		logger.Error("Response: " + string(body))
	}
	return healthy
}

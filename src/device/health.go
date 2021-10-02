package device

import (
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func (d Device) CheckHealth() {
	deviceLogger := getDeviceLogger()

	url := d.URL() + "/health"

	deviceLogger.Info("Checking Device:", d.ID, url)

	resp, err := http.Get(url)
	if err != nil {
		deviceLogger.Error(url)
		deviceLogger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
		deviceLogger.Error("Message: " + err.Error())
	}

	result := d.evaluateHealthCheckResponse(resp)
	d.ProcessHealthCheckResult(result)

	// TODO: need to cleanup unresponsive nodes somewhere

	// if d.Unresponsive() {
	// 	healthDeregistrationLogger := healthLogger.WithFields(log.Fields{"event": "deregistration"})
	// 	healthDeregistrationLogger.Info("Failure Threshold Reached... Removing Device: " + d.ID)
	// 	c.RemoveDevice(d.ID)
	// }
}

func (d Device) evaluateHealthCheckResponse(resp *http.Response) bool {
	deviceLogger := getDeviceLogger()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	healthy := false
	switch resp.StatusCode {
	case 200:
		deviceLogger.WithFields(log.Fields{"event": "isHealthy"}).Info(d.ID)
		healthy = true
	case 404:
		deviceLogger.Error("Registered Device Not Found: " + d.ID)
	default:
		deviceLogger.Error("Unexpected Result: " + d.ID)
		deviceLogger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
		deviceLogger.Error("Response: " + string(body))
	}
	return healthy
}

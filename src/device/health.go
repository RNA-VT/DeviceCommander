package device

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// CheckHealth probes the health endpoint of the device in question. The health
// endpoint is currently at Device.URL()/health
func (d Device) RunHealthCheck(client IDeviceClient) (Device, error) {
	logger := getDeviceLogger()

	// client := NewHttpDeviceClient()

	resp, err := client.Health(d)
	if err != nil {
		logger.Warn(fmt.Sprintf("Error checking [%s] %s", d.Device.ID.String(), err))
	}

	result := client.EvaluateHealthCheckResponse(resp, d)

	if result {
		log.Trace(fmt.Sprintf("device [%s] is healthy", d.Device.ID.String()))
	} else {
		log.Trace(fmt.Sprintf("device [%s] is not healthy", d.Device.ID.String()))
	}

	d.Device.Failures = d.ProcessHealthCheckResult(result)

	// TODO: need to cleanup unresponsive nodes somewhere
	// if d.Unresponsive() {
	// 	healthDeregistrationLogger := healthLogger.WithFields(log.Fields{"event": "deregistration"})
	// 	healthDeregistrationLogger.Info("Failure Threshold Reached... Removing Device: " + d.ID)
	// 	c.RemoveDevice(d.ID)
	// }
	return d, nil
}

// // evaluateHealthCheckResponse inspects the repsponse from a device and extracts
// // a few details. Firstly it will create useful logs for better understanding
// // the response from the device health check. Secondly, it will return a true/false
// // determining the health of the device.
// func (d Device) EvaluateHealthCheckResponse(resp *http.Response) bool {
// 	logger := getDeviceLogger()
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		logger.Error("reading the healtheck failed")
// 		return false
// 	}
// 	healthy := false
// 	switch resp.StatusCode {
// 	case 200:
// 		logger.WithFields(log.Fields{"event": "isHealthy"}).Info(d.Device.ID)
// 		healthy = true
// 	case 404:
// 		logger.Error("Registered Device Not Found: " + d.Device.ID.String())
// 	default:
// 		logger.Error("Unexpected Result: " + d.Device.ID.String())
// 		logger.Error("Status Code: " + strconv.Itoa(resp.StatusCode))
// 		logger.Error("Response: " + string(body))
// 	}
// 	return healthy
// }

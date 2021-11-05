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

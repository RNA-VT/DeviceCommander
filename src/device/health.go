package device

import (
	"fmt"
)

// CheckHealth probes the health endpoint of the device in question. The health
// endpoint is currently at Device.URL()/health
func (d Device) RunHealthCheck(client IDeviceClient) (Device, error) {
	resp, err := client.Health(d)
	if err != nil {
		d.logger.Warn(fmt.Sprintf("Error checking [%s] %s", d.Device.ID.String(), err))
	}

	result := client.EvaluateHealthCheckResponse(resp, d)

	if result {
		d.logger.Trace(fmt.Sprintf("device [%s] is healthy", d.Device.ID.String()))
	} else {
		d.logger.Trace(fmt.Sprintf("device [%s] is not healthy", d.Device.ID.String()))
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

func (d Device) GetSpecification(client IDeviceClient) (Device, error) {
	return Device{}, nil
}

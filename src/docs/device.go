package docs

import "github.com/rna-vt/devicecommander/src/device"

// GetAllDeviceResponse is a response for the GetAllDevices endpoint.
// swagger:response getAllDeviceResponse
type GetAllDeviceResponse struct {
	// The response contains an array of devices.
	// in: body
	Body struct {
		// Required: true
		// Example: Expected type []device.Device
		Devices []device.Device
	}
}

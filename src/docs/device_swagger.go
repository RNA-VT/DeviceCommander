package docs

import (
	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/rpc/method"
)

// swagger:route POST /device idOfDevice
// Get information about a single device.
// responses:
//   200: deviceResponse

// This text will appear as description of your response body.
// swagger:response deviceResponse
type deviceResponseWrapper struct {
	// in:body
	Body model.Device
}

// swagger:parameters idOfDevice
type deviceParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body method.GetDevicePayload
}

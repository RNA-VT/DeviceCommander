package docs

// GetAllDeviceResponse is a response for the GetAllDevices endpoint.
// swagger:response getAllDeviceResponse
type GenericError struct {
	// A Generic error response
	// in: body
	Body struct {
		// Required: true
		// Example: Expected type string
		Message string
	}
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		// Example: Expected type int
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

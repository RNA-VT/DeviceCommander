package device

import (
	"errors"
)

var ErrSpecificationRequestFailed error = errors.New("failed to request specification")
var ErrSpecificationRequestNon200 error = errors.New("specification request returned non 200 status code")
var ErrSpecificationFailedToDecode error = errors.New("specication response failed to decode")

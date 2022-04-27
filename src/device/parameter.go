package device

import (
	"github.com/google/uuid"
)

type Parameter struct {
	ID          uuid.UUID `json:"ID" faker:"uuid_hyphenated"`
	EndpointID  uuid.UUID `json:"EndpointID" faker:"uuid_hyphenated"`
	Name        string    `json:"Name"`
	Description *string   `json:"Description"`
	Type        string    `json:"Type" faker:"oneof: string, int, bool"`
}

// FromNewParameter generates a Parameter{} from a NewParameter with the correctly
// instantiated fields. This should be the primary way in which a Parameter is generated.
func FromNewParameter(input NewParameterParams) (Parameter, error) {
	endpointID, err := uuid.Parse(input.EndpointID)
	if err != nil {
		return Parameter{}, err
	}
	param := Parameter{
		ID:         uuid.New(),
		EndpointID: endpointID,
		Name:       input.Name,
		Type:       input.Type,
	}

	if input.Description != nil {
		param.Description = input.Description
	}

	return param, nil
}

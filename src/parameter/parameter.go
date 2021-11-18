package parameter

import (
	"github.com/google/uuid"

	"github.com/rna-vt/devicecommander/src/graph/model"
)

type Parameter struct {
	model.Parameter
}

// FromNewParameter generates a Parameter{} from a NewParameter with the correctly
// instantiated fields. This should be the primary way in which a Parameter is generated.
func FromNewParameter(input model.NewParameter) (model.Parameter, error) {
	endpointID, err := uuid.Parse(input.EndpointID)
	if err != nil {
		return model.Parameter{}, err
	}
	param := model.Parameter{
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

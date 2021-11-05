package endpoint

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"

	"github.com/rna-vt/devicecommander/src/graph/model"
)

// NewEndpointFromNewEndpoint generates an Endpoint from a NewEndpoint with the correctly
// instantiated fields. This should be the primary way in which an Endpoint is generated.
func NewEndpointFromNewEndpoint(input model.NewEndpoint) *model.Endpoint {
	deviceUUID, err := uuid.Parse(input.DeviceID)
	if err != nil {
		log.Error(err)
	}

	end := model.Endpoint{
		ID:       uuid.New(),
		DeviceID: deviceUUID,
		Type:     input.Type,
		Method:   input.Method,
	}

	if input.Description != nil {
		end.Description = input.Description
	}

	if input.Parameters != nil {
		end.Parameters = []model.Parameter{}

		for _, p := range input.Parameters {
			tmpP := NewParameterFromNewParameter(*p)
			end.Parameters = append(end.Parameters, tmpP)
		}
	} else {
		end.Parameters = []model.Parameter{}
	}

	return &end
}

// NewParameterFromNewParameter generates a Parameter{} from a NewParameter with the correctly
// instantiated fields. This should be the primary way in which a Parameter is generated.
func NewParameterFromNewParameter(input model.NewParameter) model.Parameter {
	endpointID, err := uuid.Parse(input.EndpointID)
	if err != nil {
		log.Error(err)
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

	return param
}

package endpoint

import (
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"

	"github.com/rna-vt/devicecommander/graph/model"
)

func NewEndpointFromNewEndpoint(input model.NewEndpoint) *model.Endpoint {
	deviceUUID, err := uuid.Parse(input.DeviceID)
	if err != nil {
		log.Error(err)
	}

	endID := uuid.New()
	end := model.Endpoint{
		ID:       endID,
		DeviceID: deviceUUID,
		Type:     input.Type,
		Method:   input.Method,
	}

	if input.Description != nil {
		end.Description = input.Description
	}

	if input.Parameters != nil {
		end.Parameters = []*model.Parameter{}

		for _, p := range input.Parameters {
			tmpP := NewParameterFromNewParameter(*p)
			end.Parameters = append(end.Parameters, tmpP)
		}
	} else {
		end.Parameters = []*model.Parameter{}
	}

	return &end
}

func NewParameterFromNewParameter(input model.NewParameter) *model.Parameter {
	param := model.Parameter{
		ID:         uuid.New().String(),
		EndpointID: input.EndpointID,
		Name:       input.Name,
		Type:       input.Type,
	}

	if input.Description != nil {
		param.Description = input.Description
	}

	return &param
}

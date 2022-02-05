package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/graph/model"
)

func (r *endpointResolver) ID(ctx context.Context, obj *model.Endpoint) (string, error) {
	return obj.ID.String(), nil
}

func (r *endpointResolver) DeviceID(ctx context.Context, obj *model.Endpoint) (string, error) {
	return obj.DeviceID.String(), nil
}

func (r *parameterResolver) ID(ctx context.Context, obj *model.Parameter) (string, error) {
	return obj.ID.String(), nil
}

func (r *parameterResolver) EndpointID(ctx context.Context, obj *model.Parameter) (string, error) {
	return obj.EndpointID.String(), nil
}

func (r *queryResolver) Endpoints(ctx context.Context) ([]*model.Endpoint, error) {
	endpoints, err := r.EndpointRepository.GetAll()
	if err != nil {
		return endpoints, err
	}
	return endpoints, nil
}

// Endpoint returns generated1.EndpointResolver implementation.
func (r *Resolver) Endpoint() generated1.EndpointResolver { return &endpointResolver{r} }

// Parameter returns generated1.ParameterResolver implementation.
func (r *Resolver) Parameter() generated1.ParameterResolver { return &parameterResolver{r} }

type (
	endpointResolver  struct{ *Resolver }
	parameterResolver struct{ *Resolver }
)

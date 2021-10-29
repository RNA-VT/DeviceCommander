package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/graph/model"
)

func (r *endpointResolver) ID(ctx context.Context, obj *model.Endpoint) (string, error) {
	return obj.ID.String(), nil
}

func (r *endpointResolver) DeviceID(ctx context.Context, obj *model.Endpoint) (string, error) {
	return obj.DeviceID.String(), nil
}

// Endpoint returns generated.EndpointResolver implementation.
func (r *Resolver) Endpoint() generated.EndpointResolver { return &endpointResolver{r} }

type endpointResolver struct{ *Resolver }

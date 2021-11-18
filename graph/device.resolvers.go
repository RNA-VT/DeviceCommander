package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rna-vt/devicecommander/src/graph/generated"
	"github.com/rna-vt/devicecommander/src/graph/model"
)

func (r *deviceResolver) ID(ctx context.Context, obj *model.Device) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.UpdateDevice) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDevice(ctx context.Context, id string) (*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	deviceResolver   struct{ *Resolver }
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

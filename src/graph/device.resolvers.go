package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/graph/model"
)

func (r *deviceResolver) ID(ctx context.Context, obj *model.Device) (string, error) {
	return obj.ID.String(), nil
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	newDevice, err := r.DeviceService.Create(input)
	if err != nil {
		return newDevice, err
	}

	return newDevice, nil
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.UpdateDevice) (string, error) {
	err := r.DeviceService.Update(input)
	if err != nil {
		return "", err
	}
	return input.ID, nil
}

func (r *mutationResolver) DeleteDevice(ctx context.Context, uuid string) (*model.Device, error) {
	newDevice, err := r.DeviceService.Delete(uuid)
	if err != nil {
		return newDevice, err
	}

	return newDevice, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	devices, err := r.DeviceService.GetAll()
	if err != nil {
		return devices, err
	}
	return devices, nil
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

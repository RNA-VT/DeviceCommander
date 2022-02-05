package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	generated1 "github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/graph/model"
)

func (r *deviceResolver) ID(ctx context.Context, obj *model.Device) (string, error) {
	return obj.ID.String(), nil
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	newDevice, err := r.DeviceRepository.Create(input)
	if err != nil {
		return newDevice, err
	}

	return newDevice, nil
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.UpdateDevice) (string, error) {
	err := r.DeviceRepository.Update(input)
	if err != nil {
		return "", err
	}
	return input.ID, nil
}

func (r *mutationResolver) DeleteDevice(ctx context.Context, id string) (*model.Device, error) {
	newDevice, err := r.DeviceRepository.Delete(id)
	if err != nil {
		return newDevice, err
	}

	return newDevice, nil
}

func (r *queryResolver) Devices(ctx context.Context) ([]*model.Device, error) {
	devices, err := r.DeviceRepository.GetAll()
	if err != nil {
		return devices, err
	}
	return devices, nil
}

// Device returns generated1.DeviceResolver implementation.
func (r *Resolver) Device() generated1.DeviceResolver { return &deviceResolver{r} }

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type (
	deviceResolver   struct{ *Resolver }
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

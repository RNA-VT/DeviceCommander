package endpoint

import "github.com/rna-vt/devicecommander/graph/model"

// Repository prototypes the required interfaces for CRUD
// management of a collection of Endpoints.
type Repository interface {
	Create(model.NewEndpoint) (*model.Endpoint, error)
	Update(model.UpdateEndpoint) error
	Delete(string) (*model.Endpoint, error)

	Get(model.Endpoint) ([]*model.Endpoint, error)
	GetAll() ([]*model.Endpoint, error)
}

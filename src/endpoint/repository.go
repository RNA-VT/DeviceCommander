package endpoint

import "github.com/rna-vt/devicecommander/src/graph/model"

// BaseRepository prototypes the required interfaces for a CRUD postgres repository.
type IEndpointCRUDRepository interface {
	Create(model.NewEndpoint) (*model.Endpoint, error)
	Update(model.UpdateEndpoint) error
	Delete(string) (*model.Endpoint, error)

	Get(model.Endpoint) ([]*model.Endpoint, error)
	GetAll() ([]*model.Endpoint, error)
}

package endpoint

// Repository prototypes the required interfaces for CRUD
// management of a collection of Endpoints.
type Repository interface {
	Create(NewEndpointParams) (*Endpoint, error)
	Update(UpdateEndpointParams) error
	Delete(string) (*Endpoint, error)
	Get(Endpoint) ([]*Endpoint, error)
	GetAll() ([]*Endpoint, error)
}

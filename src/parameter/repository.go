package parameter

import "github.com/rna-vt/devicecommander/src/graph/model"

// CRUDRepository prototypes the required interfaces for a Parameter CRUD postgres repository.
type IParameterCRUDRepository interface {
	Create(model.NewParameter) (*model.Parameter, error)
	Update(model.UpdateParameter) error
	Delete(string) (*model.Parameter, error)

	Get(model.Parameter) ([]*model.Parameter, error)
	GetAll() ([]*model.Parameter, error)
}

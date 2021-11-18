package parameter

import "github.com/rna-vt/devicecommander/src/graph/model"

// IParameterCRUDRepository prototypes the required interfaces for
// CRUD actions on a collection of Parameters.
type IParameterCRUDRepository interface {
	Create(model.NewParameter) (*model.Parameter, error)
	Update(model.UpdateParameter) error
	Delete(string) (*model.Parameter, error)

	Get(model.Parameter) ([]*model.Parameter, error)
	GetAll() ([]*model.Parameter, error)
}

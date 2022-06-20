package parameter

// Repository prototypes the required interfaces for
// CRUD actions on a collection of Parameters.
type Repository interface {
	Create(NewParameterParams) (*Parameter, error)
	Update(UpdateParameterParams) error
	Delete(string) (*Parameter, error)
	Get(Parameter) ([]*Parameter, error)
	GetAll() ([]*Parameter, error)
}

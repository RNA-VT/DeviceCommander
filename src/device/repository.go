package device

import "github.com/rna-vt/devicecommander/src/graph/model"

// DeviceCRUDRepository prototypes the required interfaces necessary to
// interact with a collection of Devices in a Postgres DB.
type IDeviceCRUDRepository interface {
	Create(model.NewDevice) (*model.Device, error)
	Update(model.UpdateDevice) error
	Delete(string) (*model.Device, error)

	Get(model.Device) ([]*model.Device, error)
	GetAll() ([]*model.Device, error)
}

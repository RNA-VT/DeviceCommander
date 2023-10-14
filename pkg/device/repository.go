package device

// DeviceCRUDRepository prototypes the required interfaces necessary to
// interact with a collection of Devices in a Postgres DB.
type Repository interface {
	Create(NewDeviceParams) (*Device, error)
	Update(UpdateDeviceParams) error
	Delete(string) (*Device, error)
	Get(Device) ([]*Device, error)
	GetAll() ([]*Device, error)
}

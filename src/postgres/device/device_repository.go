package device

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/device/parameter"
	"github.com/rna-vt/devicecommander/src/postgres"
)

// DeviceRepository implements the BaseRepository for CRUD actions involving the Devices.
type Repository struct {
	DbConfig     postgres.DBConfig
	DBConnection *gorm.DB
	Initialized  bool
	logger       *log.Entry
}

// NewDeviceRepository creates a new instance of a DeviceRepository with a DBConfig.
func NewRepository(config postgres.DBConfig) (Repository, error) {
	repository := Repository{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "repository": "device"}),
	}
	repository, err := repository.Initialize()
	if err != nil {
		return repository, err
	}
	repository.Initialized = true
	return repository, nil
}

func (r Repository) Initialize() (Repository, error) {
	db, err := postgres.GetDBConnection(r.DbConfig)
	if err != nil {
		return r, err
	}

	r.DBConnection = db

	return r, nil
}

// Create on the DeviceRepository creates a new row in the Device table.
// Due to the nested nature of Parameters.
func (r Repository) Create(newDeviceArgs device.NewDeviceParams) (*device.Device, error) {
	newDevice := device.FromNewDevice(newDeviceArgs)
	result := r.DBConnection.Create(&newDevice)
	if result.Error != nil {
		return &newDevice, result.Error
	}

	r.logger.Trace("Created device " + newDevice.ID.String())
	return &newDevice, nil
}

// Update on the DeviceRepository updates a single Device based off the ID of the UpdateDevice argument.
// It will return an error if no device is updated.
func (r Repository) Update(input device.UpdateDeviceParams) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}
	device := device.Device{ID: id}
	result := r.DBConnection.Model(device).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return postgres.NewNonExistentError("device", "update", input.ID)
	}

	r.logger.Trace("Updated device " + device.ID.String())
	return nil
}

func (r Repository) Delete(id string) (*device.Device, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return &device.Device{}, err
	}

	toBeDeleted := device.Device{
		ID: uid,
	}

	results, err := r.Get(toBeDeleted)
	if err != nil {
		return &toBeDeleted, err
	}

	if len(results) == 0 {
		return &toBeDeleted, postgres.NewNonExistentError("device", "delete", id)
	}

	for _, e := range results[0].Endpoints {
		r.DBConnection.Delete(parameter.Parameter{}, parameter.Parameter{
			EndpointID: e.ID,
		})
	}

	r.DBConnection.Delete(device.Endpoint{}, device.Endpoint{
		DeviceID: uid,
	})

	// TODO: Implement soft deletes
	r.DBConnection.Select("Endpoints").Delete(device.Device{}, toBeDeleted)

	r.logger.Trace("Deleted device " + id)
	return &toBeDeleted, nil
}

func (r Repository) Get(devQuery device.Device) ([]*device.Device, error) {
	devices := []*device.Device{}
	result := r.DBConnection.Preload(clause.Associations).Where(devQuery).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (r Repository) GetAll() ([]*device.Device, error) {
	devices := []*device.Device{}
	result := r.DBConnection.Preload(clause.Associations).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/device"
	"github.com/rna-vt/devicecommander/graph/model"
)

// BaseService prototypes the required interfaces for a CRUD postgres service.
type DeviceCRUDService interface {
	Create(model.NewDevice) (*model.Device, error)
	Update(model.UpdateDevice) error
	Delete(string) (*model.Device, error)

	Get(model.Device) ([]*model.Device, error)
	GetAll() ([]*model.Device, error)
}

// DeviceService implements the BaseService for CRUD actions involving the Devices.
type DeviceService struct {
	DbConfig     DBConfig
	DBConnection *gorm.DB
	Initialized  bool
}

// NewDeviceService creates a new instance of a DeviceService with a DBConfig.
func NewDeviceService(config DBConfig) (DeviceService, error) {
	service := DeviceService{
		DbConfig:    config,
		Initialized: false,
	}
	service, err := service.Initialise()
	if err != nil {
		return service, err
	}
	service.Initialized = true
	return service, nil
}

func (s DeviceService) Initialise() (DeviceService, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", s.DbConfig.Host, s.DbConfig.UserName, s.DbConfig.Password, s.DbConfig.Name, s.DbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return s, err
	}

	s.DBConnection = db

	err = RunMigration(db)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (s DeviceService) Create(newDeviceArgs model.NewDevice) (*model.Device, error) {
	logger := getPostgresLogger()
	newDevice := device.NewDeviceFromNewDevice(newDeviceArgs)
	result := s.DBConnection.Create(&newDevice)
	if result.Error != nil {
		return &newDevice, result.Error
	}

	logger.Debug("Created device " + newDevice.ID.String())
	return &newDevice, nil
}

func (s DeviceService) Update(input model.UpdateDevice) error {
	logger := getPostgresLogger()
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}
	device := model.Device{ID: id}
	result := s.DBConnection.Model(device).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	logger.Debug("Updated device " + device.ID.String())
	return nil
}

func (s DeviceService) Delete(id string) (*model.Device, error) {
	logger := getPostgresLogger()
	var toBeDeleted model.Device
	result := s.DBConnection.First(&toBeDeleted, "ID = ?", id)
	if result.Error != nil {
		return &toBeDeleted, result.Error
	}

	logger.Debug(toBeDeleted)

	uid, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = uid

	results, err := s.Get(toBeDeleted)
	if err != nil {
		return &toBeDeleted, err
	}

	if len(results) == 0 {
		return &toBeDeleted, fmt.Errorf("the device %s has already been deleted", id)
	}

	for _, e := range results[0].Endpoints {
		s.DBConnection.Delete(model.Parameter{}, model.Parameter{
			EndpointID: e.ID,
		})
	}

	s.DBConnection.Delete(model.Endpoint{}, model.Endpoint{
		DeviceID: uid,
	})

	// TODO: Implement soft deletes
	s.DBConnection.Select("Endpoints").Delete(model.Device{}, toBeDeleted)

	logger.Debug("Deleted device " + id)
	return &toBeDeleted, nil
}

func (s DeviceService) Get(devQuery model.Device) ([]*model.Device, error) {
	devices := []*model.Device{}
	result := s.DBConnection.Preload(clause.Associations).Where(devQuery).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (s DeviceService) GetAll() ([]*model.Device, error) {
	devices := []*model.Device{}
	result := s.DBConnection.Preload(clause.Associations).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/rna-vt/devicecommander/graph/model"
)

// BaseService prototypes the required interfaces for a CRUD postgres service.
type DeviceCRUDService interface {
	Initialise() (*gorm.DB, error)
	Create(model.NewDevice) (*model.Device, error)
	Update(model.UpdateDevice) error
	Delete(string) (*model.Device, error)

	Get(model.Device) ([]*model.Device, error)
	GetAll() ([]*model.Device, error)
}

// DeviceService implements the BaseService for CRUD actions involving the Devices.
type MockDeviceService struct {
	DbConfig         DBConfig
	dBConnection     *gorm.DB
	Initialized      bool
	DeviceCollection []model.Device
}

// NewDeviceService creates a new instance of a DeviceService with a DBConfig.
func NewMockDeviceService(config DBConfig) (DeviceService, error) {
	service := DeviceService{
		DbConfig:    config,
		Initialized: false,
	}
	db, err := service.Initialise()
	if err != nil {
		return service, err
	}
	service.dBConnection = db
	service.Initialized = true
	return service, nil
}

func (s MockDeviceService) Create(newDeviceArgs model.NewDevice) (*model.Device, error) {
	logger := getPostgresLogger()
	newDevice := model.Device{
		ID:   uuid.New(),
		Host: newDeviceArgs.Host,
		Port: newDeviceArgs.Port,
	}

	if newDeviceArgs.Mac != nil {
		newDevice.MAC = *newDeviceArgs.Mac
	}

	if newDeviceArgs.Name != nil {
		newDevice.Name = *newDeviceArgs.Name
	}

	if newDeviceArgs.Description != nil {
		newDevice.Description = *newDeviceArgs.Description
	}

	result := s.dBConnection.Create(&newDevice)
	if result.Error != nil {
		return &newDevice, result.Error
	}

	logger.Debug("Created device " + newDevice.ID.String())
	return &newDevice, nil
}

func (s MockDeviceService) Update(input model.UpdateDevice) error {
	return nil
}

func (s MockDeviceService) Delete(id string) (*model.Device, error) {
	return &model.Device{}, nil
}

func (s MockDeviceService) Get(devQuery model.Device) ([]*model.Device, error) {
	devices := []*model.Device{}
	result := s.dBConnection.Where(devQuery).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (s MockDeviceService) GetAll() ([]*model.Device, error) {
	return s.DeviceCollection, nil
}

// func (s MockDeviceService) Initialise() (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", s.DbConfig.Host, s.DbConfig.UserName, s.DbConfig.Password, s.DbConfig.Name, s.DbConfig.Port)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return &gorm.DB{}, err
// 	}

// 	err = db.AutoMigrate(&model.Device{})
// 	if err != nil {
// 		return &gorm.DB{}, err
// 	}

// 	return db, nil
// }

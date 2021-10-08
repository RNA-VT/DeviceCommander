package postgres

import (
	"fmt"

	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rna-vt/devicecommander/graph/model"
)

type BaseService interface {
	Initialise() error
	Create(model.NewDevice) (*model.Device, error)
	Update(model.UpdateDevice) error
	Delete(model.Device) (*model.Device, error)

	Get(model.Device) (*model.Device, error)
	GetAll(host string, port int) ([]*model.Device, error)
}

type DeviceService struct {
	DbConfig     DbConfig
	dBConnection *gorm.DB
	Initialized  bool
}

func NewDeviceService(config DbConfig) (DeviceService, error) {
	service := DeviceService{
		DbConfig:    config,
		Initialized: false,
	}
	err := service.Initialise()
	if err != nil {
		return service, err
	}
	service.Initialized = true
	return service, nil
}

func (s *DeviceService) Create(newDeviceArgs model.NewDevice) (*model.Device, error) {
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

func (s *DeviceService) Update(input model.UpdateDevice) error {
	logger := getPostgresLogger()
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}
	device := model.Device{ID: id}
	result := s.dBConnection.Model(device).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	logger.Debug("Updated device " + device.ID.String())
	return nil
}

func (s *DeviceService) Delete(id string) (*model.Device, error) {
	logger := getPostgresLogger()
	var toBeDeleted model.Device
	result := s.dBConnection.First(&toBeDeleted, "ID = ?", id)
	if result.Error != nil {
		return &toBeDeleted, result.Error
	}

	logger.Debug(toBeDeleted)

	uid, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = uid

	// TODO: Implement soft deletes
	s.dBConnection.Delete(model.Device{}, toBeDeleted)

	logger.Debug("Deleted device " + id)
	return &toBeDeleted, nil
}

func (s *DeviceService) Get(devQuery model.Device) ([]*model.Device, error) {
	devices := []*model.Device{}
	result := s.dBConnection.Where(devQuery).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (s *DeviceService) GetAll() ([]*model.Device, error) {
	devices := []*model.Device{}
	result := s.dBConnection.Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}

	return devices, nil
}

func (s *DeviceService) Initialise() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", s.DbConfig.Host, s.DbConfig.UserName, s.DbConfig.Password, s.DbConfig.Name, s.DbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	s.dBConnection = db

	err = db.AutoMigrate(&model.Device{})
	if err != nil {
		return err
	}

	return nil
}

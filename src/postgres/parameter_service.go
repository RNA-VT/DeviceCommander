package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/endpoint"
	"github.com/rna-vt/devicecommander/graph/model"
)

// ParameterCRUDService prototypes the required interfaces for a Parameter CRUD postgres service.
type ParameterCRUDService interface {
	Create(model.NewParameter) (*model.Parameter, error)
	Update(model.UpdateParameter) error
	Delete(string) (*model.Parameter, error)

	Get(model.Parameter) ([]*model.Parameter, error)
	GetAll() ([]*model.Parameter, error)
}

// ParameterService implements the BaseService for CRUD actions involving the Devices.
type ParameterService struct {
	DbConfig     DBConfig
	DBConnection *gorm.DB
	Initialized  bool
}

// NewParameterService creates a new instance of a DeviceService with a DBConfig.
func NewParameterService(config DBConfig) (ParameterService, error) {
	service := ParameterService{
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

func (s ParameterService) Initialise() (ParameterService, error) {
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

func (s ParameterService) Create(newParameterArgs model.NewParameter) (*model.Parameter, error) {
	logger := getPostgresLogger()
	newParameter := endpoint.NewParameterFromNewParameter(newParameterArgs)

	result := s.DBConnection.Create(&newParameter)
	if result.Error != nil {
		return &newParameter, result.Error
	}

	logger.Debug("Created Parameter " + newParameter.ID.String())
	return &newParameter, nil
}

func (s ParameterService) Update(input model.UpdateParameter) error {
	logger := getPostgresLogger()

	parameterID, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := model.Parameter{ID: parameterID}

	result := s.DBConnection.Session(&gorm.Session{FullSaveAssociations: true}).Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	logger.Debug("Updated Parameter " + end.ID.String())
	return nil
}

func (s ParameterService) Delete(id string) (*model.Parameter, error) {
	logger := getPostgresLogger()
	var toBeDeleted model.Parameter

	parameterID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = parameterID

	results, err := s.Get(toBeDeleted)
	if err != nil {
		return &toBeDeleted, err
	}

	if len(results) == 0 {
		return &toBeDeleted, fmt.Errorf("parameter %s has already been deleted", id)
	}

	toBeDeleted = *results[0]

	// TODO: Implement soft deletes
	s.DBConnection.Delete(model.Parameter{}, toBeDeleted)

	logger.Debug("Deleted Parameter " + id)
	return &toBeDeleted, nil
}

func (s ParameterService) Get(query model.Parameter) ([]*model.Parameter, error) {
	Parameters := []*model.Parameter{}
	result := s.DBConnection.Preload(clause.Associations).Where(query).Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

func (s ParameterService) GetAll() ([]*model.Parameter, error) {
	Parameters := []*model.Parameter{}
	result := s.DBConnection.Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

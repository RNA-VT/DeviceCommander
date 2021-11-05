package postgres

import (
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/src/endpoint"
	"github.com/rna-vt/devicecommander/src/graph/model"
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
	logger       *log.Entry
}

// NewParameterService creates a new instance of a DeviceService with a DBConfig.
func NewParameterService(config DBConfig) (ParameterService, error) {
	service := ParameterService{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "service": "parameter"}),
	}
	service, err := service.Initialise()
	if err != nil {
		return service, err
	}
	service.Initialized = true
	return service, nil
}

// Initialise on the ParameterService struct opens the postgres connection defined in the ParameterService.DBConfig
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

// Create on the ParameterService creates a new row in the Parameter table in the postgres database.
func (s ParameterService) Create(newParameterArgs model.NewParameter) (*model.Parameter, error) {
	newParameter := endpoint.NewParameterFromNewParameter(newParameterArgs)

	result := s.DBConnection.Create(&newParameter)
	if result.Error != nil {
		return &newParameter, result.Error
	}

	s.logger.Trace("Created Parameter " + newParameter.ID.String())
	return &newParameter, nil
}

// Update on the ParameterService updates a new single row in the Parameter table according to the specified UpdateParameter.ID.
func (s ParameterService) Update(input model.UpdateParameter) error {
	parameterID, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := model.Parameter{ID: parameterID}

	result := s.DBConnection.Session(&gorm.Session{FullSaveAssociations: true}).Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	s.logger.Trace("Updated Parameter " + end.ID.String())
	return nil
}

// Delete on the ParameterService removes a single row from the Parameter table by the specific ID.
func (s ParameterService) Delete(id string) (*model.Parameter, error) {
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

	s.logger.Trace("Deleted Parameter " + id)
	return &toBeDeleted, nil
}

// Get on the ParameterService will retrieve all of the rows that match the query. The
// associated objects (endpoints) will be preloaded for convenience.
func (s ParameterService) Get(query model.Parameter) ([]*model.Parameter, error) {
	Parameters := []*model.Parameter{}
	result := s.DBConnection.Preload(clause.Associations).Where(query).Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

// GetAll on the ParameterService will retrieve all of the rows in the Parameter table. The
// associated objects (endpoints) will be preloaded for convenience.
func (s ParameterService) GetAll() ([]*model.Parameter, error) {
	Parameters := []*model.Parameter{}
	result := s.DBConnection.Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

package parameter

import (
	"fmt"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/src/device/parameter"
	"github.com/rna-vt/devicecommander/src/postgres"
)

// ParameterRepository implements the BaseRepository for CRUD actions involving the Devices.
type Repository struct {
	DbConfig     postgres.DBConfig
	DBConnection *gorm.DB
	Initialized  bool
	logger       *log.Entry
}

// NewParameterRepository creates a new instance of a DeviceRepository with a DBConfig.
func NewParameterRepository(config postgres.DBConfig) (Repository, error) {
	repository := Repository{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "repository": "parameter"}),
	}
	repository, err := repository.Initialize()
	if err != nil {
		return repository, err
	}
	repository.Initialized = true
	return repository, nil
}

// Initialize on the ParameterRepository struct opens the postgres connection defined in the ParameterRepository.DBConfig.
func (r Repository) Initialize() (Repository, error) {
	db, err := postgres.GetDBConnection(r.DbConfig)
	if err != nil {
		return r, err
	}

	r.DBConnection = db

	return r, nil
}

// Create on the ParameterRepository creates a new row in the Parameter table in the postgres database.
func (r Repository) Create(newParameterArgs parameter.NewParameterParams) (*parameter.Parameter, error) {
	newParameter, err := parameter.FromNewParameter(newParameterArgs)
	if err != nil {
		return &newParameter, err
	}

	result := r.DBConnection.Create(&newParameter)
	if result.Error != nil {
		return &newParameter, result.Error
	}

	r.logger.Trace("Created Parameter " + newParameter.ID.String())
	return &newParameter, nil
}

// Update on the ParameterRepository updates a new single row in the Parameter table according to the specified UpdateParameter.ID.
func (r Repository) Update(input parameter.UpdateParameterParams) error {
	parameterID, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := parameter.Parameter{ID: parameterID}

	result := r.DBConnection.Session(&gorm.Session{FullSaveAssociations: true}).Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	r.logger.Trace("Updated Parameter " + end.ID.String())
	return nil
}

// Delete on the ParameterRepository removes a single row from the Parameter table by the specific ID.
func (r Repository) Delete(id string) (*parameter.Parameter, error) {
	var toBeDeleted parameter.Parameter

	parameterID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = parameterID

	results, err := r.Get(toBeDeleted)
	if err != nil {
		return &toBeDeleted, err
	}

	if len(results) == 0 {
		return &toBeDeleted, fmt.Errorf("parameter %s has already been deleted", id)
	}

	toBeDeleted = *results[0]

	// TODO: Implement soft deletes
	r.DBConnection.Delete(parameter.Parameter{}, toBeDeleted)

	r.logger.Trace("Deleted Parameter " + id)
	return &toBeDeleted, nil
}

// Get on the ParameterRepository will retrieve all of the rows that match the query. The
// associated object (endpoint) will be preloaded for convenience.
func (r Repository) Get(query parameter.Parameter) ([]*parameter.Parameter, error) {
	Parameters := []*parameter.Parameter{}
	result := r.DBConnection.Preload(clause.Associations).Where(query).Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

// GetAll on the ParameterRepository will retrieve all of the rows in the Parameter table. The
// associated objects (endpoints) will be preloaded for convenience.
func (r Repository) GetAll() ([]*parameter.Parameter, error) {
	Parameters := []*parameter.Parameter{}
	result := r.DBConnection.Find(&Parameters)
	if result.Error != nil {
		return Parameters, result.Error
	}

	return Parameters, nil
}

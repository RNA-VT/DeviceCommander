package endpoint

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
)

// EndpointRepository implements the BaseRepository for CRUD actions involving Endpoints.
type Repository struct {
	DbConfig     postgres.DBConfig
	DBConnection *gorm.DB
	Initialized  bool
	logger       *log.Entry
}

// NewEndpointRepository creates a new instance of an EndpointRepository with a DBConfig.
func NewRepository(config postgres.DBConfig) (Repository, error) {
	repository := Repository{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "repository": "endpoint"}),
	}
	repository, err := repository.Initialize()
	if err != nil {
		return repository, err
	}
	repository.Initialized = true
	return repository, nil
}

// Initialize on the EndpointRepository struct opens the postgres connection
// defined in the EndpointRepository.DBConfig.
func (r Repository) Initialize() (Repository, error) {
	db, err := postgres.GetDBConnection(r.DbConfig)
	if err != nil {
		return r, err
	}

	r.DBConnection = db

	return r, nil
}

func (r Repository) Create(newDeviceArgs device.NewEndpointParams) (*device.Endpoint, error) {
	newEndpoint, err := device.FromNewEndpoint(newDeviceArgs)
	if err != nil {
		return &newEndpoint, err
	}

	result := r.DBConnection.Create(&newEndpoint)
	if result.Error != nil {
		return &newEndpoint, result.Error
	}

	r.logger.Trace("Created endpoint " + newEndpoint.ID.String())
	return &newEndpoint, nil
}

func (r Repository) Update(input device.UpdateEndpointParams) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := device.Endpoint{ID: id}

	result := r.DBConnection.Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return postgres.NewNonExistentError("endpoint", "update", input.ID)
	}

	r.logger.Trace("Updated endpoint " + end.ID.String())
	return nil
}

// Delete on the EndpointRepository removes a single row from the Endpoint table by the
// specific ID AND all of the Parameters associated with the EndpointID.
func (r Repository) Delete(id string) (*device.Endpoint, error) {
	var toBeDeleted device.Endpoint
	endUUID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = endUUID

	r.DBConnection.Delete(device.Parameter{}, device.Parameter{
		EndpointID: endUUID,
	})

	r.DBConnection.Delete(device.Endpoint{}, toBeDeleted)

	r.logger.Debug("Deleted endpoint " + id)
	return &toBeDeleted, nil
}

// Get on the EndpointRepository will retrieve all of the rows that match the query. The
// associated objects (parameters) will be preloaded for convenience.
func (r Repository) Get(query device.Endpoint) ([]*device.Endpoint, error) {
	endpoints := []*device.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Where(query).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

func (r Repository) GetAll() ([]*device.Endpoint, error) {
	endpoints := []*device.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

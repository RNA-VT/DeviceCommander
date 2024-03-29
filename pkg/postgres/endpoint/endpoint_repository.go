package endpoint

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/rna-vt/devicecommander/pkg/device/endpoint"
	"github.com/rna-vt/devicecommander/pkg/device/endpoint/parameter"
	"github.com/rna-vt/devicecommander/pkg/postgres"
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

func (r Repository) Create(newDeviceArgs endpoint.NewEndpointParams) (*endpoint.Endpoint, error) {
	newEndpoint, err := endpoint.FromNewEndpoint(newDeviceArgs)
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

func (r Repository) Update(input endpoint.UpdateEndpointParams) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := endpoint.Endpoint{ID: id}

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
func (r Repository) Delete(id string) (*endpoint.Endpoint, error) {
	var toBeDeleted endpoint.Endpoint
	endUUID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = endUUID

	r.DBConnection.Delete(parameter.Parameter{}, parameter.Parameter{
		EndpointID: endUUID,
	})

	r.DBConnection.Delete(endpoint.Endpoint{}, toBeDeleted)

	r.logger.Debug("Deleted endpoint " + id)
	return &toBeDeleted, nil
}

// Get on the EndpointRepository will retrieve all of the rows that match the query. The
// associated objects (parameters) will be preloaded for convenience.
func (r Repository) Get(query endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Where(query).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

func (r Repository) GetAll() ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

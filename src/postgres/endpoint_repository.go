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

// EndpointRepository implements the BaseRepository for CRUD actions involving Endpoints.
type EndpointRepository struct {
	DbConfig     DBConfig
	DBConnection *gorm.DB
	Initialized  bool
	logger       *log.Entry
}

// NewEndpointRepository creates a new instance of an EndpointRepository with a DBConfig.
func NewEndpointRepository(config DBConfig) (EndpointRepository, error) {
	repository := EndpointRepository{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "repository": "endpoint"}),
	}
	repository, err := repository.Initialise()
	if err != nil {
		return repository, err
	}
	repository.Initialized = true
	return repository, nil
}

// Initialise on the EndpointRepository struct opens the postgres connection
// defined in the EndpointRepository.DBConfig.
func (r EndpointRepository) Initialise() (EndpointRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", r.DbConfig.Host, r.DbConfig.UserName, r.DbConfig.Password, r.DbConfig.Name, r.DbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return r, err
	}

	r.DBConnection = db

	err = RunMigration(db)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (r EndpointRepository) Create(newDeviceArgs model.NewEndpoint) (*model.Endpoint, error) {
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

func (r EndpointRepository) Update(input model.UpdateEndpoint) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := model.Endpoint{ID: id}

	result := r.DBConnection.Session(&gorm.Session{FullSaveAssociations: true}).Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	r.logger.Trace("Updated endpoint " + end.ID.String())
	return nil
}

// Delete on the EndpointRepository removes a single row from the Endpoint table by the
// specific ID AND all of the Parameters associated with the EndpointID.
func (r EndpointRepository) Delete(id string) (*model.Endpoint, error) {
	var toBeDeleted model.Endpoint
	endUUID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = endUUID

	r.DBConnection.Delete(model.Parameter{}, model.Parameter{
		EndpointID: endUUID,
	})

	r.DBConnection.Delete(model.Endpoint{}, toBeDeleted)

	r.logger.Debug("Deleted endpoint " + id)
	return &toBeDeleted, nil
}

// Get on the EndpointRepository will retrieve all of the rows that match the query. The
// associated objects (parameters) will be preloaded for convenience.
func (r EndpointRepository) Get(query model.Endpoint) ([]*model.Endpoint, error) {
	endpoints := []*model.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Where(query).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

func (r EndpointRepository) GetAll() ([]*model.Endpoint, error) {
	endpoints := []*model.Endpoint{}
	result := r.DBConnection.Preload(clause.Associations).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

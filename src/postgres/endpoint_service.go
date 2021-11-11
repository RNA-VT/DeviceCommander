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

// BaseService prototypes the required interfaces for a CRUD postgres service.
type EndpointCRUDService interface {
	Create(model.NewEndpoint) (*model.Endpoint, error)
	Update(model.UpdateEndpoint) error
	Delete(string) (*model.Endpoint, error)

	Get(model.Endpoint) ([]*model.Endpoint, error)
	GetAll() ([]*model.Endpoint, error)
}

// EndpointService implements the BaseService for CRUD actions involving Endpoints.
type EndpointService struct {
	DbConfig     DBConfig
	DBConnection *gorm.DB
	Initialized  bool
	logger       *log.Entry
}

// NewEndpointService creates a new instance of an EndpointService with a DBConfig.
func NewEndpointService(config DBConfig) (EndpointService, error) {
	service := EndpointService{
		DbConfig:    config,
		Initialized: false,
		logger:      log.WithFields(log.Fields{"module": "postgres", "service": "endpoint"}),
	}
	service, err := service.Initialise()
	if err != nil {
		return service, err
	}
	service.Initialized = true
	return service, nil
}

// Initialise on the EndpointService struct opens the postgres connection defined in the EndpointService.DBConfig.
func (s EndpointService) Initialise() (EndpointService, error) {
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

func (s EndpointService) Create(newDeviceArgs model.NewEndpoint) (*model.Endpoint, error) {
	newEndpoint := endpoint.EndpointFromNewEndpoint(newDeviceArgs)

	result := s.DBConnection.Create(&newEndpoint)
	if result.Error != nil {
		return newEndpoint, result.Error
	}

	s.logger.Trace("Created endpoint " + newEndpoint.ID.String())
	return newEndpoint, nil
}

func (s EndpointService) Update(input model.UpdateEndpoint) error {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return err
	}

	end := model.Endpoint{ID: id}

	result := s.DBConnection.Session(&gorm.Session{FullSaveAssociations: true}).Model(end).Updates(input)
	if result.Error != nil {
		return result.Error
	}

	s.logger.Trace("Updated endpoint " + end.ID.String())
	return nil
}

// Delete on the EndpointService removes a single row from the Endpoint table by the specific ID AND all of the Parameters
// associated with the EndpointID.
func (s EndpointService) Delete(id string) (*model.Endpoint, error) {
	var toBeDeleted model.Endpoint
	endUUID, err := uuid.Parse(id)
	if err != nil {
		return &toBeDeleted, err
	}
	toBeDeleted.ID = endUUID

	s.DBConnection.Delete(model.Parameter{}, model.Parameter{
		EndpointID: endUUID,
	})

	s.DBConnection.Delete(model.Endpoint{}, toBeDeleted)

	s.logger.Debug("Deleted endpoint " + id)
	return &toBeDeleted, nil
}

// Get on the EndpointService will retrieve all of the rows that match the query. The
// associated objects (parameters) will be preloaded for convenience.
func (s EndpointService) Get(query model.Endpoint) ([]*model.Endpoint, error) {
	endpoints := []*model.Endpoint{}
	result := s.DBConnection.Preload(clause.Associations).Where(query).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

func (s EndpointService) GetAll() ([]*model.Endpoint, error) {
	endpoints := []*model.Endpoint{}
	result := s.DBConnection.Preload(clause.Associations).Find(&endpoints)
	if result.Error != nil {
		return endpoints, result.Error
	}

	return endpoints, nil
}

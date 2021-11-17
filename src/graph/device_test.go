package graph

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/rna-vt/devicecommander/src/graph/model"
	"github.com/rna-vt/devicecommander/src/mocks"
)

type DeviceGraphQLSuite struct {
	suite.Suite
	resolver             Resolver
	mockDeviceRepository mocks.IDeviceCRUDRepository
	ctx                  context.Context
}

func (s *DeviceGraphQLSuite) SetupSuite() {
	s.mockDeviceRepository = mocks.IDeviceCRUDRepository{}
	s.resolver = Resolver{
		DeviceRepository: &s.mockDeviceRepository,
	}
	s.ctx = context.Background()
}

func (s *DeviceGraphQLSuite) TestCreateDevice() {
	mutator := s.resolver.Mutation()

	newDevice := model.NewDevice{
		Host: "0.0.0.0",
		Port: 0000,
	}

	s.mockDeviceRepository.On("Create", newDevice).Return(&model.Device{}, nil)
	_, err := mutator.CreateDevice(s.ctx, newDevice)
	assert.Nil(s.T(), err)

	s.mockDeviceRepository.AssertCalled(s.T(), "Create", newDevice)

	s.mockDeviceRepository.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestGetDevices() {
	queryResolver := s.resolver.Query()

	s.mockDeviceRepository.On("GetAll").Return([]*model.Device{}, nil)
	_, err := queryResolver.Devices(s.ctx)
	assert.Nil(s.T(), err)

	s.mockDeviceRepository.AssertCalled(s.T(), "GetAll")

	s.mockDeviceRepository.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestDeleteDevice() {
	mutator := s.resolver.Mutation()

	randomUUID := uuid.New().String()

	s.mockDeviceRepository.On("Delete", randomUUID).Return(&model.Device{}, nil)
	_, err := mutator.DeleteDevice(s.ctx, randomUUID)
	assert.Nil(s.T(), err)

	s.mockDeviceRepository.AssertCalled(s.T(), "Delete", randomUUID)

	s.mockDeviceRepository.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestUpdateDevice() {
	mutator := s.resolver.Mutation()
	tmpName := "McTesterson"
	updateInput := model.UpdateDevice{
		ID:   "uuid.string",
		Name: &tmpName,
	}

	s.mockDeviceRepository.On("Update", updateInput).Return(nil)
	_, err := mutator.UpdateDevice(s.ctx, updateInput)
	assert.Nil(s.T(), err)

	s.mockDeviceRepository.AssertCalled(s.T(), "Update", updateInput)

	s.mockDeviceRepository.AssertExpectations(s.T())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeviceGraphQLTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceGraphQLSuite))
}

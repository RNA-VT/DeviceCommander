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
	resolver          Resolver
	mockDeviceService mocks.DeviceCRUDService
	ctx               context.Context
}

func (s *DeviceGraphQLSuite) SetupSuite() {
	s.mockDeviceService = mocks.DeviceCRUDService{}
	s.resolver = Resolver{
		DeviceService: &s.mockDeviceService,
	}
	s.ctx = context.Background()
}

func (s *DeviceGraphQLSuite) TestCreateDevice() {
	mutator := s.resolver.Mutation()

	newDevice := model.NewDevice{
		Host: "0.0.0.0",
		Port: 0o000,
	}

	s.mockDeviceService.On("Create", newDevice).Return(&model.Device{}, nil)
	_, err := mutator.CreateDevice(s.ctx, newDevice)
	assert.Nil(s.T(), err)

	s.mockDeviceService.AssertCalled(s.T(), "Create", newDevice)

	s.mockDeviceService.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestGetDevices() {
	queryResolver := s.resolver.Query()

	s.mockDeviceService.On("GetAll").Return([]*model.Device{}, nil)
	_, err := queryResolver.Devices(s.ctx)
	assert.Nil(s.T(), err)

	s.mockDeviceService.AssertCalled(s.T(), "GetAll")

	s.mockDeviceService.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestDeleteDevice() {
	mutator := s.resolver.Mutation()

	randomUUID := uuid.New().String()

	s.mockDeviceService.On("Delete", randomUUID).Return(&model.Device{}, nil)
	_, err := mutator.DeleteDevice(s.ctx, randomUUID)
	assert.Nil(s.T(), err)

	s.mockDeviceService.AssertCalled(s.T(), "Delete", randomUUID)

	s.mockDeviceService.AssertExpectations(s.T())
}

func (s *DeviceGraphQLSuite) TestUpdateDevice() {
	mutator := s.resolver.Mutation()
	tmpName := "McTesterson"
	updateInput := model.UpdateDevice{
		ID:   "uuid.string",
		Name: &tmpName,
	}

	s.mockDeviceService.On("Update", updateInput).Return(nil)
	_, err := mutator.UpdateDevice(s.ctx, updateInput)
	assert.Nil(s.T(), err)

	s.mockDeviceService.AssertCalled(s.T(), "Update", updateInput)

	s.mockDeviceService.AssertExpectations(s.T())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeviceGraphQLTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceGraphQLSuite))
}

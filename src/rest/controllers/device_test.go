package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	mocks "github.com/rna-vt/devicecommander/mocks/device"
	"github.com/rna-vt/devicecommander/src/device"
)

type DeviceControllerSuite struct {
	suite.Suite
	controller           DeviceController
	mockDeviceRepository mocks.Repository
	ctx                  context.Context
}

func (s *DeviceControllerSuite) SetupSuite() {
	s.mockDeviceRepository = mocks.Repository{}
	s.controller = DeviceController{
		Repository: &s.mockDeviceRepository,
	}
	s.ctx = context.Background()
}

func NewEchoContext(method, path string, data interface{}) echo.Context {
	requestByte, _ := json.Marshal(data)
	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest(method, path, requestReader)
	if err != nil {
		logrus.Error(err)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	res := httptest.NewRecorder()
	return e.NewContext(req, res)
}

func (s *DeviceControllerSuite) TestCreateDevice() {
	newDevices := device.GenerateRandomNewDeviceParams(1)

	s.mockDeviceRepository.On("Create", newDevices[0]).Return(&device.Device{}, nil)

	err := s.controller.Create(NewEchoContext("POST", "/v1/device", newDevices[0]))
	assert.Nil(s.T(), err)
}

func (s *DeviceControllerSuite) TestGetDevices() {

	s.mockDeviceRepository.On("GetAll").Return([]*device.Device{}, nil)
	err := s.controller.GetAll(NewEchoContext("GET", "/v1/device", nil))
	assert.Nil(s.T(), err)
}

func (s *DeviceControllerSuite) TestDeleteDevice() {
	randomUUID := uuid.New().String()

	ctx := NewEchoContext("DELETE", "/v1/device/:id", nil)

	ctx.SetParamNames("id")
	ctx.SetParamValues(randomUUID)

	s.mockDeviceRepository.On("Delete", randomUUID).Return(&device.Device{}, nil)
	err := s.controller.Delete(ctx)
	assert.Nil(s.T(), err)
}

func (s *DeviceControllerSuite) TestUpdateDevice() {
	tmpName := "McTesterson"
	updateInput := device.UpdateDeviceParams{
		ID:   "uuid.string",
		Name: &tmpName,
	}
	ctx := NewEchoContext("POST", "/v1/device", updateInput)

	s.mockDeviceRepository.On("Update", updateInput).Return(nil)
	err := s.controller.Update(ctx)
	assert.Nil(s.T(), err)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run.
func TestDeviceRESTControllerTestSuite(t *testing.T) {
	suite.Run(t, new(DeviceControllerSuite))
}

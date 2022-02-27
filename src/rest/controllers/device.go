package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/src/postgres/device"
)

type DeviceController struct {
	Repository device.Repository
}

func NewDeviceController() (DeviceController, error) {
	dbConfig := postgres.GetDBConfigFromEnv()
	deviceRepository, err := postgresDevice.NewRepository(dbConfig)
	if err != nil {
		log.Error(err)
		return DeviceController{}, err
	}

	return DeviceController{
		Repository: deviceRepository,
	}, nil
}

func (controller DeviceController) Create(c echo.Context) error {
	dev, err := device.BasicDevice{}.NewDeviceFromRequestBody(c.Request().Body)
	if err != nil {
		return err
	}

	newDevice, err := controller.Repository.Create(dev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newDevice)
}

func (controller DeviceController) GetAll(c echo.Context) error {
	devices, err := controller.Repository.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, devices)
}

func (controller DeviceController) GetDevice(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	tmpDev := model.Device{ID: id}
	device, err := controller.Repository.Get(tmpDev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &device)
}

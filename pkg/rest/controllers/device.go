package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/pkg/postgres/device"
	"github.com/rna-vt/devicecommander/pkg/utils"
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
	dev, err := device.NewDeviceFromRequestBody(c.Request().Body)
	if err != nil {
		return err
	}

	newDevice, err := controller.Repository.Create(dev)
	if err != nil {
		return err
	}

	fmt.Println(utils.PrettyPrintJSON(dev))
	fmt.Println(utils.PrettyPrintJSON(newDevice))

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

	tmpDev := device.Device{ID: id}
	device, err := controller.Repository.Get(tmpDev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &device)
}

func (controller DeviceController) Update(c echo.Context) error {
	updateParams := device.UpdateDeviceParams{}
	if err := c.Bind(&updateParams); err != nil {
		return err
	}

	err := controller.Repository.Update(updateParams)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}

func (controller DeviceController) Delete(c echo.Context) error {
	deviceID := c.Param("id")
	log.Infof("DEVICE_ID='%s'", deviceID)
	toDelete, err := controller.Repository.Delete(deviceID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &toDelete)
}

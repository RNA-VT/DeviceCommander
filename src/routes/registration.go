package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/src/device"
)

func (a *APIService) addRegistrationRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.POST("/join_network", a.joinNetwork)
}

// joinNetwork enables self sign-up for devices that have previously connected
// and devices that need to be manually added.
func (a *APIService) joinNetwork(c echo.Context) error {
	log.Println("Device asked to join cluster")

	dev, err := device.BasicDevice{}.NewDeviceFromRequestBody(c.Request().Body)
	if err != nil {
		return err
	}

	newDevice, err := a.DeviceRepository.Create(dev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newDevice)
}

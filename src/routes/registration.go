package routes

import (
	"devicecommander/device"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a *APIService) addRegistrationRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.POST("/join_network", a.joinNetwork)
}

// joinNetwork enables self sign-up for devices that have previously connected
// and devices that need to be manually added
func (a *APIService) joinNetwork(c echo.Context) error {
	log.Println("Device asked to join cluster")

	dev, err := device.NewDeviceFromRequestBody(c.Request().Body)
	if err != nil {
		return err
	}

	a.Cluster.AddDevice(dev)

	return c.JSON(http.StatusOK, "Registered.")
}

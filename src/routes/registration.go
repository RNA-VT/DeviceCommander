package routes

import (
	"devicecommander/cluster"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a *APIService) addRegistrationRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.POST("/join_network", a.joinNetwork)
}

func (a *APIService) joinNetwork(c echo.Context) error {
	log.Println("Device asked to join cluster")

	dev := cluster.DeviceFromRegistrationRequestBody(c.Request().Body)
	a.Cluster.AddDevice(dev)

	return c.JSON(http.StatusOK, "Registered.")
}

package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func (a *APIService) addInfoRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/cluster_info", a.getClusterInfo)
	api.GET("/device", a.getDevices)
	api.GET("/device/:id", a.getDevice)
	api.GET("/health", a.health)
}

func (a APIService) health(c echo.Context) error {
	return c.JSON(http.StatusOK, "I'm Alive")
}

func (a APIService) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster)
}

func (a APIService) getDevices(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster.Devices)
}

func (a APIService) getDevice(c echo.Context) error {
	devices := a.Cluster.GetDevices()
	dev, ok := devices[c.Param("id")]
	if !ok {
		return c.JSON(http.StatusNotFound, "ID Not Found")
	}
	return c.JSON(http.StatusOK, dev)
}

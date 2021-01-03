package routes

import (
	"net/http"
	"strconv"

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
	return c.JSON(http.StatusOK, a.Cluster.GetConfig())
}

func (a APIService) getDevices(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster.Devices)
}

func (a APIService) getDevice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusTeapot, "Fuck")
	}
	micros := a.Cluster.GetDevices()
	micro, ok := micros[id]
	if !ok {
		return c.JSON(http.StatusTeapot, "Fuck")
	}
	return c.JSON(http.StatusOK, micro)
}

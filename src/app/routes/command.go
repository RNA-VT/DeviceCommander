package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a APIService) addCommandRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/cmd", a.getCommands)
	api.GET("/component/:id/fire/:duration", a.fireSolenoid)
	api.GET("/component/:id/open", a.openSolenoid)
	api.GET("/component/:id/close", a.closeSolenoid)
	api.GET("/component/:id/enable", a.enableSolenoid)
	api.GET("/component/:id/disable", a.disableSolenoid)
}

func (a APIService) getCommands(c echo.Context) error {
	//TODO: De-sass this endpoint
	return c.JSON(http.StatusMethodNotAllowed, "You cannot control me")
}

func (a APIService) openSolenoid(c echo.Context) error {
	component, err := a.Cluster.GetComponent(c.Param("id"))
	if err != nil {
		return err
	}
	component.Open()
	return c.JSON(http.StatusOK, component.State())
}

func (a APIService) fireSolenoid(c echo.Context) error {
	components := a.Cluster.GetComponents()
	componentID := c.Param("id")
	duration, err := strconv.Atoi(c.Param("duration"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Duration invalid.")
	}

	component, ok := components[componentID]
	if !ok {
		return c.JSON(http.StatusBadRequest, "Component Not Found.")
	}
	component.OpenFor(duration)
	return c.JSON(http.StatusOK, a.Cluster.SlaveMicrocontrolers)
}
func (a APIService) closeSolenoid(c echo.Context) error {
	component, err := a.Cluster.GetComponent(c.Param("id"))
	if err != nil {
		return err
	}
	component.Close(0)
	return c.JSON(http.StatusOK, component.State())
}
func (a APIService) disableSolenoid(c echo.Context) error {
	component, err := a.Cluster.GetComponent(c.Param("id"))
	if err != nil {
		return err
	}
	component.Disable()
	return c.JSON(http.StatusOK, component.State())
}

func (a APIService) enableSolenoid(c echo.Context) error {
	component, err := a.Cluster.GetComponent(c.Param("id"))
	if err != nil {
		return err
	}
	//I Guess Always
	component.Enable(true)
	return c.JSON(http.StatusOK, component.State())
}

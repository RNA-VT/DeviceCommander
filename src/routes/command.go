package routes

import (
	"encoding/json"
	"firecontroller/cluster"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a *APIService) addCommandRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/cmd", a.getCommands)
	api.GET("component/:id/cmd", a.processCommand)
}

func (a *APIService) processCommand(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	body, err := c.Request().GetBody()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to get warning message body")
	}
	decoder := json.NewDecoder(body)
	var msg cluster.CommandMessage
	err = decoder.Decode(&msg)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Malformatted Command")
	}

	sol, err := a.Cluster.Me.GetSolenoid(id)
	if err == nil {
		//It's a Solenoid!
		sol.Command(msg.Command)
	}

	igniter, err := a.Cluster.Me.GetIgniter(id)
	if err == nil {
		//It's an Igniter!
		igniter.Command(msg.Command)
	}
	return c.JSON(http.StatusBadRequest, "Not Found")

}

func (a APIService) getCommands(c echo.Context) error {
	//TODO: De-sass this endpoint
	return c.JSON(http.StatusMethodNotAllowed, "You cannot control me")
}

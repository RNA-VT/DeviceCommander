package routes

import (
	"encoding/json"
	"firecontroller/cluster"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a APIService) addCommandRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/cmd", a.getCommands)
	api.POST("/component/:id/cmd", a.processCommand)
}

func (a APIService) processCommand(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	body := c.Request().Body

	decoder := json.NewDecoder(body)

	var msg cluster.CommandMessage
	err = decoder.Decode(&msg)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Malformatted Command")
	}
	fmt.Println(msg)

	if msg.ComponentType == "solenoid" {
		sol, err := a.Cluster.Me.GetSolenoid(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Solenoid does not exist")
		}
		//It's a Solenoid!
		sol.Command(msg.Command)
		return c.JSON(http.StatusOK, "Solenoid has been commanded")
	} else if msg.ComponentType == "igniter" {
		igniter, err := a.Cluster.Me.GetIgniter(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Igniter does not exist")
		}

		//It's an Igniter!
		igniter.Command(msg.Command)
		return c.JSON(http.StatusOK, "Igniter has been commanded")
	}

	return c.JSON(http.StatusOK, "No components have been commanded")
}

func (a APIService) getCommands(c echo.Context) error {
	//TODO: De-sass this endpoint
	return c.JSON(http.StatusMethodNotAllowed, "You cannot control me")
}

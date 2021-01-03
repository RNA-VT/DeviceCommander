package routes

import (
	"encoding/json"
	"firecontroller/cluster"
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

	body := c.Request().Body
	decoder := json.NewDecoder(body)
	var msg cluster.JoinNetworkMessage
	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}

	a.Cluster.AddDevice(msg.ImNewHere)

	return c.JSON(http.StatusOK, "Registered.")
}

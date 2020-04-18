package routes

import (
	"encoding/json"
	"firecontroller/cluster"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a *APIService) addErrorRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/uhoh", a.getErrors)
	api.POST("/uhoh", a.handleConcern)
}

func (a APIService) getErrors(c echo.Context) error {
	// TODO De-sass this endpoint
	return c.JSON(http.StatusMethodNotAllowed, "Nothing has ever been wrong with anything.")
}

func (a APIService) handleConcern(c echo.Context) error {
	body, err := c.Request().GetBody()
	if err != nil {
		log.Println("Failed to get warning message body")
	}
	decoder := json.NewDecoder(body)
	var msg cluster.PeerErrorMessage
	err = decoder.Decode(&msg)

	a.Cluster.ReceiveError(msg)
	return c.JSON(200, "")
}

package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/rna-vt/devicecommander/pkg/rest/controllers"
)

type OpsRouter struct {
	OpsController controllers.OpsController
}

func (r OpsRouter) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/v1/ops")

	api.GET("/run-arp-scan", r.OpsController.RunARPScan)
}

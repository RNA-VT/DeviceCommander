package routes

import (
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/src/rest/controllers"
)

type DeviceRouter struct {
	DeviceController controllers.DeviceController
}

func (r DeviceRouter) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/v1/device")
	api.POST("/", r.DeviceController.Create)
	api.GET("/", r.DeviceController.GetAll)
	api.GET("/device/:id", r.DeviceController.GetDevice)
}

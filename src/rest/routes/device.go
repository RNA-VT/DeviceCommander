package routes

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/rna-vt/devicecommander/src/rest/controllers"
)

type DeviceRouter struct {
	DeviceController controllers.DeviceController
}

func (r DeviceRouter) RegisterRoutes(e *echo.Echo) {
	log.Println("register device routes")
	api := e.Group("/v1")

	api.POST("/", r.DeviceController.Create)

	// swagger:route GET /v1/device getAllDevices
	//
	// Lists all Devices
	//
	// This will show all devices stored in the DB.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Deprecated: false
	//
	//     Responses:
	//       default: genericError
	//       200: getAllDeviceResponse
	//       422: validationError
	api.GET("/device", r.DeviceController.GetAll)

	api.GET("/device/:id", r.DeviceController.GetDevice)
	api.DELETE("/device/:id", r.DeviceController.Delete)
}

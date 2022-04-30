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
	//       200: someResponse
	//       422: validationError
	api.GET("/", r.DeviceController.GetAll)

	api.GET("/:id", r.DeviceController.GetDevice)
	api.DELETE("/:id", r.DeviceController.Delete)
}

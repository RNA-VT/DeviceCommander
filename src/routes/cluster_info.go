package routes

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/graph/model"
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
	// swagger:route GET /cluster-info clusterInfo
	//
	// Get general data about the cluster.
	//
	// This will only show surface details. Further information
	// can be queried via other routes.
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

	devices, err := a.DeviceRepository.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, devices)
}

func (a APIService) getDevices(c echo.Context) error {
	devices, err := a.DeviceRepository.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, devices)
}

func (a *APIService) getDevice(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	tmpDev := model.Device{ID: id}
	device, err := a.DeviceRepository.Get(tmpDev)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &device)
}

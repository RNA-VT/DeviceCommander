package app

import (
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/postgres"
	"github.com/rna-vt/devicecommander/src/routes"
)

// The Application encapsulates the required state for the running software
type Application struct {
	Cluster         cluster.Cluster
	Echo            *echo.Echo
	Hostname        string
	DeviceService   postgres.DeviceCRUDService
	EndpointService postgres.EndpointCRUDService
}

// SystemInfo returns a stringified version of this api
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

func (a *Application) Start() {
	api := routes.NewAPIService(&a.Cluster, a.DeviceService, a.EndpointService)

	a.Cluster.Start()

	api.ConfigureRoutes(a.Hostname, a.Echo)
}

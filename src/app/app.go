package app

import (
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/cluster"
	"github.com/rna-vt/devicecommander/postgres"
	"github.com/rna-vt/devicecommander/routes"
)

// The Application encapsulates the required state for the running software
type Application struct {
	Cluster       cluster.Cluster
	Echo          *echo.Echo
	Hostname      string
	DeviceService postgres.DeviceCRUDService
}

// SystemInfo returns a stringified version of this api
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

func (a *Application) Start() {
	api := routes.APIService{
		Cluster:       &a.Cluster,
		DeviceService: a.DeviceService,
	}

	a.Cluster.Start()

	routes.ConfigureRoutes(a.Hostname, a.Echo, &api, a.DeviceService)
}

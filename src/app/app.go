package app

import (
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/endpoint"
	"github.com/rna-vt/devicecommander/src/routes"
)

// The Application encapsulates the required state for the running software.
type Application struct {
	Cluster            cluster.Cluster
	Echo               *echo.Echo
	Hostname           string
	DeviceRepository   device.Repository
	EndpointRepository endpoint.Repository
}

// SystemInfo returns a stringified version of this api.
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name() + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

func (a *Application) Start() {
	api := routes.NewAPIService(&a.Cluster, a.DeviceRepository, a.EndpointRepository)

	a.Cluster.Start()

	api.ConfigureRoutes(a.Hostname, a.Echo)
}

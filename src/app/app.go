package app

import (
	"firecontroller/cluster"

	"github.com/labstack/echo"
)

// The Application encapsulates the required state for the running software
type Application struct {
	Cluster cluster.Cluster
	Echo    *echo.Echo
}

//SystemInfo returns a stringified version of this api
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

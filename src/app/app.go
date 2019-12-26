package app

import (
	"firecontroller/nodecluster"

	"github.com/labstack/echo"
)

// The Application encapsulates the required state for the running software
type Application struct {
	Cluster nodecluster.Cluster
	Echo    *echo.Echo
}

//SystemInfo returns a stringified version of this api
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

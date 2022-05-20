package app

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/rest/routes"
)

// The Application encapsulates the required state for the running software.
type Application struct {
	Cluster  cluster.Cluster
	Echo     *echo.Echo
	Hostname string
	Router   routes.Router
}

// SystemInfo returns a stringified version of this api.
func (a *Application) SystemInfo() string {
	return "Cluster: " + a.Cluster.Name() + "\nEcho Server: " + a.Echo.Server.TLSConfig.ServerName
}

func (a *Application) Start() {
	a.startListening()
	a.startMaintainingCluster()
}

func (a *Application) startListening() {

	a.Router.RegisterRoutes(a.Echo)
	log.Info("Configured routes listening on " + a.Hostname)

	log.Println("*****************************************************")
	log.Println("~Rejoice~ The Device Commander Lives Again! ~Rejoice~")
	log.Println("*****************************************************")

	// Start server
	a.Echo.Logger.Fatal(a.Echo.Start(a.Hostname))
}

func (a *Application) startMaintainingCluster() {
	a.Cluster.Start()
}

package app

import (
	echo "github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/rest/routes"
)

// The Application encapsulates the required state for the running software.
type Application struct {
	Echo     *echo.Echo
	Hostname string
	Routers  []routes.Router
}

// SystemInfo returns a stringified version of this api.
func (a *Application) SystemInfo() string {
	return "Echo Server: " + a.Echo.Server.TLSConfig.ServerName
}

func (a *Application) Start() {
	a.startListening()
}

func (a *Application) startListening() {
	for _, router := range a.Routers {
		router.RegisterRoutes(a.Echo)
	}
	log.Info("Configured routes listening on " + a.Hostname)

	log.Println("*****************************************************")
	log.Println("~Rejoice~ The Device Commander Lives Again! ~Rejoice~")
	log.Println("*****************************************************")

	// Start server
	a.Echo.Logger.Fatal(a.Echo.Start(a.Hostname))
}

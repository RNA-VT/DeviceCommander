package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/cluster"
	"github.com/rna-vt/devicecommander/postgres"
)

// APIService -
type APIService struct {
	Cluster       *cluster.Cluster
	DeviceService *postgres.DeviceService
}

// ConfigureRoutes will use Echo to start listening on the appropriate paths
func ConfigureRoutes(listenURL string, e *echo.Echo, API APIService, deviceService *postgres.DeviceService) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	frontendRoot := "../frontend/build/"
	if viper.GetString("ENV") == "production" {
		frontendRoot = "/src/build/"
	}
	// Routes
	e.Static("/static", frontendRoot+"static")
	e.File("/*", frontendRoot+"index.html")
	e.GET("/v1", API.defaultGet)

	API.addRegistrationRoutes(e)
	API.addInfoRoutes(e)
	API.addGraphQLRoutes(e, deviceService)

	log.WithFields(log.Fields{
		"module": "routes",
	}).Info("Configured routes listening on " + listenURL)

	log.Println("*****************************************************")
	log.Println("~Rejoice~ The Device Commander Lives Again! ~Rejoice~")
	log.Println("*****************************************************")

	// Start server
	e.Logger.Fatal(e.Start(listenURL))
}

func (a APIService) defaultGet(c echo.Context) error {
	log.Println("Someone is touching me", a.Cluster)
	return c.String(http.StatusOK, "Help Me! I'm trapped in the Server! You're the only one receiving this message.")
}

package routes

import (
	"firecontroller/cluster"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

//APIService -
type APIService struct {
	Cluster *cluster.Cluster
}

const apiVersion = "1"

// ConfigureRoutes will use Echo to start listening on the appropriate paths
func ConfigureRoutes(listenURL string, e *echo.Echo, API APIService) {

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
	API.addErrorRoutes(e)
	API.addCommandRoutes(e)
	API.addManageRoutes(e)

	log.Println("Configure routes listening on " + listenURL)

	log.Println("***************************************")
	log.Println("~Rejoice~ GoFire Lives Again! ~Rejoice~")
	log.Println("***************************************")

	// Start server
	e.Logger.Fatal(e.Start(listenURL))
}

func (a APIService) defaultGet(c echo.Context) error {
	log.Println("Someone is touching me", a.Cluster.GetConfig())
	return c.String(http.StatusOK, "Help Me! I'm trapped in the Server! You're the only one receiving this message.")
}

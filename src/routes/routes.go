package routes

import (
	"devicecommander/cluster"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// Routes
	addBaseRoutes(e, API)

	API.addRegistrationRoutes(e)
	API.addInfoRoutes(e)
	API.addManageRoutes(e)

	log.Println("Configure routes listening on " + listenURL)

	log.Println("*****************************************************")
	log.Println("~Rejoice~ The Device Commander Lives Again! ~Rejoice~")
	log.Println("*****************************************************")

	// Start server
	e.Logger.Fatal(e.Start(listenURL))
}

func addBaseRoutes(e *echo.Echo, API APIService) {
	e.Static("static", "../frontend/build/static")
	e.File("*", "../frontend/build/index.html")
	api := e.Group("/v1")
	api.GET("/", API.defaultGet)
}

func (a APIService) defaultGet(c echo.Context) error {
	log.Println("Someone is touching me", a.Cluster)
	return c.String(http.StatusOK, "Help Me! I'm trapped in the Server! You're the only one receiving this message.")
}

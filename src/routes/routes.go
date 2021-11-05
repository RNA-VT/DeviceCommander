package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/postgres"
)

// APIService -
type APIService struct {
	Cluster         *cluster.Cluster
	DeviceService   postgres.DeviceCRUDService
	EndpointService postgres.EndpointCRUDService
	logger          *log.Entry
}

func NewAPIService(cluster *cluster.Cluster, deviceService postgres.DeviceCRUDService, endpointService postgres.EndpointCRUDService) *APIService {
	api := APIService{
		Cluster:         cluster,
		DeviceService:   deviceService,
		EndpointService: endpointService,
		logger:          log.WithFields(log.Fields{"module": "routes"}),
	}

	return &api
}

// ConfigureRoutes will use Echo to start listening on the appropriate paths
func (api APIService) ConfigureRoutes(listenURL string, e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
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
	e.GET("/v1", api.defaultGet)

	api.addRegistrationRoutes(e)
	api.addInfoRoutes(e)
	api.addGraphQLRoutes(e, api.DeviceService, api.EndpointService)

	api.logger.Info("Configured routes listening on " + listenURL)

	api.logger.Println("*****************************************************")
	api.logger.Println("~Rejoice~ The Device Commander Lives Again! ~Rejoice~")
	api.logger.Println("*****************************************************")

	// Start server
	e.Logger.Fatal(e.Start(listenURL))
}

func (api APIService) defaultGet(c echo.Context) error {
	log.Println("Someone is touching me", api.Cluster)
	return c.String(http.StatusOK, "Help Me! I'm trapped in the Server! You're the only one receiving this message.")
}

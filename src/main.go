package main

/* Al useful imports */
import (
	"fmt"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/app"
	"github.com/rna-vt/devicecommander/cluster"
	"github.com/rna-vt/devicecommander/postgres"
	"github.com/rna-vt/devicecommander/routes"
	"github.com/rna-vt/devicecommander/utilities"
)

/* The entry point for our System */
func main() {
	/* Load Defaut Env Var Config */
	utilities.ConfigureEnvironment()

	dbConfig := postgres.DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}
	deviceService, err := postgres.NewDeviceService(dbConfig)
	if err != nil {
		log.Error(err)
		return
	}

	endpointService, err := postgres.NewEndpointService(dbConfig)
	if err != nil {
		log.Error(err)
		return
	}

	app := app.Application{
		Cluster: cluster.Cluster{
			Name:          viper.GetString("CLUSTER_NAME"),
			DeviceService: deviceService,
		},
		Echo:            echo.New(),
		Hostname:        fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
		DeviceService:   deviceService,
		EndpointService: endpointService,
	}
	var API routes.APIService

	API.Cluster = &app.Cluster
	app.Start()
}

package main

/* Al useful imports */
import (
	"fmt"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/app"
	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
	"github.com/rna-vt/devicecommander/src/routes"
	"github.com/rna-vt/devicecommander/src/utilities"
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
	deviceRepository, err := postgres.NewDeviceRepository(dbConfig)
	if err != nil {
		log.Error(err)
		return
	}

	deviceClient := device.NewHTTPDeviceClient()

	endpointRepository, err := postgres.NewEndpointRepository(dbConfig)
	if err != nil {
		log.Error(err)
		return
	}

	app := app.Application{
		Cluster:            cluster.NewCluster(viper.GetString("CLUSTER_NAME"), deviceRepository, deviceClient),
		Echo:               echo.New(),
		Hostname:           fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
		DeviceRepository:   deviceRepository,
		EndpointRepository: endpointRepository,
	}
	var API routes.APIService

	API.Cluster = &app.Cluster
	app.Start()
}

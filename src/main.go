package main

/* Al useful imports */
import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/app"
	"github.com/rna-vt/devicecommander/cluster"
	"github.com/rna-vt/devicecommander/postgres"
	"github.com/rna-vt/devicecommander/routes"
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()

	host := viper.GetString("HOST")
	port := viper.GetString("PORT")

	fullHostname := host + ":" + port

	dbConfig := postgres.DbConfig{
		Name:     "postgres",
		Host:     "0.0.0.0",
		Port:     "5432",
		UserName: "postgres",
		Password: "changeme",
	}
	deviceService, err := postgres.NewDeviceService(dbConfig)
	if err != nil {
		log.Error(err)
		return
	}

	app := app.Application{
		Cluster: cluster.Cluster{
			Name: viper.GetString("CLUSTER_NAME"),
		},
		Echo:          echo.New(),
		Hostname:      fullHostname,
		DeviceService: deviceService,
	}
	var API routes.APIService

	API.Cluster = &app.Cluster

	app.Cluster.Start()
	routes.ConfigureRoutes(fullHostname, app.Echo, API, &app.DeviceService)
}

func configureEnvironment() {
	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s ", err))
	// }

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	viper.AutomaticEnv()

	viper.SetDefault("ENV", "local") // local or production only
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", 8001)
	viper.SetDefault("CLUSTER_NAME", "Flaming Hot Fleet Directory")
	viper.SetDefault("IP_ADDRESS_ROOT", "192.16.1.")
	viper.SetDefault("DISCOVERY_PERIOD", 30)
	viper.SetDefault("HEALTH_CHECK_PERIOD", 60)
}

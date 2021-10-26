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
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()

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

	app := app.Application{
		Cluster: cluster.Cluster{
			Name:          viper.GetString("CLUSTER_NAME"),
			DeviceService: deviceService,
		},
		Echo:          echo.New(),
		Hostname:      fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
		DeviceService: deviceService,
	}
	var API routes.APIService

	API.Cluster = &app.Cluster
	app.Start()
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
	viper.SetDefault("IP_ADDRESS_ROOT", "192.168.1.")
	viper.SetDefault("DISCOVERY_PERIOD", 30)
	viper.SetDefault("HEALTH_CHECK_PERIOD", 60)

	viper.SetDefault("POSTGRES_NAME", "postgres")
	viper.SetDefault("POSTGRES_HOST", "0.0.0.0")
	viper.SetDefault("POSTGRES_PORT", 5432)
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "changeme")
}

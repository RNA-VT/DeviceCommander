package main

/* Al useful imports */
import (
	"devicecommander/app"
	"devicecommander/cluster"
	"devicecommander/routes"
	"fmt"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()

	host := viper.GetString("HOST")
	port := viper.GetString("PORT")

	fullHostname := host + ":" + port

	app := app.Application{
		Cluster: cluster.Cluster{
			Name: viper.GetString("CLUSTER_NAME"),
		},
		Echo: echo.New(),
	}
	var API routes.APIService

	API.Cluster = &app.Cluster

	app.Cluster.Start()
	routes.ConfigureRoutes(fullHostname, app.Echo, API)
}

func configureEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	viper.AutomaticEnv()

	viper.SetDefault("ENV", "local") // local or production only
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", 8001)
	viper.SetDefault("CLUSTER_NAME", "Flaming Hot Fleet Directory")
	viper.SetDefault("IP_ADDRESS_ROOT", "192.168.1.")
	viper.SetDefault("DISCOVERY_PERIOD", 30)
	viper.SetDefault("HEALTH_CHECK_PERIOD", 5)
}

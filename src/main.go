package main

/* Al useful imports */
import (
	"firecontroller/app"
	"firecontroller/cluster"
	"firecontroller/routes"
	"fmt"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()

	//Pick Listening Port
	port := "8001"
	host := "device1.mindshark.io"

	if viper.GetBool("GOFIRE_MASTER") {
		host = viper.GetString("GOFIRE_MASTER_HOST")
		port = viper.GetString("GOFIRE_MASTER_PORT")
	} else {
		host = viper.GetString("GOFIRE_HOST")
		port = viper.GetString("GOFIRE_PORT")
	}

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
	viper.SetDefault("GOFIRE_MASTER", false)
	viper.SetDefault("GOFIRE_HOST", "127.0.0.1")
	viper.SetDefault("GOFIRE_PORT", 8001)
	viper.SetDefault("GOFIRE_MASTER_PORT", 8000)
	viper.SetDefault("GOFIRE_MASTER_HOST", "127.0.0.1")
	viper.SetDefault("GOFIRE_MOCK_GPIO", true)
	viper.SetDefault("CLUSTER_NAME", "MasterOfHot")
	viper.SetDefault("MICROCONTORLLER_LIMIT", 255)
}

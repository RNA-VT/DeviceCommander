package mock

import (
	"github.com/labstack/echo"
	"github.com/rna-vt/devicecommander/app"
	"github.com/rna-vt/devicecommander/cluster"
	"github.com/spf13/viper"
)

func NewMockApp() app.Application {
	app := app.Application{
		Cluster: cluster.Cluster{
			Name:          viper.GetString("CLUSTER_NAME"),
			DeviceService: deviceService,
		},
		Echo:          echo.New(),
		Hostname:      "0.0.0.0:8001",
		DeviceService: deviceService,
	}

	return app
}

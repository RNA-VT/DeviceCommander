package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/app"
	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/src/postgres/device"
	"github.com/rna-vt/devicecommander/src/rest/controllers"
	"github.com/rna-vt/devicecommander/src/rest/routes"
)

func init() {
	RootCmd.AddCommand(NewServerCommand())
}

func NewServerCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "server",
		Short: "Run a device-commander server instance.",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			dbConfig := postgres.GetDBConfigFromEnv()
			deviceRepository, err := postgresDevice.NewRepository(dbConfig)
			if err != nil {
				log.Error(err)
				return err
			}

			echoInstance := echo.New()

			router := routes.BaseRouter{
				DeviceRouter: routes.DeviceRouter{
					DeviceController: controllers.DeviceController{
						Repository: deviceRepository,
					},
				},
			}

			app := app.Application{
				Cluster: cluster.NewDeviceCluster(
					viper.GetString("CLUSTER_NAME"),
					deviceRepository, device.NewHTTPDeviceClient(),
				),
				Echo:     echoInstance,
				Hostname: fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
				Router:   router,
			}

			app.Start()

			return nil
		},
	}

	return &command
}

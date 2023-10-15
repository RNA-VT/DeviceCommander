package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/pkg/app"
	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/device/registration"
	"github.com/rna-vt/devicecommander/pkg/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/pkg/postgres/device"
	"github.com/rna-vt/devicecommander/pkg/rest/controllers"
	"github.com/rna-vt/devicecommander/pkg/rest/routes"
	"github.com/rna-vt/devicecommander/pkg/scanner"
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

			deviceClient := device.NewHTTPDeviceClient()

			echoInstance := echo.New()

			app := app.Application{
				Echo:     echoInstance,
				Hostname: fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
				Routers:  getRouters(deviceClient, deviceRepository),
			}

			app.Start()

			return nil
		},
	}

	return &command
}

func getRouters(deviceClient device.HTTPDeviceClient, deviceRepository device.Repository) []routes.Router {
	return []routes.Router{
		routes.BaseRouter{
			DeviceRouter: routes.DeviceRouter{
				DeviceController: controllers.DeviceController{
					Repository: deviceRepository,
				},
			},
		},
		routes.OpsRouter{
			OpsController: controllers.OpsController{
				DeviceScanner: scanner.NewDeviceScanner(
					deviceClient,
				),
				DeviceRegistrar: registration.NewDeviceRegistrar(
					deviceClient,
					deviceRepository,
				),
			},
		},
	}
}

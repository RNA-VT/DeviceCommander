package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/app"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/device/registration"
	"github.com/rna-vt/devicecommander/src/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/src/postgres/device"
	"github.com/rna-vt/devicecommander/src/rest/controllers"
	"github.com/rna-vt/devicecommander/src/rest/routes"
	"github.com/rna-vt/devicecommander/src/scanner"
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

			baseRouter := routes.BaseRouter{
				DeviceRouter: routes.DeviceRouter{
					DeviceController: controllers.DeviceController{
						Repository: deviceRepository,
					},
				},
			}

			opsRouter := routes.OpsRouter{
				OpsController: controllers.OpsController{
					DeviceScanner: scanner.NewDeviceScanner(
						deviceClient,
					),
					DeviceRegistrar: registration.NewDeviceRegistrar(
						deviceClient,
						deviceRepository,
					),
				},
			}

			app := app.Application{
				Echo:     echoInstance,
				Hostname: fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
				Routers: []routes.Router{
					baseRouter,
					opsRouter,
				},
			}

			app.Start()

			return nil
		},
	}

	return &command
}

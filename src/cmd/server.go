package cmd

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/src/app"
	"github.com/rna-vt/devicecommander/src/cluster"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/postgres"
	postgresDevice "github.com/rna-vt/devicecommander/src/postgres/device"
	postgresEndpoint "github.com/rna-vt/devicecommander/src/postgres/endpoint"
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

			endpointRepository, err := postgresEndpoint.NewRepository(dbConfig)
			if err != nil {
				log.Error(err)
				return err
			}

			app := app.Application{
				Cluster: cluster.NewDeviceCluster(
					viper.GetString("CLUSTER_NAME"),
					deviceRepository, device.NewHTTPDeviceClient(),
				),
				Echo:               echo.New(),
				Hostname:           fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
				DeviceRepository:   deviceRepository,
				EndpointRepository: endpointRepository,
			}

			app.Start()

			return nil
		},
	}

	return &command
}

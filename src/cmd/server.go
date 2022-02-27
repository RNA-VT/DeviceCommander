package cmd

import (
	"github.com/spf13/cobra"

	"github.com/rna-vt/devicecommander/src/rpc"
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
			// dbConfig := postgres.GetDBConfigFromEnv()
			// deviceRepository, err := postgresDevice.NewRepository(dbConfig)
			// if err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			// endpointRepository, err := postgresEndpoint.NewRepository(dbConfig)
			// if err != nil {
			// 	log.Error(err)
			// 	return err
			// }

			// app := app.Application{
			// 	Cluster: cluster.NewDeviceCluster(
			// 		viper.GetString("CLUSTER_NAME"),
			// 		deviceRepository, device.NewHTTPDeviceClient(),
			// 	),
			// 	Echo:               echo.New(),
			// 	Hostname:           fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT")),
			// 	DeviceRepository:   deviceRepository,
			// 	EndpointRepository: endpointRepository,
			// }

			// app.Start()

			rpcServer := rpc.NewDCServer()
			rpcServer = rpcServer.Listen()

			return nil
		},
	}

	return &command
}

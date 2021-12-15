package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/rna-vt/devicecommander/src/postgres"
)

func init() {
	RootCmd.AddCommand(NewMigrateDBCommand())
}

func NewMigrateDBCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "migrate-db",
		Short: "Run all migrations on DB.",
		Long: `Run the GORM postgres AutoMigrations for all tables.
This command can be used to initialize a database or attempt
to update an existing database. Eventually a more complex 
migration tool could be useful.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbConfig := postgres.GetDBConfigFromEnv()

			db, err := postgres.GetDBConnection(dbConfig)
			if err != nil {
				log.Fatal("connecting to the DB should not throw an error", err)
			}

			err = postgres.RunMigration(db)
			if err != nil {
				log.Fatal("connecting to the DB should not throw an error", err)
			}

			log.Info("DB migrations have completed successfully...")

			return nil
		},
	}

	return &command
}

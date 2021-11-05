package postgres

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/rna-vt/devicecommander/src/graph/model"
)

// DBConfig encapsulates the information rquired for connecting to a database.
type DBConfig struct {
	Name     string
	Host     string
	Port     string
	UserName string
	Password string
}

// GetDBConfigFromEnv loads the required DBConfig from env_vars.
func GetDBConfigFromEnv() DBConfig {
	return DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}
}

// RunMigration makes sure each of the important models are fully migrated.
func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Device{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Endpoint{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Parameter{})
	if err != nil {
		return err
	}
	return nil
}

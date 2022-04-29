package postgres

import (
	"fmt"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/device/parameter"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig encapsulates the information required for connecting to a database.
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

func GetDBConnection(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.UserName,
		config.Password,
		config.Name,
		config.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return db, err
	}

	return db, nil
}

// RunMigration makes sure each of the important models are fully migrated.
func RunMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(&device.Device{}, &device.Endpoint{}, &parameter.Parameter{}); err != nil {
		return err
	}

	return nil
}

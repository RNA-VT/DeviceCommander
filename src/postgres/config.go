package postgres

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/rna-vt/devicecommander/graph/model"
)

type DBConfig struct {
	Name     string
	Host     string
	Port     string
	UserName string
	Password string
}

func GetDBConfigFromEnv() DBConfig {
	return DBConfig{
		Name:     viper.GetString("POSTGRES_NAME"),
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		UserName: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
	}
}

func getPostgresLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "postgres"})
}

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

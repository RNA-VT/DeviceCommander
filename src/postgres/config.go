package postgres

import (
	log "github.com/sirupsen/logrus"
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

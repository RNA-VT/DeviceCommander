package postgres

import (
	log "github.com/sirupsen/logrus"
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
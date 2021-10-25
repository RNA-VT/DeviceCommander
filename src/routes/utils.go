package routes

import (
	log "github.com/sirupsen/logrus"
)

func getRouteLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "routes"})
}

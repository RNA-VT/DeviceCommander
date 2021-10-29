package cluster

import (
	log "github.com/sirupsen/logrus"
)

func getClusterLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "cluster"})
}

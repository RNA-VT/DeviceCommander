package memory

import (
	log "github.com/sirupsen/logrus"
)

func getMemoryLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "memory"})
}

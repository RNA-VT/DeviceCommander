package scanner

import (
	log "github.com/sirupsen/logrus"
)

func getScannerLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "scanner"})
}

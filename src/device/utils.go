package device

import (
	log "github.com/sirupsen/logrus"
)

func getDeviceLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "device"})
}

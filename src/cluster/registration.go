package cluster

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/scanner"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

// DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	logger := getRegistrationLogger()
	logger.Info("Device discovery")

	deviceList, err := scanner.ScanNetwork(viper.Get("IP_ADDRESS_ROOT").(string))
	if err != nil {
		logger.Error(err)
		return
	}

	if len(deviceList) > 0 {
		c.AddDevices(deviceList)
	}
}

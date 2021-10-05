package cluster

import (
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/scanner"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

// DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	logger := getRegistrationLogger()
	logger.Info("Device discovery")

	ipResults, err := scanner.GetLocalAddresses()
	if err != nil {
		logger.Error(err)
		return
	}

	deviceList, err := scanner.ScanIPs(ipResults.IPv4Addresses)
	if err != nil {
		logger.Error(err)
		return
	}

	if len(deviceList) > 0 {
		// TODO: Bulk insert
		for _, d := range deviceList {
			c.DeviceService.Create(d)
		}
	}
}

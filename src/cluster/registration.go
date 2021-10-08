package cluster

import (
	"github.com/rna-vt/devicecommander/scanner"
	log "github.com/sirupsen/logrus"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

// DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	logger := getRegistrationLogger()
	logger.Info("Device discovery")

	arpScanner := scanner.ArpScanner{
		LoopDelay:     60,
		DeviceService: &c.DeviceService,
	}

	arpScanner.Start()

	// ipResults, err := scanner.GetLocalAddresses()
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// deviceList, err := scanner.ScanIPs(ipResults.IPv4Addresses)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// if len(deviceList) > 0 {
	// 	// TODO: Bulk insert
	// 	for _, d := range deviceList {
	// 		_, err := c.DeviceService.Create(d)
	// 		if err != nil {
	// 			logger.Error(err)
	// 		}
	// 	}
	// }
}

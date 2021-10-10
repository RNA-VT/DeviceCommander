package cluster

import (
	"fmt"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/scanner"
	log "github.com/sirupsen/logrus"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

// DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	logger := getRegistrationLogger()

	newDevices := make(chan model.NewDevice, 10)
	defer close(newDevices)
	stop := make(chan struct{})
	defer close(stop)
	arpScanner := scanner.ArpScanner{
		LoopDelay:     60,
		NewDeviceChan: newDevices,
		Stop:          stop,
	}

	go arpScanner.Start()

	searchLimit := 10
	for i := 0; i < searchLimit; i++ {
		tmpNewDevice := <-newDevices
		devSearch := model.Device{
			MAC: *tmpNewDevice.Mac,
		}
		results, err := c.DeviceService.Get(devSearch)
		if err != nil {
			logger.Error(err)
		} else {
			if len(results) == 0 {
				completeDevice, err := c.DeviceService.Create(tmpNewDevice)
				if err != nil {
					logger.Error(err)
				} else {
					logger.Debug(fmt.Sprintf("registered mac address [%s] with id [%s]", completeDevice.MAC, completeDevice.ID))
				}
			}
		}
	}
}

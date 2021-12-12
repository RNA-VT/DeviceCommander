package cluster

import (
	"fmt"
	"time"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/scanner"
)

// DeviceDiscovery will start an ArpScanner and use its results to create new
// Devices in the database if they do not already exist.
func (c DeviceCluster) DeviceDiscovery(scanDurationSeconds int) {
	newDevices := make(chan model.NewDevice)
	defer close(newDevices)
	stop := make(chan struct{})
	// defer close(stop)
	arpScanner := scanner.NewArpScanner(newDevices, stop)

	go arpScanner.Start()

	// listen for scanDurationSeconds... then shut it down.
	c.logger.Info(fmt.Sprintf("ARP scanning for %d seconds...", scanDurationSeconds))
	time.AfterFunc(time.Duration(scanDurationSeconds)*time.Second, func() {
		close(stop)
	})

	for {
		select {
		case <-stop:
			c.logger.Debug("Exit NewDevice stream watch")
			return
		case tmpNewDevice := <-newDevices:
			if err := c.HandleDiscoveredDevice(tmpNewDevice); err != nil {
				c.logger.Error(err)
			}
		}
	}
}

func (c DeviceCluster) HandleDiscoveredDevice(newDevice model.NewDevice) error {
	results, err := c.DeviceRepository.Get(model.Device{
		MAC: *newDevice.Mac,
	})
	if err != nil {
		return err
	}

	if len(results) > 0 {
		return nil
	}

	completeDevice, err := c.DeviceRepository.Create(newDevice)
	if err != nil {
		return err
	}

	c.logger.Debug(fmt.Sprintf("registered mac address [%s] with id [%s]", completeDevice.MAC, completeDevice.ID))

	deviceWrapper := device.NewDeviceWrapper(*completeDevice)
	err = deviceWrapper.RunHealthCheck(c.DeviceClient)
	if err != nil {
		return err
	}

	return nil
}

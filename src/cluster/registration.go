package cluster

import (
	"errors"
	"fmt"
	"strconv"
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

// Once a Device is found on the network it needs to get processed into the platform.
// HandleDiscoveredDevice does this with some additional steps. For example:
// 1. does the Device already exist in the DB? (MAC address is the unique identifier in this case).
// 2. immediately check its health.
func (c DeviceCluster) HandleDiscoveredDevice(newDevice model.NewDevice) error {
	results, err := c.DeviceRepository.Get(model.Device{
		MAC: *newDevice.Mac,
	})
	if err != nil {
		c.logger.Print(err)
		return err
	}

	var discoveredDevice *model.Device
	switch len(results) {
	case 0:
		//New Device
		discoveredDevice, err = c.DeviceRepository.Create(newDevice)
		if err != nil {
			c.logger.Error(err)
			return err
		}
	case 1:
		// Known Device
		discoveredDevice = results[0]

		// Activate
		discoveredDevice.Active = true

		// Update device
		var m model.UpdateDevice = model.UpdateDevice{
			Mac:         &discoveredDevice.MAC,
			Name:        newDevice.Name,
			Description: newDevice.Description,
			Host:        &newDevice.Host,
			Port:        &newDevice.Port,
			Active:      &discoveredDevice.Active,
		}
		c.DeviceRepository.Update(m)
	default:
		return errors.New("multiple results returned for 1 mac address")
	}

	c.logger.Debug(fmt.Sprintf("registered mac address [%s] with id [%s] at [%s]:[%s]", discoveredDevice.MAC, discoveredDevice.ID, newDevice.Host, strconv.Itoa(newDevice.Port)))

	// Immediatly run health check
	if err := device.NewDeviceWrapper(*discoveredDevice).RunHealthCheck(c.DeviceClient); err != nil {
		return err
	}

	return nil
}

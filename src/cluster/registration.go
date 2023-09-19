package cluster

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/scanner"
)

// DeviceDiscovery will start an ArpScanner and use its results to create new
// Devices in the database if they do not already exist.
func (c DeviceCluster) DeviceDiscovery(scanDurationSeconds int) {
	newDevices := make(chan device.NewDeviceParams)
	defer close(newDevices)
	scannerStop := make(chan struct{})
	stop := make(chan struct{})
	defer close(scannerStop)
	defer close(stop)
	arpScanner := scanner.NewArpScanner(newDevices, scannerStop)

	go arpScanner.Start()

	// listen for scanDurationSeconds... then shut it down.
	c.logger.Info(fmt.Sprintf("ARP scanning for %d seconds...", scanDurationSeconds))
	time.AfterFunc(time.Duration(scanDurationSeconds)*time.Second, func() {
		c.logger.Info("ARP scan complete... shutting down ARP scanner.")
		scannerStop <- struct{}{}
		stop <- struct{}{}
		// close(stop)
	})

	for {
		select {
		case <-stop:
			c.logger.Debug("Exit NewDevice stream watch")
			return
		case tmpNewDevice := <-newDevices:
			d, err := c.HandleDiscoveredDevice(tmpNewDevice)
			if err != nil {
				c.logger.Warn(err)
			}

			err = d.RunHealthCheck(c.DeviceClient)
			if err != nil {
				c.logger.Warn(err)
			}
		}
	}
}

// Once a Device is found on the network it needs to get processed into the platform.
// HandleDiscoveredDevice does this with some additional steps. For example:
// 1. does the Device already exist in the DB? (MAC address is the unique identifier in this case).
// 2. immediately check its health.
func (c DeviceCluster) HandleDiscoveredDevice(newDevice device.NewDeviceParams) (device.Device, error) {
	results, err := c.DeviceRepository.Get(device.Device{
		MAC: *newDevice.MAC,
	})
	if err != nil {
		return device.Device{}, err
	}

	discoveredDevice := new(device.Device)
	switch len(results) {
	case 0:
		discoveredDevice, err = c.DeviceRepository.Create(newDevice)
		if err != nil {
			return device.Device{}, err
		}

		c.logger.Debugf("registered new device -- mac address [%s] with id [%s] at [%s]:[%s]",
			discoveredDevice.MAC,
			discoveredDevice.ID,
			newDevice.Host,
			strconv.Itoa(newDevice.Port),
		)
	case 1:
		partiallyUpdatedDevice := updateDeviceFromDevice(&newDevice)
		partiallyUpdatedDevice.ID = results[0].ID.String()

		if err := c.DeviceRepository.Update(partiallyUpdatedDevice); err != nil {
			return device.Device{}, err
		}

		results, err = c.DeviceRepository.Get(device.Device{
			ID: discoveredDevice.ID,
		})
		if err != nil {
			return device.Device{}, err
		}
		discoveredDevice = results[0]

		c.logger.Debugf(
			"updated known device -- mac address [%s] with id [%s] at [%s]:[%s]",
			discoveredDevice.MAC,
			discoveredDevice.ID,
			newDevice.Host,
			strconv.Itoa(newDevice.Port),
		)
	default:
		return device.Device{}, errors.New("multiple results returned for 1 mac address")
	}

	return *discoveredDevice, nil
}

// updateDeviceFromDevice builds a device.UpdateDeviceParams from a device.Device.
func updateDeviceFromDevice(d *device.NewDeviceParams) device.UpdateDeviceParams {
	updateDevice := device.UpdateDeviceParams{
		MAC:  d.MAC,
		Host: &d.Host,
		Port: &d.Port,
	}

	return updateDevice
}

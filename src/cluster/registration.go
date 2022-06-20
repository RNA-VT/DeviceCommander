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
			d, err := c.HandleDiscoveredDevice(tmpNewDevice)
			if err != nil {
				c.logger.Error(err)
			}

			err = d.RunHealthCheck(c.DeviceClient)
			if err != nil {
				c.logger.Error(err)
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
		MAC: *newDevice.Mac,
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
	case 1:
		discoveredDevice := device.FromNewDevice(newDevice)
		discoveredDevice.Active = true

		if err := c.DeviceRepository.Update(updateDeviceFromDevice(&discoveredDevice)); err != nil {
			return device.Device{}, err
		}
	default:
		return device.Device{}, errors.New("multiple results returned for 1 mac address")
	}

	c.logger.Debugf("registered mac address [%s] with id [%s] at [%s]:[%s]",
		discoveredDevice.MAC,
		discoveredDevice.ID,
		newDevice.Host,
		strconv.Itoa(newDevice.Port))

	return *discoveredDevice, nil
}

// updateDeviceFromDevice builds a device.UpdateDeviceParams from a device.Device.
func updateDeviceFromDevice(d *device.Device) device.UpdateDeviceParams {
	updateDevice := device.UpdateDeviceParams{
		Mac:  &d.MAC,
		Host: &d.Host,
		Port: &d.Port,
	}

	return updateDevice
}

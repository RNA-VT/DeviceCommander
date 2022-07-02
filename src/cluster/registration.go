package cluster

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

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

	// Add discovered device to platform db
	// Verify that discovered device produces a DeviceCommander compliant api
	// and collect device details
	// Activate verified device
	// Update db
	c.ListenForDevices(stop, newDevices)
}

func (c DeviceCluster) ListenForDevices(stop chan struct{}, newDevices chan device.NewDeviceParams) {
	for {
		select {
		case <-stop:
			c.logger.Debug("Exit NewDevice stream watch")
			return
		case tmpNewDevice := <-newDevices:

			d, err := c.HandleDiscoveredDevice(tmpNewDevice)
			if err != nil {
				c.logger.WithFields(logrus.Fields{
					"host": d.Host,
					"port": d.Port,
					"mac":  d.MAC,
				}).Error("discovered device init failed: failed to register new device with platform")
				c.logger.Error(err)
			}

			d, verified := c.VerifyDeviceAPI(d)
			if verified {
				d.Activate()

				if err := c.DeviceRepository.Update(updateDeviceFromDevice(&d)); err != nil {
					c.logger.Error(err)
				}
			}
		}
	}
}

// VerifyDeviceAPI confirms that the device is compliant & returns the device hydrated with return data from the device.
func (c DeviceCluster) VerifyDeviceAPI(d device.Device) (device.Device, bool) {
	// Verify that Health Check endpoint responds and device is healthy
	err := d.RunHealthCheck(c.DeviceClient)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"host": d.Host,
			"port": d.Port,
			"mac":  d.MAC,
		}).Error("device failed api verification: health check failed")
		c.logger.Error(err)
		return device.Device{}, false
	}

	// Get Device Spec
	spec, err := d.RequestSpecification(c.DeviceClient)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"host": d.Host,
			"port": d.Port,
			"mac":  d.MAC,
		}).Error("device failed api verification: failed to request and load device specification")
		c.logger.Error(err)
		return device.Device{}, false
	}

	// Return spec'd device
	return d.LoadFromSpecification(spec), true
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

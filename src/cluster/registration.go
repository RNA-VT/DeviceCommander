package cluster

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/pkg/errors"
// 	"github.com/rna-vt/devicecommander/src/device"
// 	"github.com/rna-vt/devicecommander/src/scanner"
// 	"github.com/rna-vt/devicecommander/src/utils"
// )

// // DeviceDiscovery will start an ArpScanner and use its results to create new
// // Devices in the database if they do not already exist.
// func (c DeviceCluster) DeviceDiscovery(scanDurationSeconds int) {
// 	foundDevices := make(chan scanner.FoundDevice)
// 	defer close(foundDevices)
// 	scannerStop := make(chan struct{})
// 	stop := make(chan struct{})
// 	defer close(scannerStop)
// 	defer close(stop)
// 	arpScanner := scanner.NewArpScanner(foundDevices, scannerStop)

// 	go arpScanner.Start()

// 	// listen for scanDurationSeconds... then shut it down.
// 	c.logger.Info(fmt.Sprintf("ARP scanning for %d seconds...", scanDurationSeconds))
// 	time.AfterFunc(time.Duration(scanDurationSeconds)*time.Second, func() {
// 		c.logger.Info("ARP scan complete... shutting down ARP scanner.")
// 		scannerStop <- struct{}{}
// 		stop <- struct{}{}
// 		// close(stop)
// 	})

// 	for {
// 		select {
// 		case <-stop:
// 			c.logger.Debug("Exit NewDevice stream watch")
// 			return
// 		case tmpNewDevice := <-foundDevices:
// 			d, err := c.HandleDiscoveredDevice(tmpNewDevice)
// 			if err != nil {
// 				c.logger.Warn(err)
// 			}

// 			err = d.RunHealthCheck(c.DeviceClient)
// 			if err != nil {
// 				c.logger.Warn(err)
// 			}
// 		}
// 	}
// }

// // Once a Device is found on the network it needs to get processed into the platform.
// // HandleDiscoveredDevice does this with some additional steps. For example:
// // 1. does the Device already exist in the DB? (MAC address is the unique identifier in this case).
// // 2. immediately check its health.
// func (c DeviceCluster) HandleDiscoveredDevice(foundDevice scanner.FoundDevice) (device.Device, error) {
// 	results, err := c.DeviceRepository.Get(device.Device{
// 		MAC: foundDevice.MAC,
// 	})
// 	if err != nil {
// 		return device.Device{}, err
// 	}

// 	tmpSpec, err := c.DeviceClient.Specification(device.Device{
// 		Host: foundDevice.IP,
// 		Port: foundDevice.Port,
// 	})
// 	if err != nil {
// 		return device.Device{}, errors.Wrap(err, "error getting device specification")
// 	}

// 	c.logger.Debug(utils.PrettyPrintJSON(tmpSpec))

// 	discoveredDevice := &device.Device{}
// 	switch len(results) {
// 	case 0:
// 		discoveredDevice, err = c.DeviceRepository.Create(newDeviceFromFoundDevice(foundDevice))
// 		if err != nil {
// 			return device.Device{}, err
// 		}

// 		c.logger.Debugf("registered new device -- mac address [%s] with id [%s] at [%s]:[%s]",
// 			discoveredDevice.MAC,
// 			discoveredDevice.ID,
// 			foundDevice.IP,
// 			strconv.Itoa(foundDevice.Port),
// 		)

// 	case 1:
// 		discoveredDevice, err = c.handleKnownDevice(foundDevice, *results[0])
// 		if err != nil {
// 			return device.Device{}, err
// 		}

// 	default:
// 		return device.Device{}, errors.New("multiple results returned for 1 mac address")
// 	}

// 	return *discoveredDevice, nil
// }

// func newDeviceFromFoundDevice(d scanner.FoundDevice) device.NewDeviceParams {
// 	return device.NewDeviceParams{
// 		MAC:  &d.MAC,
// 		Host: d.IP,
// 		Port: d.Port,
// 	}
// }

// func (c DeviceCluster) handleKnownDevice(foundDevice scanner.FoundDevice, existingDevice device.Device) (*device.Device, error) {
// 	if existingDevice.MAC != foundDevice.MAC {
// 		return &device.Device{}, errors.New("mac address mismatch")
// 	}

// 	existingDevice.Host = foundDevice.IP
// 	existingDevice.Port = foundDevice.Port

// 	partiallyUpdatedDevice := updateDeviceFromFoundDevice(existingDevice.ID.String(), foundDevice)
// 	partiallyUpdatedDevice.ID = existingDevice.ID.String()

// 	if err := c.DeviceRepository.Update(partiallyUpdatedDevice); err != nil {
// 		return &device.Device{}, err
// 	}

// 	results, err := c.DeviceRepository.Get(device.Device{
// 		ID: existingDevice.ID,
// 	})
// 	if err != nil {
// 		return &device.Device{}, err
// 	}
// 	resultingDevice := results[0]

// 	c.logger.Debugf(
// 		"updated known device -- mac address [%s] with id [%s] at [%s]:[%s]",
// 		resultingDevice.MAC,
// 		resultingDevice.ID,
// 		foundDevice.IP,
// 		strconv.Itoa(foundDevice.Port),
// 	)

// 	return resultingDevice, nil
// }

// // updateDeviceFromDevice builds a device.UpdateDeviceParams from a device.Device.
// func updateDeviceFromFoundDevice(targetID string, d scanner.FoundDevice) device.UpdateDeviceParams {
// 	updateDevice := device.UpdateDeviceParams{
// 		ID:   targetID,
// 		MAC:  &d.MAC,
// 		Host: &d.IP,
// 		Port: &d.Port,
// 	}

// 	return updateDevice
// }

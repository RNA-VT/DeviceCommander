package registration

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/rna-vt/devicecommander/pkg/device"
	"github.com/rna-vt/devicecommander/pkg/scanner"
	"github.com/rna-vt/devicecommander/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type Registrar interface {
	HandleProspects([]scanner.FoundDevice) ([]device.Device, error)
}

type DeviceRegistrar struct {
	DeviceClient     device.Client
	DeviceRepository device.Repository
	logger           *log.Entry
}

func NewDeviceRegistrar(deviceClient device.Client, deviceRepository device.Repository) DeviceRegistrar {
	return DeviceRegistrar{
		DeviceClient:     deviceClient,
		DeviceRepository: deviceRepository,
		logger:           log.WithFields(log.Fields{"module": "device_registrar"}),
	}
}

func (s DeviceRegistrar) HandleProspects(prospects []scanner.FoundDevice) ([]device.Device, error) {
	newDevices := []device.Device{}

	for _, prospect := range prospects {
		newDevice, err := s.handleDiscoveredDevice(prospect)
		if err != nil {
			return []device.Device{}, err
		}
		newDevices = append(newDevices, newDevice)
	}

	return newDevices, nil
}

// Once a Device is found on the network it needs to get processed into the platform.
// HandleDiscoveredDevice does this with some additional steps. For example:
// 1. does the Device already exist in the DB? (MAC address is the unique identifier in this case).
// 2. immediately check its health.
func (s DeviceRegistrar) handleDiscoveredDevice(foundDevice scanner.FoundDevice) (device.Device, error) {
	results, err := s.DeviceRepository.Get(device.Device{
		MAC: foundDevice.MAC,
	})
	if err != nil {
		return device.Device{}, err
	}

	tmpSpec, err := s.DeviceClient.Specification(device.Device{
		Host: foundDevice.IP,
		Port: foundDevice.Port,
	})
	if err != nil {
		return device.Device{}, errors.Wrap(err, "error getting device specification")
	}

	s.logger.Debug(utils.PrettyPrintJSON(tmpSpec))

	discoveredDevice := &device.Device{}
	switch len(results) {
	case 0:
		discoveredDevice, err = s.DeviceRepository.Create(newDeviceFromFoundDevice(foundDevice))
		if err != nil {
			return device.Device{}, err
		}

		s.logger.Debugf("registered new device -- mac address [%s] with id [%s] at [%s]:[%s]",
			discoveredDevice.MAC,
			discoveredDevice.ID,
			foundDevice.IP,
			strconv.Itoa(foundDevice.Port),
		)

	case 1:
		discoveredDevice, err = s.handleKnownDevice(foundDevice, *results[0])
		if err != nil {
			return device.Device{}, err
		}

	default:
		return device.Device{}, errors.New("multiple results returned for 1 mac address")
	}

	return *discoveredDevice, nil
}

func (s DeviceRegistrar) handleKnownDevice(foundDevice scanner.FoundDevice, existingDevice device.Device) (*device.Device, error) {
	if existingDevice.MAC != foundDevice.MAC {
		return &device.Device{}, errors.New("mac address mismatch")
	}

	existingDevice.Host = foundDevice.IP
	existingDevice.Port = foundDevice.Port

	partiallyUpdatedDevice := updateDeviceFromFoundDevice(existingDevice.ID.String(), foundDevice)
	partiallyUpdatedDevice.ID = existingDevice.ID.String()

	if err := s.DeviceRepository.Update(partiallyUpdatedDevice); err != nil {
		return &device.Device{}, err
	}

	results, err := s.DeviceRepository.Get(device.Device{
		ID: existingDevice.ID,
	})
	if err != nil {
		return &device.Device{}, err
	}
	resultingDevice := results[0]

	s.logger.Debugf(
		"updated known device -- mac address [%s] with id [%s] at [%s]:[%s]",
		resultingDevice.MAC,
		resultingDevice.ID,
		foundDevice.IP,
		strconv.Itoa(foundDevice.Port),
	)

	return resultingDevice, nil
}

// updateDeviceFromDevice builds a device.UpdateDeviceParams from a device.Device.
func updateDeviceFromFoundDevice(targetID string, d scanner.FoundDevice) device.UpdateDeviceParams {
	updateDevice := device.UpdateDeviceParams{
		ID:   targetID,
		MAC:  &d.MAC,
		Host: &d.IP,
		Port: &d.Port,
	}

	return updateDevice
}

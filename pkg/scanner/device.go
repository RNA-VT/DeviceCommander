package scanner

import (
	"fmt"
	"time"

	"github.com/rna-vt/devicecommander/pkg/device"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Scanner interface {
	Scan() ([]FoundDevice, error)
}

type FoundDevice struct {
	MAC  string
	IP   string
	Port int
}

type DeviceScanner struct {
	DeviceClient    device.Client
	logger          *log.Entry
	discoveryStop   chan bool
	discoveryPeriod int
}

func NewDeviceScanner(deviceClient device.Client) DeviceScanner {
	return DeviceScanner{
		DeviceClient:    deviceClient,
		logger:          log.WithFields(log.Fields{"module": "scanner"}),
		discoveryStop:   make(chan bool),
		discoveryPeriod: viper.GetInt("DISCOVERY_PERIOD"),
	}
}

// DeviceDiscovery will start an ArpScanner and use its results to create new
// Devices in the database if they do not already exist.
func (s DeviceScanner) Scan() ([]FoundDevice, error) {
	discoveredDeviceChan := make(chan FoundDevice)
	discoveredDevices := []FoundDevice{}
	defer close(discoveredDeviceChan)
	scannerStop := make(chan struct{})
	stop := make(chan struct{})
	defer close(scannerStop)
	defer close(stop)
	arpScanner := NewArpScanner(discoveredDeviceChan, scannerStop)

	go arpScanner.Start()

	// listen for scanDurationSeconds... then shut it down.
	s.logger.Info(fmt.Sprintf("ARP scanning for %d seconds...", s.discoveryPeriod))
	time.AfterFunc(time.Duration(s.discoveryPeriod)*time.Second, func() {
		s.logger.Info("ARP scan complete... shutting down ARP scanner.")
		scannerStop <- struct{}{}
		stop <- struct{}{}
	})

	for {
		select {
		case <-stop:
			s.logger.Debug("Exit device scanner.")
			return discoveredDevices, nil
		case tmpNewDevice := <-discoveredDeviceChan:
			discoveredDevices = append(discoveredDevices, tmpNewDevice)
		}
	}
}

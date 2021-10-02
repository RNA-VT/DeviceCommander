package cluster

import (
	device "devicecommander/device"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Cluster - This object defines an array of Devices
type Cluster struct {
	Name    string
	Devices []device.Device
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c Cluster) PrintClusterInfo() {
	for i := 0; i < len(c.Devices); i++ {
		log.Println("----Device---")
		log.Println(c.Devices[i])
	}
	log.Println()
}

func (c Cluster) GetDeviceByHost(host string) (device.Device, error) {
	for _, dev := range c.Devices {
		if dev.Host == host {
			return dev, nil
		}
	}
	return device.Device{}, errors.New("Failed to find device in existing list.")
}

//GetDevices returns a map of all registered
func (c Cluster) GetDevices() map[string]device.Device {
	devices := make(map[string]device.Device)
	for i := 0; i < len(c.Devices); i++ {
		devices[c.Devices[i].ID] = c.Devices[i]
	}
	return devices
}

//AddDevice attempts to add a device to the cluster and returns the response data.
func (c *Cluster) AddDevice(newDevice device.Device) {
	clusterLogger := getClusterLogger()
	for _, dev := range c.Devices {
		if dev.URL() == newDevice.URL() {
			clusterLogger.Debug("This host & port combination are already registered to this cluster")
			break
		}
	}

	c.Devices = append(c.Devices, newDevice)

	c.PrintClusterInfo()
}

//RemoveDevice -
func (c *Cluster) RemoveDevice(deviceID string) {
	for index, device := range c.Devices {
		if device.ID == deviceID {
			s := c.Devices
			count := len(c.Devices)
			s[count-1], s[index] = s[index], s[count-1]
			c.Devices = s[:len(s)-1]
			return
		}
	}
}

//Start begins the registration and health check goroutines
func (c *Cluster) Start() {
	discoveryPeriod := viper.GetInt("DISCOVERY_PERIOD")
	healthCheckPeriod := viper.GetInt("HEALTH_CHECK_PERIOD")
	clusterLogger := getClusterLogger()

	// Device Discovery
	go func() {
		ticker := time.NewTicker(time.Duration(discoveryPeriod) * time.Second)
		for range ticker.C {
			clusterLogger.Info("Begin Device Discovery... ")
			DeviceDiscovery(c)
		}
	}()

	// Health Check
	go func() {
		ticker := time.NewTicker(time.Duration(healthCheckPeriod) * time.Second)
		for range ticker.C {
			clusterLogger.Info("Begin Health Checks... ")
			for _, device := range c.Devices {
				device.CheckHealth()
			}
		}
	}()
}

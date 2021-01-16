package cluster

import (
	device "devicecommander/device"
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of Devices
type Cluster struct {
	Name    string
	Devices []device.Device
}

//GetDevices returns a map of all registered
func (c Cluster) GetDevices() map[int]device.Device {
	micros := make(map[int]device.Device)
	for i := 0; i < len(c.Devices); i++ {
		micros[c.Devices[i].ID] = c.Devices[i]
	}
	return micros
}

//AddDevice attempts to add a device to the cluster and returns the response data.
func (c *Cluster) AddDevice(newDevice device.Device) error {
	if viper.GetString("ENV") == "production" {
		for _, micro := range c.Devices {
			if micro.Host == newDevice.Host && newDevice.Port == micro.Port {
				return errors.New("This host & port combination are already registered to this cluster")
			}
		}
	}

	c.Devices = append(c.Devices, newDevice)

	PrintClusterInfo(*c)
	return nil
}

//RemoveDevice -
func (c *Cluster) RemoveDevice(ImDoneHere device.Device) {
	for index, device := range c.Devices {
		if device.ID == ImDoneHere.ID {
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

	// Device Discovery
	go func() {
		ticker := time.NewTicker(time.Duration(discoveryPeriod) * time.Second)
		for {
			select {
			case t := <-ticker.C:
				log.Println("Begin Device Discovery... ", t)
				DeviceDiscovery(c)
			}
		}
	}()

	// Health Check
	go func() {
		ticker := time.NewTicker(time.Duration(healthCheckPeriod) * time.Second)
		for {
			select {
			case t := <-ticker.C:
				log.Println("Begin Health Checks... ", t)
				for _, device := range c.Devices {
					c.DeviceHealthCheck(device)
				}
			}
		}
	}()
}

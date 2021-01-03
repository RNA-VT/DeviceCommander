package cluster

import (
	device "devicecommander/device"
	"devicecommander/utilities"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of Devices
type Cluster struct {
	Name    string
	Devices []device.Device
	Me      *device.Device
}

//Config -
type Config struct {
	Name    string `yaml:"Name"`
	Devices []device.Config
}

//GetConfig -
func (c Cluster) GetConfig() (config Config) {
	config.Name = c.Name
	config.Devices = make([]device.Config, len(c.Devices))
	for i, micro := range c.Devices {
		config.Devices[i] = micro.GetConfig()
	}

	return
}

//Load -
func (c *Cluster) Load(config Config) {
	c.Name = config.Name
	newPeers := []device.Device{
		*c.Me,
	}
	//c.Devices = make([]device.Device, len(config.Devices))
	for _, dev := range config.Devices {
		if c.Me.ID != dev.ID {
			newPeers = append(newPeers, device.Device{})
			newPeers[len(newPeers)-1].Load(dev)
		}
	}
	c.Devices = newPeers
}

func (c Cluster) String() string {
	cluster, err := utilities.StringJSON(c)
	if err != nil {
		return ""
	}
	return cluster
}

//GetDevices returns a map of all registered
func (c Cluster) GetDevices() map[int]device.Device {
	micros := make(map[int]device.Device)
	for i := 0; i < len(c.Devices); i++ {
		micros[c.Devices[i].ID] = c.Devices[i]
	}
	return micros
}

//AddDevice attempts to add a device to the cluster and returns the response data. This should only be run by the master.
func (c *Cluster) AddDevice(newdevice device.Config) error {
	var newGuy device.Device
	newGuy.Load(newdevice)
	if viper.GetString("ENV") == "production" {
		for _, micro := range c.Devices {
			if micro.Host == newGuy.Host {
				//This guy ain't so new!
				return errors.New("Requesting instance is running on a device already registered to this cluster")
			}
		}
	}

	c.Devices = append(c.Devices, newGuy)

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

//Start begins the registration and health check go processes
func (c *Cluster) Start() {
	// Device Discovery
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for {
			select {
			case t := <-ticker.C:
				log.Println("Begin Device Discovery...", t)
				for i := 1; i < 255; i++ {

					host := viper.Get("SUBNET_ROOT").(string) + strconv.Itoa(i)
					unregistered := !c.isRegistered(host)

					if unregistered {
						resp, err := http.Get("http://" + host + "/registration")
						if err != nil || resp.StatusCode != 200 {
							log.Println(host + " failed to respond to a registration request.")
						} else {
							log.Println(host + " is ready to register!")

						}
					}
				}
			}
		}
	}()

	// Health Check
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		failCount := 0
		failThreshold := 5
		for {
			select {
			case t := <-ticker.C:
				log.Println("Begin Heartbeat Check", t)
				for _, m := range c.Devices {
					log.Println("Checking Peer:", m.Name, m.ToFullAddress())
					url := "http://" + m.ToFullAddress() + "/v1/health"
					resp, err := http.Get(url)
					if err != nil || resp.StatusCode != 200 {
						log.Println(m.Name + " @" + m.ToFullAddress() + " is NOT ok")
						failCount++
						if failCount >= failThreshold {
							log.Println("Failure Threshold Reached. Deregistering Device...")
							c.RemoveDevice(m)
							failCount = 0
						}
					} else {
						log.Println(m.Name + " @" + m.ToFullAddress() + " is ok")
					}
				}
			}
		}
	}()
}

func (c Cluster) isRegistered(address string) bool {
	for _, m := range c.Devices {
		if address == m.ToFullAddress() {
			return true
		}
	}
	return false
}

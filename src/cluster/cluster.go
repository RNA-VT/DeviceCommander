package cluster

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/src/device"
)

type Cluster interface {
	Name() string
	PrintClusterInfo()
	Start()
	DeviceDiscovery(int)
	RunHealthCheckLoop(int)
	StopHealth()
	StopDiscovery()
	HandleDiscoveredDevice(model.NewDevice) (model.Device, error)
}

// Cluster is responsible for maintaining the cluster like state of DeviceCommander.
// It does things like probe the current active set for health and collection
// of new devices. The important differentiation between the Cluster and a repository
// of Devices is the active nature of the devices tracked in a Cluster. An active
// device is one that is currently responding to the device-commander protocol.
type DeviceCluster struct {
	name             string
	DeviceRepository device.Repository
	DeviceClient     device.Client
	logger           *log.Entry
	discoveryStop    chan bool
	healthStop       chan bool
}

func NewDeviceCluster(
	name string,
	deviceRepository device.Repository,
	deviceClient device.Client,
) DeviceCluster {
	return DeviceCluster{
		name:             name,
		DeviceRepository: deviceRepository,
		DeviceClient:     deviceClient,
		logger:           log.WithFields(log.Fields{"module": "cluster"}),
		discoveryStop:    make(chan bool),
		healthStop:       make(chan bool),
	}
}

// PrintClusterInfo will cleanly print out info about the cluster.
func (c DeviceCluster) PrintClusterInfo() {
	devices, err := c.DeviceRepository.GetAll()
	if err != nil {
		c.logger.Error(err)
		return
	}
	for i := 0; i < len(devices); i++ {
		c.logger.Println("----Device---")
		c.logger.Println(fmt.Sprintf("%+v", devices[i]))
	}
	c.logger.Println()
}

func (c DeviceCluster) Name() string {
	return c.name
}

// Start begins the collection of new devices (registration) and device health
// check goroutines.
func (c DeviceCluster) Start() {
	// Discovery and collection of new devices.
	go c.RunDeviceDiscoveryLoop(viper.GetInt("DISCOVERY_PERIOD"))

	// Health Check
	go c.RunHealthCheckLoop(viper.GetInt("HEALTH_CHECK_PERIOD"))
}

func (c DeviceCluster) StopHealth() {
	c.healthStop <- true
}

func (c DeviceCluster) StopDiscovery() {
	c.discoveryStop <- true
}

// RunDeviceDiscoveryLoop continuously searches for other responsive devices on the network.
func (c DeviceCluster) RunDeviceDiscoveryLoop(discoveryPeriod int) {
	ticker := time.NewTicker(time.Duration(discoveryPeriod) * time.Second)
	for range ticker.C {
		select {
		case <-c.discoveryStop:
			c.logger.Info("Ticker stopped by healthStop chan")
			return
		default:
			c.logger.Info("Begin Device Discovery... ")
			c.DeviceDiscovery(viper.GetInt("ARP_SCAN_DURATION"))
		}
	}
}

// RunHealthCheckLoop continuously probes the Health state of Active nodes stored in the DB.
// The important results of this health check will be tracked in the DB.
func (c DeviceCluster) RunHealthCheckLoop(healthCheckPeriod int) {
	ticker := time.NewTicker(time.Duration(healthCheckPeriod) * time.Second)
	for range ticker.C {
		select {
		case <-c.healthStop:
			c.logger.Info("Ticker stopped by healthStop chan")
			return
		default:
			devices, err := c.DeviceRepository.Get(model.Device{
				Active: true,
			})

			c.logger.Info(fmt.Sprintf("Begin Health Checks for %d devices... ", len(devices)))

			if err != nil {
				c.logger.Error(err)
				return
			}

			for _, d := range devices {
				dev := device.NewDeviceWrapper(*d)

				resp, err := c.DeviceClient.Health(dev)
				if err != nil {
					c.logger.Warn(fmt.Sprintf("error checking health for device [%s] %s", d.ID.String(), err))
				} else {
					result := c.DeviceClient.EvaluateHealthCheckResponse(resp, dev)
					if result {
						c.logger.Trace(fmt.Sprintf("device [%s] is healthy", d.ID.String()))
					} else {
						c.logger.Trace(fmt.Sprintf("device [%s] is not healthy", d.ID.String()))
					}
				}
			}
		}
	}
}

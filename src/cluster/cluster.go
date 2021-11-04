package cluster

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/device"
	"github.com/rna-vt/devicecommander/graph/model"
	"github.com/rna-vt/devicecommander/postgres"
)

type ICluster interface {
	PrintClusterInfo()
	Start()
	DeviceDiscovery()
}

// Cluster is responsible for maintaing the cluster like state of DeviceCommander.
// It does things like probe the current active set for health and collection
// of new devices.
type Cluster struct {
	Name          string
	DeviceService postgres.DeviceCRUDService
	DeviceClient  device.IDeviceClient
	logger        *log.Entry
	healthStop    chan bool
}

func NewCluster(
	name string,
	deviceService postgres.DeviceCRUDService,
	deviceClient device.IDeviceClient,
) Cluster {
	return Cluster{
		Name:          name,
		DeviceService: deviceService,
		DeviceClient:  deviceClient,
		logger:        log.WithFields(log.Fields{"module": "cluster"}),
		healthStop:    make(chan bool),
	}
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c Cluster) PrintClusterInfo() {
	devices, err := c.DeviceService.GetAll()
	if err != nil {
		c.logger.Error(err)
		return
	}
	for i := 0; i < len(devices); i++ {
		log.Println("----Device---")
		log.Println(fmt.Sprintf("%+v", devices[i]))
	}
	log.Println()
}

// Start begins the collection of new devices (registration) and device health
// check goroutines.
func (c *Cluster) Start() {
	// Discovery and collection of new devices.
	go c.RunDeviceDiscoveryLoop(viper.GetInt("DISCOVERY_PERIOD"))

	// Health Check
	go c.RunHealthCheckLoop(viper.GetInt("HEALTH_CHECK_PERIOD"))
}

func (c Cluster) RunDeviceDiscoveryLoop(discoveryPeriod int) {
	ticker := time.NewTicker(time.Duration(discoveryPeriod) * time.Second)
	for range ticker.C {
		c.logger.Info("Begin Device Discovery... ")
		c.DeviceDiscovery()
	}
}

func (c Cluster) RunHealthCheckLoop(healthCheckPeriod int) {
	ticker := time.NewTicker(time.Duration(healthCheckPeriod) * time.Second)
	for range ticker.C {
		select {
		case <-c.healthStop:
			c.logger.Info("Ticker stopped by healthStop chan")
			return
		default:
			devs, err := c.DeviceService.Get(model.Device{
				Active: true,
			})

			c.logger.Info(fmt.Sprintf("Begin Health Checks for %d devices... ", len(devs)))

			if err != nil {
				c.logger.Error(err)
				return
			}

			for _, d := range devs {
				dev := device.NewDeviceWrapper(*d)

				resp, err := c.DeviceClient.Health(dev)
				if err != nil {
					c.logger.Warn(fmt.Sprintf("error checking health for device [%s] %s", d.ID.String(), err))
				}

				result := c.DeviceClient.EvaluateHealthCheckResponse(resp, dev)
				if result {
					log.Trace(fmt.Sprintf("device [%s] is healthy", d.ID.String()))
				} else {
					log.Trace(fmt.Sprintf("device [%s] is not healthy", d.ID.String()))
				}
			}
		}
	}
}

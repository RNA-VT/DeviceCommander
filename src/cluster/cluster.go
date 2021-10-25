package cluster

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/device"
	"github.com/rna-vt/devicecommander/postgres"
)

// Cluster is responsible for maintaing the cluster like state of DeviceCommander.
// It does things like probe the current active set for health and collection
// of new devices.
type Cluster struct {
	Name          string
	DeviceService postgres.DeviceCRUDService
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c Cluster) PrintClusterInfo() {
	logger := getClusterLogger()
	devices, err := c.DeviceService.GetAll()
	if err != nil {
		logger.Error(err)
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
	logger := getClusterLogger()
	discoveryPeriod := viper.GetInt("DISCOVERY_PERIOD")
	healthCheckPeriod := viper.GetInt("HEALTH_CHECK_PERIOD")
	clusterLogger := getClusterLogger()

	// Discover and collection of new devices.
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
			devices, err := c.DeviceService.GetAll()
			if err != nil {
				logger.Error(err)
				return
			}

			for _, d := range devices {
				tmp, err := device.NewDeviceObj(d)
				if err != nil {
					logger.Error(err)
				} else {
					tmp.CheckHealth()
				}
			}
		}
	}()
}

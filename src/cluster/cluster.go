package cluster

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/rna-vt/devicecommander/device"
	"github.com/rna-vt/devicecommander/postgres"
)

// Cluster - This object defines an array of Devices
type Cluster struct {
	Name          string
	DeviceService postgres.DeviceService
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

// Start begins the registration and health check goroutines
func (c *Cluster) Start() {
	logger := getClusterLogger()
	discoveryPeriod := viper.GetInt("DISCOVERY_PERIOD")
	healthCheckPeriod := viper.GetInt("HEALTH_CHECK_PERIOD")
	clusterLogger := getClusterLogger()

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

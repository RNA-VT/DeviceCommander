package cluster

import (
	dev "devicecommander/device"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

//generateUniqueID returns a unique id for assigning to a new device
func (c Cluster) generateUniqueID() int {
	limit := viper.GetInt("MICROCONTORLLER_LIMIT")
	randID := rand.Intn(limit)
	for len(c.getDeviceByID(randID)) > 0 {
		randID = rand.Intn(limit)
	}
	return randID
}

// getSlaveByID find all the slave for a given ID
func (c Cluster) getDeviceByID(targetID int) []dev.Device {
	var micros []dev.Device

	for i := 0; i < len(c.Devices); i++ {
		if c.Devices[i].ID == targetID {
			return append(micros, c.Devices[i])
		}
	}

	return micros
}

func isExcluded(m dev.Device, exclusions []dev.Device) bool {
	for i := 0; i < len(exclusions); i++ {
		if m.Host == exclusions[i].Host && m.Port == exclusions[i].Port {
			return true
		}
	}
	return false
}

// PrintClusterInfo will cleanly print out info about the cluster
func PrintClusterInfo(c Cluster) {
	for i := 0; i < len(c.Devices); i++ {
		log.Println("----Device---")
		log.Println(c.Devices[i])
	}
	log.Println()
}

// test check if master exists
func test(URL string) error {
	parsedURL, err := url.Parse("http://" + URL)
	if err != nil {
		log.Println("Failed to Parse URL")
		return err
	}
	_, err = http.Get(parsedURL.String())
	return err
}

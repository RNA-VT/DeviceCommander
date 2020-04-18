package cluster

import (
	mc "firecontroller/microcontroller"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

//generateUniqueID returns a unique id for assigning to a new microcontroller
func (c Cluster) generateUniqueID() int {
	limit := viper.GetInt("MICROCONTORLLER_LIMIT")
	randID := rand.Intn(limit)
	for len(c.getSlavesByID(randID)) > 0 {
		randID = rand.Intn(limit)
	}
	return randID
}

// getSlaveByID find all the slave for a given ID
func (c Cluster) getSlavesByID(targetID int) []mc.Microcontroller {
	var micros []mc.Microcontroller

	for i := 0; i < len(c.Microcontrollers); i++ {
		if c.Microcontrollers[i].ID == targetID {
			return append(micros, c.Microcontrollers[i])
		}
	}

	return micros
}

func isExcluded(m mc.Microcontroller, exclusions []mc.Config) bool {
	for i := 0; i < len(exclusions); i++ {
		if m.Host == exclusions[i].Host && m.Port == exclusions[i].Port {
			return true
		}
	}
	return false
}

// PrintClusterInfo will cleanly print out info about the cluster
func PrintClusterInfo(c Cluster) {
	log.Println()
	log.Println("====Master====")
	log.Println(c.Master())

	log.Println()

	for i := 0; i < len(c.Microcontrollers); i++ {
		log.Println("----Peer---")
		log.Println(c.Microcontrollers[i])
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

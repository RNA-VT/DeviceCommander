package cluster

import (
	"log"
)

// PrintClusterInfo will cleanly print out info about the cluster
func PrintClusterInfo(c Cluster) {
	for i := 0; i < len(c.Devices); i++ {
		log.Println("----Device---")
		log.Println(c.Devices[i])
	}
	log.Println()
}

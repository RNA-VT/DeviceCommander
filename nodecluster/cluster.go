package nodecluster

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Cluster struct {
	SlaveNodes []NodeInfo
	MasterNode NodeInfo
}

func (c *Cluster) AddSlaveNode(node NodeInfo) {
	c.SlaveNodes = append(c.SlaveNodes, node)
}

// GenerateUniqueID returns a unique id for asigning to a new node
func (c *Cluster) GenerateUniqueID() int {
	randID := rand.Intn(100)
	for len(c.GetSlaveByID(randID)) > 0 {
		randID = rand.Intn(100)
		fmt.Println(len(c.GetSlaveByID(randID)))
	}
	return randID
}

// GenerateUniquePort returns a unique id for asigning to a new node
func (c *Cluster) GenerateUniquePort(targetIP string) string {
	randPort := 8002
	nodesInQuestion := c.GetAllSlavesByIP(targetIP)

	for i := 0; i < len(nodesInQuestion); i++ {
		if nodesInQuestion[i].Port == strconv.Itoa(randPort) {
			randPort++
		}
	}

	return strconv.Itoa(randPort)
}

// GetSlaveByID find a slave node by its ID
func (c *Cluster) GetSlaveByID(targetID int) []NodeInfo {
	var nodes []NodeInfo

	for i := 0; i < len(c.SlaveNodes); i++ {
		if c.SlaveNodes[i].NodeId == targetID {
			return append(nodes, c.SlaveNodes[i])
		}
	}

	return nodes
}

// GetAllSlavesByIP find all slave node by its IP
func (c *Cluster) GetAllSlavesByIP(targetIP string) []NodeInfo {
	var nodes []NodeInfo

	for i := 0; i < len(c.SlaveNodes); i++ {
		if c.SlaveNodes[i].NodeIpAddr == targetIP {
			nodes = append(nodes, c.SlaveNodes[i])
		}
	}

	return nodes
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c *Cluster) PrintClusterInfo() {
	fmt.Println()
	fmt.Println("====Master====")
	fmt.Println(c.MasterNode)

	fmt.Println()

	for i := 0; i < len(c.SlaveNodes); i++ {
		fmt.Println("----Node---")
		fmt.Println(c.SlaveNodes[i])
	}
	fmt.Println()
}

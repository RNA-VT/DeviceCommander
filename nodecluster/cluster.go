package nodecluster

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Cluster struct {
	SlaveNodes []NodeInfo
	MasterIp   *string
	MasterNode NodeInfo
}

func (c *Cluster) AddSlaveNode(node NodeInfo) {
	c.SlaveNodes = append(c.SlaveNodes, node)
}

// GenerateUniqueID returns a unique id for asigning to a new node
func (c *Cluster) GenerateUniqueID() int {
	randID := rand.Intn(100)
	fmt.Println(randID)
	for len(c.GetSlaveByID(randID)) > 0 {
		randID = rand.Intn(100)
		fmt.Println(randID)
		fmt.Println(len(c.GetSlaveByID(randID)))
	}
	return randID
}

// GenerateUniquePort returns a unique id for asigning to a new node
func (c *Cluster) GenerateUniquePort(targetIP string) string {
	randPort := 8001
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

	fmt.Println(c.SlaveNodes)

	for i := 0; i < len(c.SlaveNodes); i++ {
		if c.SlaveNodes[i].NodeId == targetID {
			return append(nodes, c.SlaveNodes[i])
		}
	}

	fmt.Println(nodes)
	return nodes
}

// GetAllSlavesByIP find all slave node by its IP
func (c *Cluster) GetAllSlavesByIP(targetIP string) []NodeInfo {
	var nodes []NodeInfo

	fmt.Println(c.SlaveNodes)

	for i := 0; i < len(c.SlaveNodes); i++ {
		if c.SlaveNodes[i].NodeIpAddr == targetIP {
			nodes = append(nodes, c.SlaveNodes[i])
		}
	}

	fmt.Println(nodes)
	return nodes
}

func (c *Cluster) PrintClusterInfo() {
	fmt.Println(c)
}

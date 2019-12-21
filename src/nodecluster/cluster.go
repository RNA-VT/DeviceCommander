package nodecluster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Cluster object definition
type Cluster struct {
	SlaveNodes []NodeInfo
	MasterNode NodeInfo
}

//AddSlaveNode appends a device controller to the node list
func (c *Cluster) AddSlaveNode(node NodeInfo) {
	c.SlaveNodes = append(c.SlaveNodes, node)
}

//GenerateUniqueID returns a unique id for asigning to a new node
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
		if c.SlaveNodes[i].NodeID == targetID {
			return append(nodes, c.SlaveNodes[i])
		}
	}

	return nodes
}

// GetAllSlavesByIP find all slave node by its IP
func (c *Cluster) GetAllSlavesByIP(targetIP string) []NodeInfo {
	var nodes []NodeInfo

	for i := 0; i < len(c.SlaveNodes); i++ {
		if c.SlaveNodes[i].NodeIPAddr == targetIP {
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

// SendMessageToAllNodes will take a byte slice and POST it to each node
func (c *Cluster) SendMessageToAllNodes(urlPath string, message AddToClusterMessage, excludeNodes []NodeInfo) error {
	for i := 0; i < len(c.SlaveNodes); i++ {
		if !isExcluded(c.SlaveNodes[i], excludeNodes) {
			bytesRepresentation, err := json.Marshal(c)
			if err != nil {
				fmt.Println("Failed to convert cluster to json: ", c)
				return err
			}
			currURL := c.SlaveNodes[i].ToFullAddress() + urlPath

			resp, err := http.Post(currURL, "application/json", bytes.NewBuffer(bytesRepresentation))
			if err != nil {
				fmt.Println("WARNING: Failed to POST to Peer: ", c.SlaveNodes[i].String(), currURL)
			} else {
				var result map[string]interface{}
				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&result)
				log.Println(result)
				log.Println(result["data"])
			}
		}
	}
	return nil
}

func isExcluded(node NodeInfo, exclusions []NodeInfo) bool {
	for i := 0; i < len(exclusions); i++ {
		if node.String() == exclusions[i].String() {
			return true
		}
	}
	return false
}

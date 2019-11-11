package nodecluster

import (
	"fmt"
)

type Cluster struct {
	SlaveNodes []NodeInfo
	MasterIp   *string
}

func (cluster *Cluster) AddSlaveNode(node NodeInfo) {
	cluster.SlaveNodes = append(cluster.SlaveNodes, node)
}

func (cluster *Cluster) PrintClusterInfo() {
	fmt.Println(cluster)
}

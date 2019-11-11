package nodecluster

import (
	"strconv"
)

/* Information/Metadata about node */
type NodeInfo struct {
	NodeId     int
	NodeIpAddr string
	Port       string
}

/* Just for pretty printing the node info */
func (node NodeInfo) String() string {
	return "NodeInfo:{ nodeId:" + strconv.Itoa(node.NodeId) + ", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + " }"
}

/* Just for pretty printing the node info */
func (node NodeInfo) ToFullAdress() string {
	return node.NodeIpAddr + ":" + node.Port
}

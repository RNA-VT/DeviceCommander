package nodecluster

import (
	"strconv"
)

// NodeInfo Information/Metadata about node
type NodeInfo struct {
	NodeId     int
	NodeIpAddr string
	Port       string
}

/* Just for pretty printing the node info */
func (node *NodeInfo) String() string {
	return "NodeInfo:{ nodeId:" + strconv.Itoa(node.NodeId) + ", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + " }"
}

// ToFullAddress returns a network address including the ip address and port that this node is listening on
func (node *NodeInfo) ToFullAddress() string {
	/* Just for pretty printing the node info */
	return node.NodeIpAddr + ":" + node.Port
}

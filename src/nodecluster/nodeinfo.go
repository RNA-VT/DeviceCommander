package nodecluster

import (
	"strconv"
)

//NodeInfo Information/Metadata about node
type NodeInfo struct {
	NodeID int
	Host   string
	Port   string
}

/*String Just for pretty printing the node info */
func (node *NodeInfo) String() string {
	//strings.Split(myIP[0].String(), "/")[0]
	return "NodeInfo:{ nodeId:" + strconv.Itoa(node.NodeID) + ", Host:" + node.Host + ", port:" + node.Port + " }"
}

//ToFullAddress returns a network address including the ip address and port that this node is listening on
func (node *NodeInfo) ToFullAddress() string {
	/* Just for pretty printing the node info */
	return node.Host + ":" + node.Port
}

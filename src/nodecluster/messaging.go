package nodecluster

// "fmt"
// "net"
// "time"
// "encoding/json"

/*AddToClusterMessage A standard format for a Request/Response for adding node to cluster */
type AddToClusterMessage struct {
	Source  NodeInfo
	Dest    NodeInfo
	Cluster Cluster
}

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	Node NodeInfo
}

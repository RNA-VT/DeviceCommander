package nodecluster

// "fmt"
// "net"
// "time"
// "encoding/json"

/* A standard format for a Request/Response for adding node to cluster */
type AddToClusterMessage struct {
	Source  NodeInfo
	Dest    NodeInfo
	Cluster Cluster
}

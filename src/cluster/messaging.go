package cluster

import (
	"firecontroller/device"
)

/*AddToClusterMessage A standard format for a Request/Response for adding device to cluster */
type AddToClusterMessage struct {
	Source  device.Device
	Dest    device.Device
	Cluster Cluster
}

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	Device device.Device
}

//PeerUpdateMessage contains a source and cluster info
type PeerUpdateMessage struct {
	Source  device.Device
	Cluster Cluster
}

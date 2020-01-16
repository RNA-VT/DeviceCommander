package cluster

import (
	"firecontroller/device"
)

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	Device device.Device
}

//PeerUpdateMessage contains a source and cluster info
type PeerUpdateMessage struct {
	Source  device.Device
	Cluster Cluster
}

package cluster

import mc "firecontroller/microcontroller"

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	ImNewHere mc.Microcontroller
}

//PeerUpdateMessage contains a source and cluster info
type PeerUpdateMessage struct {
	Source  mc.Microcontroller
	Cluster Cluster
}

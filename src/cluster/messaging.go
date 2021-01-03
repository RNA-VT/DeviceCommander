package cluster

import (
	mc "devicecommander/device"
	"time"
)

//PeerErrorMessage -
type PeerErrorMessage struct {
	Panic        bool
	DeregisterMe mc.Device
	PeerInfoMessage
}

//PeerInfoMessage -
type PeerInfoMessage struct {
	Messages []string
	Header   GoFireHeader
}

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	ImNewHere mc.Config
}

//CommandMessage -
type CommandMessage struct {
	Command       string
	ComponentType string
}

//GoFireHeader -
type GoFireHeader struct {
	Source  mc.Config
	Created time.Time
}

//GetHeader -
func (c Cluster) GetHeader() GoFireHeader {
	return GoFireHeader{
		Source:  c.Me.GetConfig(),
		Created: time.Now(),
	}
}

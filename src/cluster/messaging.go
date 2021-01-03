package cluster

import (
	device "devicecommander/device"
)

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	ImNewHere device.Device
}

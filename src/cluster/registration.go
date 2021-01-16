package cluster

import (
	device "devicecommander/device"
	"encoding/json"
	"io"
	"log"
)

//DecodeRegistrationRequest - helper to get new device details from a registration request msg body
func DecodeRegistrationRequest(body io.ReadCloser) device.Device {
	decoder := json.NewDecoder(body)
	var msg JoinNetworkMessage
	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}
	return msg.ImNewHere
}

func (c Cluster) isRegistered(address string) bool {
	for _, m := range c.Devices {
		if address == m.ToFullAddress() {
			return true
		}
	}
	return false
}

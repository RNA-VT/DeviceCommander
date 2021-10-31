package test

import (
	"crypto/rand"
	"fmt"

	"github.com/rna-vt/devicecommander/graph/model"
)

func GenerateRandomMacAddress() string {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	// Set the local bit
	buf[0] |= 2
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

func GenerateRandomNewDevices(count int) []model.NewDevice {
	collection := []model.NewDevice{}
	for i := 0; i < count; i++ {
		tmpMac := GenerateRandomMacAddress()
		tmpDev := model.NewDevice{
			Mac: &tmpMac,
		}
		collection = append(collection, tmpDev)
	}
	return collection
}

func GenerateRandomNewEndpoints(deviceID string, count int) []model.NewEndpoint {
	collection := []model.NewEndpoint{}
	for i := 0; i < count; i++ {
		tmpEndpointName := "dmx-config"
		tmpEndpoint := model.NewEndpoint{
			DeviceID: deviceID,
			Method:   &tmpEndpointName,
			Type:     "SET",
			Parameters: []*model.Parameter{
				{
					Name: "universe",
					Type: "int",
				},
				{
					Name: "start-address",
					Type: "int",
				},
			},
		}

		tmpEndpointName2 := "state"
		tmpEndpoint2 := model.NewEndpoint{
			DeviceID: deviceID,
			Method:   &tmpEndpointName2,
			Type:     "GET",
			Parameters: []*model.Parameter{
				{
					Name: "universe",
					Type: "int",
				},
				{
					Name: "start-address",
					Type: "int",
				},
			},
		}
		collection = append(collection, tmpEndpoint)
		collection = append(collection, tmpEndpoint2)
	}
	return collection
}

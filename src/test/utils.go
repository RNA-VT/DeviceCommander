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
		tmpParam := GenerateRandomNewParameter(2)
		tmpEndpoint := model.NewEndpoint{
			DeviceID: deviceID,
			Method:   "dmx-config",
			Type:     "SET",
			Parameters: []*model.NewParameter{
				&tmpParam[0],
				&tmpParam[1],
			},
		}

		tmpEndpoint2 := model.NewEndpoint{
			DeviceID: deviceID,
			Method:   "state",
			Type:     "GET",
			Parameters: []*model.NewParameter{
				&tmpParam[0],
				&tmpParam[1],
			},
		}
		collection = append(collection, tmpEndpoint)
		collection = append(collection, tmpEndpoint2)
	}
	return collection
}

func GenerateRandomNewParameter(count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpParam := model.NewParameter{
			Name: "foobar",
			Type: "SET",
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

func GenerateRandomNewParameterForEndpoint(endpointID string, count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpParam := model.NewParameter{
			EndpointID: endpointID,
			Name:       "foobar",
			Type:       "SET",
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

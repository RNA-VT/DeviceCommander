package test

import (
	"fmt"
	"math/rand"
	"time"

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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 124567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRandomBoolean() bool {
	return rand.Intn(1) > 0
}

func GenerateRandomNewDevices(count int) []model.NewDevice {
	collection := []model.NewDevice{}
	for i := 0; i < count; i++ {
		tmpName := GenerateRandomString(10)
		tmpDesc := GenerateRandomString(100)
		tmpMac := GenerateRandomMacAddress()

		tmpDev := model.NewDevice{
			Name:        &tmpName,
			Description: &tmpDesc,
			Mac:         &tmpMac,
			Host:        "127.0.0.1",
			Port:        9100,
		}
		collection = append(collection, tmpDev)
	}
	return collection
}

func GenerateRandomNewEndpoints(deviceID string, count int) []model.NewEndpoint {
	collection := []model.NewEndpoint{}
	for i := 0; i < count; i++ {
		tmpParam := GenerateRandomNewParameter(2)
		tmpDesc := GenerateRandomString(100)
		tmpEndpoint := model.NewEndpoint{
			DeviceID:    deviceID,
			Method:      "dmx-config",
			Type:        "SET",
			Description: &tmpDesc,
			Parameters: []*model.NewParameter{
				&tmpParam[0],
				&tmpParam[1],
			},
		}
		collection = append(collection, tmpEndpoint)
	}
	return collection
}

func GenerateRandomNewParameter(count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpDesc := GenerateRandomString(100)
		tmpParam := model.NewParameter{
			Name:        GenerateRandomString(5),
			Description: &tmpDesc,
			Type:        "SET",
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

func GenerateRandomNewParameterForEndpoint(endpointID string, count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpDesc := GenerateRandomString(100)
		tmpParam := model.NewParameter{
			EndpointID:  endpointID,
			Name:        "foobar",
			Description: &tmpDesc,
			Type:        "SET",
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

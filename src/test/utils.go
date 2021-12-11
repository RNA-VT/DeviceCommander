package test

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"

	"github.com/rna-vt/devicecommander/graph/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomNewDevices(count int) []model.NewDevice {
	collection := []model.NewDevice{}
	for i := 0; i < count; i++ {
		tmpDev := model.NewDevice{}
		err := faker.FakeData(&tmpDev)
		if err != nil {
			fmt.Println(err)
		}
		collection = append(collection, tmpDev)
	}
	return collection
}

func GenerateRandomNewEndpoints(deviceID string, count int) []model.NewEndpoint {
	collection := []model.NewEndpoint{}
	for i := 0; i < count; i++ {
		tmpEndpoint := model.NewEndpoint{}
		err := faker.FakeData(&tmpEndpoint)
		if err != nil {
			fmt.Println(err)
		}

		tmpEndpoint.DeviceID = deviceID
		collection = append(collection, tmpEndpoint)
	}
	return collection
}

func GenerateRandomNewParameter(count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpParam := model.NewParameter{}
		err := faker.FakeData(&tmpParam)
		if err != nil {
			fmt.Println(err)
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

func GenerateRandomNewParameterForEndpoint(endpointID string, count int) []model.NewParameter {
	collection := []model.NewParameter{}
	for i := 0; i < count; i++ {
		tmpParam := model.NewParameter{}
		err := faker.FakeData(&tmpParam)
		if err != nil {
			fmt.Println(err)
		}
		tmpParam.EndpointID = endpointID

		collection = append(collection, tmpParam)
	}
	return collection
}

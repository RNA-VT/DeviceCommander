package test

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/graph/model"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomNewDevices(count int) []model.NewDevice {
	collection := make([]model.NewDevice, count)
	for index := range collection {
		tmpItem := model.NewDevice{}
		err := faker.FakeData(&tmpItem)
		if err != nil {
			log.Fatal(err)
		}
		collection[index] = tmpItem
	}
	return collection
}

func GenerateRandomNewEndpoints(deviceID string, count int) []model.NewEndpoint {
	collection := []model.NewEndpoint{}
	for i := 0; i < count; i++ {
		tmpEndpoint := model.NewEndpoint{}
		err := faker.FakeData(&tmpEndpoint)
		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
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
			log.Fatal(err)
		}
		tmpParam.EndpointID = endpointID

		collection = append(collection, tmpParam)
	}
	return collection
}

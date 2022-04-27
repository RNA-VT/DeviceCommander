package device

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

func GenerateRandomNewDeviceParams(count int) []NewDeviceParams {
	collection := make([]NewDeviceParams, count)
	for index := range collection {
		tmpItem := NewDeviceParams{}
		err := faker.FakeData(&tmpItem)
		if err != nil {
			log.Fatal(err)
		}
		collection[index] = tmpItem
	}
	return collection
}

func GenerateRandomNewEndpointParams(deviceID string, count int) []NewEndpointParams {
	collection := []NewEndpointParams{}
	for i := 0; i < count; i++ {
		tmpEndpoint := NewEndpointParams{}
		err := faker.FakeData(&tmpEndpoint)
		if err != nil {
			log.Fatal(err)
		}

		tmpEndpoint.DeviceID = deviceID
		collection = append(collection, tmpEndpoint)
	}
	return collection
}

func GenerateRandomNewParameter(count int) []NewParameterParams {
	collection := []NewParameterParams{}
	for i := 0; i < count; i++ {
		tmpParam := NewParameterParams{}
		err := faker.FakeData(&tmpParam)
		if err != nil {
			log.Fatal(err)
		}

		collection = append(collection, tmpParam)
	}
	return collection
}

func GenerateRandomNewParameterForEndpoint(endpointID string, count int) []NewParameterParams {
	collection := []NewParameterParams{}
	for i := 0; i < count; i++ {
		tmpParam := NewParameterParams{}
		err := faker.FakeData(&tmpParam)
		if err != nil {
			log.Fatal(err)
		}
		tmpParam.EndpointID = endpointID

		collection = append(collection, tmpParam)
	}
	return collection
}

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

package endpoint

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

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

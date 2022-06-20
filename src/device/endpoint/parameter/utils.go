package parameter

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

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

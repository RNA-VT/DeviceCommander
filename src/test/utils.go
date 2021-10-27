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

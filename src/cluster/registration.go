package cluster

import (
	device "devicecommander/device"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

//DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	for i := 1; i < 255; i++ {

		host := viper.Get("IP_ADDRESS_ROOT").(string) + strconv.Itoa(i)
		unregistered := !c.isRegistered(host)

		if unregistered {
			url := "http://" + host + "/registration"
			log.Println("[Registration] Attempting to Register: " + url)
			resp, err := http.Get(url)
			if err != nil {
				log.Println("[Registration] Attempt to register " + host + " resulted in an error:")
				log.Println(err)
			}
			switch resp.StatusCode {
			case 200:
				log.Println("[Registration] Registration Request Accepted: " + host)
				log.Println("[Registration] Adding New Device...")
				c.AddDevice(DecodeRegistrationRequest(resp.Body))
			case 404:
				log.Println("[Registration] Host Not Found: " + host)
			default:
				log.Println("[Registration] Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
			}
		}
	}
}

//DecodeRegistrationRequest - helper to get new device details from a registration request msg body
func DecodeRegistrationRequest(body io.ReadCloser) device.Device {
	decoder := json.NewDecoder(body)
	var dev device.Device
	err := decoder.Decode(&dev)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}
	return dev
}

func (c Cluster) isRegistered(address string) bool {
	for _, m := range c.Devices {
		if address == m.ToFullAddress() {
			return true
		}
	}
	return false
}

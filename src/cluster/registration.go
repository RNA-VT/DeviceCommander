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
			//TODO: Hit ports other than 80
			//TODO: Scan https for /registration
			url := "http://" + host + "/registration"
			log.Println("[Registration] Attempting to Register: " + url)
			resp, err := http.Get(url)
			if err != nil {
				log.Println("[Registration] [Error] Attempt to register " + host + " resulted in an error:")
				log.Println(err)
			} else {
				switch resp.StatusCode {
				case 200:
					log.Println("[Registration] [Success] Registration Request Accepted: " + host)
					log.Println("[Registration] [Success] Adding New Device...")
					dev := DeviceFromRegistrationRequestBody(resp.Body)
					c.AddDevice(dev)
				case 404:
					log.Println("[Registration] [Warning] Host Not Found: " + host)
				default:
					log.Println("[Registration] [Warning] Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
				}
			}
		}
	}
}

//DeviceFromRegistrationRequestBody - helper to get new device details from a registration request msg body
func DeviceFromRegistrationRequestBody(body io.ReadCloser) device.Device {
	defer body.Close()
	decoder := json.NewDecoder(body)
	var dev device.Device
	err := decoder.Decode(&dev)
	if err != nil {
		log.Println("[Registration] [Error] Failed to decode device config from body", err)
	}
	return dev
}

func (c Cluster) isRegistered(host string) bool {
	for _, dev := range c.Devices {
		if host == dev.Host {
			return true
		}
	}
	return false
}

package cluster

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	device "devicecommander/device"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

//DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	registrationLogger := getRegistrationLogger()

	for i := 1; i < 255; i++ {
		host := viper.Get("IP_ADDRESS_ROOT").(string) + strconv.Itoa(i)
		dev, err := c.GetDeviceByHost(host)

		if err == nil {
			//TODO: Hit ports other than 80
			//TODO: Scan https for /registration
			url := "http://" + host + "/registration"
			registrationLogger.Info("Attempting to Register: " + url)
			resp, err := http.Get(url)
			if err != nil {
				registrationLogger.Error("Attempt to register "+host+" resulted in an error: ", err)
			} else {
				switch resp.StatusCode {
				case 200:
					registrationSuccessLogger := registrationLogger.WithFields(log.Fields{
						"event": "success",
					})
					dev, err = device.NewDeviceFromRequestBody(resp.Body)
					if err != nil {
						return
					}

					registrationSuccessLogger.Info("Registration Request Accepted: " + host)
					registrationSuccessLogger.Info("Adding new Device: " + dev.ID)
					c.AddDevice(dev)
				case 404:
					registrationLogger.Debug("Host Not Found: " + host)
				default:
					registrationLogger.Debug("Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
				}
			}
		}
	}
}

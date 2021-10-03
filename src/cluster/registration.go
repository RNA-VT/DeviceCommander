package cluster

import (
	"fmt"

	"github.com/schollz/peerdiscovery"
	log "github.com/sirupsen/logrus"
)

func getRegistrationLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "registration"})
}

//DeviceDiscovery -
func DeviceDiscovery(c *Cluster) {
	registrationLogger := getRegistrationLogger()

	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit: 5,
		Notify: func(d peerdiscovery.Discovered) {
			log.Println(d)
			registrationLogger.Info("discovered ", d.Address)
		},
	})

	if len(discoveries) > 0 {
		fmt.Printf("Found %d other computers\n", len(discoveries))
		for i, d := range discoveries {
			fmt.Printf("%d) '%s' with payload '%s'\n", i, d.Address, d.Payload)
		}
	} else {
		fmt.Println("Found no devices. You need to run this on another computer at the same time.")
	}

	// for _, d := range discoveries {
	// 	registrationLogger.Info("discovered " + d.Address)
	// }

	// for i := 1; i < 255; i++ {
	// 	host := viper.Get("IP_ADDRESS_ROOT").(string) + strconv.Itoa(i)
	// 	dev, err := c.GetDeviceByHost(host)

	// 	if err == nil {
	// 		//TODO: Hit ports other than 80
	// 		//TODO: Scan https for /registration
	// 		url := "http://" + host + "/registration"
	// 		registrationLogger.Info("Attempting to Register: " + url)
	// 		resp, err := http.Get(url)
	// 		if err != nil {
	// 			registrationLogger.Error("Attempt to register "+host+" resulted in an error: ", err)
	// 		} else {
	// 			switch resp.StatusCode {
	// 			case 200:
	// 				registrationSuccessLogger := registrationLogger.WithFields(log.Fields{
	// 					"event": "success",
	// 				})
	// 				dev, err = device.NewDeviceFromRequestBody(resp.Body)
	// 				if err != nil {
	// 					return
	// 				}

	// 				registrationSuccessLogger.Info("Registration Request Accepted: " + host)
	// 				registrationSuccessLogger.Info("Adding new Device: " + dev.ID)
	// 				c.AddDevice(dev)
	// 			case 404:
	// 				registrationLogger.Debug("Host Not Found: " + host)
	// 			default:
	// 				registrationLogger.Debug("Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
	// 			}
	// 		}
	// 	}
	// }
}

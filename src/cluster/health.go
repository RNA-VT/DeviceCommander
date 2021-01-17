package cluster

import (
	device "devicecommander/device"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var failCounts map[string]int

//DeviceHealthCheck -
func (c *Cluster) DeviceHealthCheck(dev device.Device) {
	failThreshold := 3
	url := "http://" + dev.ToFullAddress() + "/v1/health"
	log.Println("[Health] Checking Device:", dev.ID, url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[Health] [Error] : " + url)
		log.Println("[Health] [Error] Status Code: " + strconv.Itoa(resp.StatusCode))
		log.Println("[Health] [Error] Message: " + err.Error())
	}

	healthy := evaluateHealthCheckResponse(resp, url)

	if healthy {
		failCounts[dev.ID] = 0
	} else {
		failCounts[dev.ID]++
		if failCounts[dev.ID] >= failThreshold {
			log.Println("[Health] [Deregistration] Failure Threshold Reached.")
			log.Println("[Health] [Deregistration] Removing Device: " + dev.ID)
			c.RemoveDevice(dev)
			failCounts[dev.ID] = 0
		}
	}
}

func evaluateHealthCheckResponse(resp *http.Response, url string) bool {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	healthy := false
	switch resp.StatusCode {
	case 200:
		log.Println("[Health] [Success]: " + url)
	case 404:
		log.Println("[Health] [Failure] Registed Device Not Found: " + url)
	default:
		log.Println("[Health] [Failure] Unexpected Result: " + url)
		log.Println("[Health] [Failure] Status Code: " + strconv.Itoa(resp.StatusCode))
		log.Println("[Health] [Failure] Response: " + string(body))
	}
	return healthy
}

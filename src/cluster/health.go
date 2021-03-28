package cluster

import (
	device "devicecommander/device"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//DeviceHealthCheck -
func (c *Cluster) DeviceHealthCheck(dev *device.Device) {
	url := dev.URL() + "/health"

	log.Println("[Health] Checking Device:", dev.ID, url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[Health] [Error] : " + url)
		log.Println("[Health] [Error] Status Code: " + strconv.Itoa(resp.StatusCode))
		log.Println("[Health] [Error] Message: " + err.Error())
	}

	result := evaluateHealthCheckResponse(resp, *dev)
	dev.ProcessHealthCheckResult(result)

	if dev.Failed() {
		log.Println("[Health] [Deregistration] Failure Threshold Reached.")
		log.Println("[Health] [Deregistration] Removing Device: " + dev.ID)
		c.RemoveDevice(dev.ID)
	}
}

func evaluateHealthCheckResponse(resp *http.Response, dev device.Device) bool {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	healthy := false
	switch resp.StatusCode {
	case 200:
		log.Println("[Health] [Success] " + dev.ID + " is Healthy")
		healthy = true
	case 404:
		log.Println("[Health] [Failure] Registered Device Not Found: " + dev.ID)
	default:
		log.Println("[Health] [Failure] Unexpected Result: " + dev.ID)
		log.Println("[Health] [Failure] Status Code: " + strconv.Itoa(resp.StatusCode))
		log.Println("[Health] [Failure] Response: " + string(body))
	}
	return healthy
}

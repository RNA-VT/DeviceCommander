package cluster

import (
	device "devicecommander/device"
	"log"
	"net/http"
)

//DeviceHealthCheck -
func (c *Cluster) DeviceHealthCheck(m device.Device) {
	failCount := 0
	failThreshold := 3
	log.Println("Checking Peer:", m.Name, m.ToFullAddress())
	url := "http://" + m.ToFullAddress() + "/v1/health"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		log.Println(m.Name + " @" + m.ToFullAddress() + " is NOT ok")
		failCount++
		if failCount >= failThreshold {
			log.Println("Failure Threshold Reached. Deregistering Device...")
			c.RemoveDevice(m)
			failCount = 0
		}
	} else {
		log.Println(m.Name + " @" + m.ToFullAddress() + " is ok")
	}
}

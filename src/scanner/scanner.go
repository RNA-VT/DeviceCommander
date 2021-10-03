package scanner

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/device"
)

func getScannerLogger() *log.Entry {
	return log.WithFields(log.Fields{"module": "scanner"})
}

type DeviceResponse struct {
	Success bool
	Device  device.Device
}

func ScanNetwork(addressRoot string) ([]device.Device, error) {
	logger := getScannerLogger()
	deviceList := []device.Device{}
	ipRange := 255

	ch := make(chan DeviceResponse)
	for i := 1; i < ipRange; i++ {
		host := addressRoot + strconv.Itoa(i)
		go ProbeHostConcurrent(host, ch)
	}

	for i := 1; i < ipRange; i++ {
		tmpResponse := <-ch
		if tmpResponse.Success {
			deviceList = append(deviceList, tmpResponse.Device)
		}
	}

	logger.Debug(fmt.Sprintf("Network Scan Results: %+v", deviceList))

	return deviceList, nil
}

func ProbeHostConcurrent(host string, ch chan<- DeviceResponse) {
	success := true
	device, err := ProbeHost(host)
	if err != nil {
		log.Trace(err)
		success = false
	}
	ch <- DeviceResponse{
		Success: success,
		Device:  device,
	}
}

func ProbeHost(host string) (device.Device, error) {
	logger := getScannerLogger()
	url := "http://" + host + "/registration"
	logger.Trace("Probing ", host)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return device.Device{}, err
	}

	switch resp.StatusCode {
	case 200:
		successLogger := logger.WithFields(log.Fields{
			"event": "success",
		})
		dev, err := device.NewDeviceFromRequestBody(resp.Body)
		if err != nil {
			return dev, err
		}

		successLogger.Info(fmt.Sprintf("Response from %s accepted", host))

		return dev, nil
	case 404:
		logger.Debug("Host Not Found: " + host)
		return device.Device{}, errors.New("host not found " + host)
	default:
		logger.Debug("Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
		return device.Device{}, fmt.Errorf("attempt to register %s resulted in an unexpected response: %d", host, resp.StatusCode)
	}
}

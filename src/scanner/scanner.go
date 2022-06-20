package scanner

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rna-vt/devicecommander/src/device"
)

type DeviceResponse struct {
	Success bool
	Device  device.NewDeviceParams
}

type IPScanResults struct {
	IPv4Addresses []net.IP
	IPv6Addresses []net.IP
}

func GetLocalAddresses() (IPScanResults, error) {
	logger := getScannerLogger()
	results := IPScanResults{
		IPv4Addresses: []net.IP{},
		IPv6Addresses: []net.IP{},
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		logger.Error(fmt.Errorf("localAddresses: %w", err))
		return results, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			logger.Error(fmt.Errorf("localAddresses: %w", err))
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				logger.Trace(fmt.Sprintf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask()))

			case *net.IPNet:
				logger.Trace(fmt.Sprintf("%s : %v [%v/%v]\n", i.Name, v, v.IP, v.Mask))
				if v.IP.To4() != nil {
					results.IPv4Addresses = append(results.IPv4Addresses, v.IP)
				} else {
					results.IPv6Addresses = append(results.IPv6Addresses, v.IP)
				}
			}
		}
	}
	return results, nil
}

func ScanIPs(ipSet []net.IP, timeoutSeconds int) ([]device.NewDeviceParams, error) {
	logger := getScannerLogger()
	deviceList := []device.NewDeviceParams{}

	logger.Info("Scan IPs: ", ipSet)

	ch := make(chan DeviceResponse)
	for _, ip := range ipSet {
		go ProbeHostConcurrent(ip.String(), ch, timeoutSeconds)
	}

	for i := 1; i < len(ipSet); i++ {
		tmpResponse := <-ch
		if tmpResponse.Success {
			deviceList = append(deviceList, tmpResponse.Device)
		}
	}

	logger.Debug(fmt.Sprintf("Network Scan Results: %+v", deviceList))

	return deviceList, nil
}

func ProbeHostConcurrent(host string, ch chan<- DeviceResponse, timeoutSeconds int) {
	success := true
	device, err := ProbeHost(host, timeoutSeconds)
	if err != nil {
		log.Trace(err)
		success = false
	}
	ch <- DeviceResponse{
		Success: success,
		Device:  device,
	}
}

func ProbeHost(host string, timeoutSeconds int) (device.NewDeviceParams, error) {
	logger := getScannerLogger()
	url := "http://" + host + "/registration"
	logger.Trace("Probing ", host)

	client := http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return device.NewDeviceParams{}, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		successLogger := logger.WithFields(log.Fields{
			"event": "success",
		})
		dev, err := device.NewDeviceFromRequestBody(resp.Body)
		if err != nil {
			return dev, err
		}

		successLogger.Info(fmt.Sprintf("Response from %s accepted", host))

		return dev, nil
	case http.StatusNotFound:
		logger.Debug("Host Not Found: " + host)
		return device.NewDeviceParams{}, errors.New("host not found " + host)
	default:
		logger.Debug("Attempt to register " + host + " resulted in an unexpected response:" + strconv.Itoa(resp.StatusCode))
		return device.NewDeviceParams{}, fmt.Errorf("attempt to register %s resulted in an unexpected response: %d", host, resp.StatusCode)
	}
}

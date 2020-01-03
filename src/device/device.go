package device

import (
	"strconv"
)

//Device Information/Metadata about device
type Device struct {
	ID   int
	Host string
	Port string
}

/*String Just for pretty printing the device info */
func (device *Device) String() string {
	//strings.Split(myIP[0].String(), "/")[0]
	return "Device:{ deviceId:" + strconv.Itoa(device.ID) + ", Host:" + device.Host + ", port:" + device.Port + " }"
}

//ToFullAddress returns a network address including the ip address and port that this device is listening on
func (device *Device) ToFullAddress() string {
	/* Just for pretty printing the device info */
	return device.Host + ":" + device.Port
}

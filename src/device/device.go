package device

import (
	"firecontroller/component"
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

//Device Information/Metadata about a device
type Device struct {
	ID        int
	Host      string
	Port      string
	Solenoids []component.Solenoid
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

//LoadSolenoids - Load Solenoid Array from Config
func (device *Device) LoadSolenoids() error {
	yamlFile, err := ioutil.ReadFile("/Users/tushar/fun/GoFire/src/app/config/solenoids.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &device.Solenoids)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}

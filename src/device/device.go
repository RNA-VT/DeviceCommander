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
	Solenoids []component.Solenoid `yaml:"solenoids"`
}

/*String Just for pretty printing the device info */
func (device *Device) String() string {
	//strings.Split(myIP[0].String(), "/")[0]
	return "Device:{ deviceId:" + strconv.Itoa(device.ID) + ", Host:" + device.Host + ", port:" + device.Port + " }"
}

//Init - Initialize all the components on this device
func (device *Device) Init() error {
	for i := 0; i < len(device.Solenoids); i++ {
		err := device.Solenoids[i].Init()
		if err != nil {
			return err
		}
		return nil
	}
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
	err = yaml.Unmarshal(yamlFile, &device)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	log.Println("Loaded the Following Solenoids: ", device.solenoidsString())
	log.Print("Performing Initializations: ")
	err = device.Init()
	if err != nil {
		return err
	}
	return nil
}

//SolenoidsString assembles a string of all the Solenoids on this device.
func (device *Device) solenoidsString() string {
	out := ""
	for i := 0; i < len(device.Solenoids); i++ {
		out += device.Solenoids[i].String()
	}
	return out
}

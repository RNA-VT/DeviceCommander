package device

import (
	"firecontroller/utilities"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

//Device represents the machine running the firecontroller and the micros connected to it. Currently only Raspberry Pis are supported
type Device struct {
	ID          int
	Name        string
	Description string
	Host        string
	Port        string
}

//Config -
type Config struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

//GetConfig -
func (m Device) GetConfig() (config Config) {
	config.ID = m.ID
	config.Host = m.Host
	config.Port = m.Port
	config.Name = m.Name
	config.Description = m.Description
	return
}

//Load -
func (m *Device) Load(config Config) {
	m.ID = config.ID
	m.Name = config.Name
	m.Description = config.Description
	m.Host = config.Host
	m.Port = config.Port
}

//NewDevice -
func NewDevice(host string, port string) (Device, error) {
	micro := Device{
		Host: host,
		Port: port,
	}
	err := micro.LoadConfigFromFile()
	if err != nil {
		return Device{}, err
	}
	return micro, nil
}

//LoadConfigFromFile - Load Solenoid Array from config
func (m *Device) LoadConfigFromFile() error {
	//TODO: replace this with a sqlite store
	yamlFile, err := ioutil.ReadFile("./app/config/device.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	var microConfig Config
	err = yaml.Unmarshal(yamlFile, &microConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	microConfig.Host = viper.GetString("HOST")
	microConfig.Port = viper.GetString("PORT")
	m.Load(microConfig)

	if err != nil {
		log.Printf("Error During Initialization: #%v ", err)
		return err
	}
	return nil
}

//ToFullAddress returns a network address including the ip address and port that this micro is listening on
func (m Device) ToFullAddress() string {
	/* Just for pretty printing the micro info */
	return m.Host + ":" + m.Port
}

/*String Just for pretty printing the Device info */
func (m Device) String() string {
	return utilities.LabelString("Device",
		utilities.LabelString("Id", strconv.Itoa(m.ID))+
			utilities.LabelString("Host", m.Host)+
			utilities.LabelString("Port", m.Port))
}

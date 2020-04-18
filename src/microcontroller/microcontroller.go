package microcontroller

import (
	"errors"
	"firecontroller/component"
	"firecontroller/utilities"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

//Microcontroller represents the machine running the firecontroller and the micros connected to it. Currently only Raspberry Pis are supported
type Microcontroller struct {
	ID          int
	Name        string
	Description string
	Host        string
	Port        string
	Master      bool
	Solenoids   []component.Solenoid
	Igniters    []component.Igniter
}

//Config -
type Config struct {
	ID          int                        `yaml:"id"`
	Name        string                     `yaml:"name"`
	Description string                     `yaml:"description"`
	Host        string                     `yaml:"host"`
	Port        string                     `yaml:"port"`
	Master      bool                       `yaml:"master"`
	Solenoids   []component.SolenoidConfig `yaml:"solenoids"`
	Igniters    []component.IgniterConfig  `yaml:"igniters"`
}

//GetConfig -
func (m Microcontroller) GetConfig() (config Config) {
	config.ID = m.ID
	config.Host = m.Host
	config.Port = m.Port
	config.Name = m.Name
	config.Description = m.Description
	config.Master = m.Master
	config.Solenoids = make([]component.SolenoidConfig, len(m.Solenoids))
	for i, sol := range m.Solenoids {
		config.Solenoids[i] = sol.GetConfig()
	}
	config.Igniters = make([]component.IgniterConfig, len(m.Igniters))
	for i, igniter := range m.Igniters {
		config.Igniters[i] = igniter.GetConfig()
	}
	return
}

//Load -
func (m *Microcontroller) Load(config Config) {
	m.ID = config.ID
	m.Name = config.Name
	m.Description = config.Description
	m.Master = config.Master
	if viper.GetBool("GOFIRE_MASTER") {
		m.Host = viper.GetString("GOFIRE_MASTER_HOST")
		m.Port = viper.GetString("GOFIRE_MASTER_PORT")
	} else {
		m.Host = viper.GetString("GOFIRE_HOST")
		m.Port = viper.GetString("GOFIRE_PORT")
	}
	if length := len(config.Solenoids); length > 0 {
		m.Solenoids = make([]component.Solenoid, length)
		for i, sol := range config.Solenoids {
			m.Solenoids[i].Load(sol)
		}
	}
	if length := len(config.Igniters); length > 0 {
		m.Igniters = make([]component.Igniter, length)
		for i, igniter := range config.Igniters {
			m.Igniters[i].Load(igniter)
		}
	}
}

/*String Just for pretty printing the Microcontroller info */
func (m Microcontroller) String() string {
	return utilities.LabelString("Microcontroller",
		utilities.LabelString("Id", strconv.Itoa(m.ID))+
			utilities.LabelString("Host", m.Host)+
			utilities.LabelString("Port", m.Port)+
			utilities.LabelString("Components", m.ComponentString()))
}

//Init - Initialize all the components on this micro
func (m *Microcontroller) Init() (err error) {
	for _, component := range m.GetComponentMap() {
		if err = component.Init(); err != nil {
			log.Println("Failed to load: ", component.String())
			return
		}
	}
	return
}

//GetComponentMap - retrieve a map of pointers to all components on this micro
func (m *Microcontroller) GetComponentMap() map[string]component.Component {
	count := len(m.Solenoids) + len(m.Igniters)
	components := make(map[string]component.Component, count)
	for i := 0; i < len(m.Solenoids); i++ {
		components["Solenoid_"+strconv.Itoa(i)] = &m.Solenoids[i]
	}
	for i := 0; i < len(m.Igniters); i++ {
		components["Igniter_"+strconv.Itoa(i)] = &m.Igniters[i]
	}
	return components
}

//GetSolenoid -
func (m *Microcontroller) GetSolenoid(id int) (sol component.Solenoid, err error) {
	for _, sol := range m.Solenoids {
		if sol.UID == id {
			return sol, nil
		}
	}
	return component.Solenoid{}, errors.New("id not found")
}

//GetIgniter -
func (m *Microcontroller) GetIgniter(id int) (sol component.Igniter, err error) {
	for _, igniter := range m.Igniters {
		if igniter.UID == id {
			return igniter, nil
		}
	}
	return component.Igniter{}, errors.New("id not found")
}

//NewMicrocontroller -
func NewMicrocontroller(host string, port string) (Microcontroller, error) {
	micro := Microcontroller{
		Host: host,
		Port: port,
	}
	err := micro.LoadConfigFromFile()
	if err != nil {
		return Microcontroller{}, err
	}
	return micro, nil

}

//ToFullAddress returns a network address including the ip address and port that this micro is listening on
func (m Microcontroller) ToFullAddress() string {
	/* Just for pretty printing the micro info */
	return m.Host + ":" + m.Port
}

//LoadConfigFromFile - Load Solenoid Array from config
func (m *Microcontroller) LoadConfigFromFile() error {
	yamlFile, err := ioutil.ReadFile("./app/config/microcontroller.yaml")
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
	m.Load(microConfig)
	log.Println("Loaded the Following Components: ", m.ComponentString())
	log.Print("Performing Initializations: ")
	err = m.Init()
	if err != nil {
		log.Printf("Error During Initialization: #%v ", err)
		return err
	}
	return nil
}

//ComponentString assembles a string of all the Solenoids on this microcontroller.
func (m Microcontroller) ComponentString() string {
	out := ""
	for name, component := range m.GetComponentMap() {
		out += "\n[" + name + "]:" + component.String()
	}
	return out
}

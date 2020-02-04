package microcontroller

import (
	"firecontroller/component"
	"firecontroller/utilities"
	"io/ioutil"
	"log"
	"strconv"

	"gopkg.in/yaml.v2"
)

//Microcontroller represents the machine running the firecontroller and the micros connected to it. Currently only Raspberry Pis are supported
type Microcontroller struct {
	ID        int
	Host      string
	Port      string
	Solenoids []component.Solenoid `yaml:"solenoids"`
}

/*String Just for pretty printing the Microcontroller info */
func (m *Microcontroller) String() string {
	return utilities.LabelString("Microcontroller", utilities.LabelString("Id", strconv.Itoa(m.ID))+utilities.LabelString("Host", m.Host)+utilities.LabelString("Port", m.Port)+utilities.LabelString("Solenoids", m.solenoidsString()))
}

//Init - Initialize all the components on this micro
func (m *Microcontroller) Init() error {
	for i := 0; i < len(m.Solenoids); i++ {
		err := m.Solenoids[i].Init()
		if err != nil {
			log.Println("Failed to load: ", m.Solenoids[i].String())
			return err
		}
	}
	return nil
}

//ToFullAddress returns a network address including the ip address and port that this micro is listening on
func (m *Microcontroller) ToFullAddress() string {
	/* Just for pretty printing the micro info */
	return m.Host + ":" + m.Port
}

//LoadSolenoids - Load Solenoid Array from Config
func (m *Microcontroller) LoadSolenoids() error {
	yamlFile, err := ioutil.ReadFile("./app/config/solenoids.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	log.Println("Loaded the Following Solenoids: ", m.solenoidsString())
	log.Print("Performing Initializations: ")
	err = m.Init()
	if err != nil {
		log.Printf("Error During Initialization: #%v ", err)
		return err
	}
	return nil
}

//SolenoidsString assembles a string of all the Solenoids on this microcontroller.
func (m *Microcontroller) solenoidsString() string {
	out := ""
	for i := 0; i < len(m.Solenoids); i++ {
		out += m.Solenoids[i].String()
	}
	return out
}

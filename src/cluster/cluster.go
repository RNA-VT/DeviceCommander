package cluster

import (
	"errors"
	"firecontroller/component"
	"firecontroller/microcontroller"
	mc "firecontroller/microcontroller"
	"firecontroller/utilities"
	"log"
	"strconv"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of microcontrollers
type Cluster struct {
	Name                  string
	SlaveMicrocontrollers []mc.Microcontroller
	Master                mc.Microcontroller
	Me                    *mc.Microcontroller
}

//Config -
type Config struct {
	Name                  string `yaml:"Name"`
	SlaveMicrocontrollers []mc.Config
	Master                mc.Config
}

//GetConfig -
func (c Cluster) GetConfig() (config Config) {
	config.Name = c.Name
	config.Master = c.Master.GetConfig()
	config.SlaveMicrocontrollers = make([]mc.Config, len(c.SlaveMicrocontrollers))
	for i, micro := range c.SlaveMicrocontrollers {
		config.SlaveMicrocontrollers[i] = micro.GetConfig()
	}
	return
}

//Load -
func (c *Cluster) Load(config Config) {
	c.Name = config.Name
	c.Master.Load(config.Master)
	c.SlaveMicrocontrollers = make([]mc.Microcontroller, len(config.SlaveMicrocontrollers))
	for i, micro := range config.SlaveMicrocontrollers {
		c.SlaveMicrocontrollers[i].Load(micro)
	}
}

func (c Cluster) String() string {
	cluster, err := utilities.StringJSON(c)
	if err != nil {
		return ""
	}
	return cluster
}

//Start registers this microcontroller, retrieves cluster config, loads local components and verifies peers
func (c *Cluster) Start() {
	//Set global ref to cluster
	gofireMaster := viper.GetBool("GOFIRE_MASTER")
	if gofireMaster {
		log.Println("Master Mode Enabled!")
		c.KingMe()
	} else {
		log.Println("Slave Mode Enabled.")
		c.ALifeOfServitude()
	}
}

//GetMicrocontrollers returns a map[microcontrollerID]microcontroller of all Microcontrollers in the cluster
func (c Cluster) GetMicrocontrollers() map[int]microcontroller.Microcontroller {
	micros := make(map[int]microcontroller.Microcontroller)
	for i := 0; i < len(c.SlaveMicrocontrollers); i++ {
		micros[c.SlaveMicrocontrollers[i].ID] = c.SlaveMicrocontrollers[i]
	}
	return micros
}

//GetComponent - gets a component by its id
func (c *Cluster) GetComponent(id string) (sol component.Solenoid, err error) {
	components := c.GetComponents()
	sol, ok := components[id]
	if !ok {
		return sol, errors.New("Component Not Found")
	}
	return sol, nil
}

//GetComponents builds a map of all the components in the cluster by a cluster wide unique key
func (c Cluster) GetComponents() map[string]component.Solenoid {
	components := make(map[string]component.Solenoid, c.countComponents())
	for i := 0; i < len(c.SlaveMicrocontrollers); i++ {
		for j := 0; j < len(c.SlaveMicrocontrollers[i].Solenoids); j++ {
			key := strconv.Itoa(c.SlaveMicrocontrollers[i].Solenoids[j].UID)
			components[key] = c.SlaveMicrocontrollers[i].Solenoids[j]
		}
	}
	return components
}
func (c Cluster) countComponents() int {
	count := 0
	for i := 0; i < len(c.SlaveMicrocontrollers); i++ {
		count += len(c.SlaveMicrocontrollers[i].Solenoids)
	}

	return count
}

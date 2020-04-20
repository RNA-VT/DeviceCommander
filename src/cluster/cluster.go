package cluster

import (
	"firecontroller/microcontroller"
	mc "firecontroller/microcontroller"
	"firecontroller/utilities"
	"log"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of microcontrollers
type Cluster struct {
	Name             string
	Microcontrollers []mc.Microcontroller
	Me               *mc.Microcontroller
}

//Config -
type Config struct {
	Name             string `yaml:"Name"`
	Microcontrollers []mc.Config
}

//GetConfig -
func (c Cluster) GetConfig() (config Config) {
	config.Name = c.Name
	config.Microcontrollers = make([]mc.Config, len(c.Microcontrollers))
	for i, micro := range c.Microcontrollers {
		config.Microcontrollers[i] = micro.GetConfig()
	}

	return
}

//Master - returns a pointer to the Master micro
func (c *Cluster) Master() *mc.Microcontroller {
	for _, micro := range c.Microcontrollers {
		if micro.Master {
			return &micro
		}
	}
	return &mc.Microcontroller{}
}

//Load -
func (c *Cluster) Load(config Config) {
	c.Name = config.Name
	newPeers := []mc.Microcontroller{
		*c.Me,
	}
	//c.Microcontrollers = make([]mc.Microcontroller, len(config.Microcontrollers))
	for _, micro := range config.Microcontrollers {
		if c.Me.ID != micro.ID {
			newPeers = append(newPeers, mc.Microcontroller{})
			newPeers[len(newPeers)-1].Load(micro)
		}
	}
	c.Microcontrollers = newPeers
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
	for i := 0; i < len(c.Microcontrollers); i++ {
		micros[c.Microcontrollers[i].ID] = c.Microcontrollers[i]
	}
	return micros
}

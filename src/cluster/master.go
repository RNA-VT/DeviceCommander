package cluster

import (
	"errors"
	mc "firecontroller/microcontroller"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

//******************************************************************************************************
//*******Master Only Methods****************************************************************************
//******************************************************************************************************

//KingMe makes this microcontroller the master
func (c *Cluster) KingMe() {
	me, err := mc.NewMicrocontroller(viper.GetString("GOFIRE_MASTER_HOST"), viper.GetString("GOFIRE_MASTER_PORT"))
	if err != nil {
		log.Println("Failed to Create New Microcontroller:", err.Error())
	}
	me.ID = c.generateUniqueID()
	me.Master = true
	//The master also serves
	c.Microcontrollers = append(c.Microcontrollers, me)
	c.Me = c.Master()

	//The Master pulls out its stethoscope ...
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for {
			select {
			case t := <-ticker.C:
				log.Println("Begin Heartbeat Check", t)
				for _, m := range c.Microcontrollers {
					if !m.Master {
						log.Println("Checking Peer:", m.Name, m.ToFullAddress())
						url := "http://" + m.ToFullAddress() + "/v1/health"
						resp, err := http.Get(url)
						if err != nil || resp.StatusCode != 200 {
							log.Println(m.Name + " @" + m.ToFullAddress() + " is NOT ok")
							log.Println("Deregistering Microcontroller...")
							c.RemoveMicrocontroller(m)
							c.SendClusterUpdate([]mc.Config{})
						} else {
							log.Println(m.Name + " @" + m.ToFullAddress() + " is ok")
						}
					}
				}
			}
		}
	}()
}

//AddMicrocontroller attempts to add a microcontroller to the cluster and returns the response data. This should only be run by the master.
func (c *Cluster) AddMicrocontroller(newMC mc.Config) error {
	var newGuy mc.Microcontroller
	newGuy.Load(newMC)
	if viper.GetString("ENV") == "production" {
		for _, micro := range c.Microcontrollers {
			if micro.Host == newGuy.Host {
				//This guy ain't so new!
				return errors.New("Requesting instance is running on a microcontroller already registered to this cluster")
			}
		}
	}

	c.Microcontrollers = append(c.Microcontrollers, newGuy)

	PrintClusterInfo(*c)
	return nil
}

//RemoveMicrocontroller -
func (c *Cluster) RemoveMicrocontroller(ImDoneHere mc.Microcontroller) {
	for index, mc := range c.Microcontrollers {
		if mc.ID == ImDoneHere.ID {
			s := c.Microcontrollers
			count := len(c.Microcontrollers)
			s[count-1], s[index] = s[index], s[count-1]
			c.Microcontrollers = s[:len(s)-1]
			return
		}
	}
}

//SendClusterUpdate -
func (c Cluster) SendClusterUpdate(exclude []mc.Config) error {
	msg := PeerUpdateMessage{
		Cluster: c.GetConfig(),
		Header:  c.GetHeader(),
	}

	exclusions := append(exclude, c.Me.GetConfig())

	err := c.UpdatePeers("/peers", msg, exclusions)
	if err != nil {
		log.Println("Unexpected Error during attempt to contact all peers: ", err)
		return err
	}
	return nil
}

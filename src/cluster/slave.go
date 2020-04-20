package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	mc "firecontroller/microcontroller"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

//******************************************************************************************************
//*******Slave Only Methods*****************************************************************************
//******************************************************************************************************

//ALifeOfServitude is all that awaits this microcontroller
func (c *Cluster) ALifeOfServitude() {
	me, err := mc.NewMicrocontroller(viper.GetString("GOFIRE_HOST"), viper.GetString("GOFIRE_PORT"))
	if err != nil {
		log.Println("Failed to Create New Microcontroller:", err.Error())
	}
	me.ID = c.generateUniqueID()
	me.Master = false
	c.Me = &me
	masterHostname := viper.GetString("GOFIRE_MASTER_HOST") + ":" + viper.GetString("GOFIRE_MASTER_PORT")
	//Try and Connect to the Master
	err = test(masterHostname)
	if err != nil {
		log.Println("Failed to Reach Master Microcontroller: PANIC")
		//TODO: Add Retry or failover maybe? panic for now
		panic(err)
	}
	err = c.JoinNetwork(masterHostname)
	if err != nil {
		log.Println("Failed to Join Network: PANIC")
		panic(err)
	}
}

// JoinNetwork checks if the master exists and joins the network
func (c *Cluster) JoinNetwork(URL string) error {
	parsedURL, err := url.Parse("http://" + URL + "/v1/join_network")
	log.Println("Trying to Join: " + parsedURL.String())

	msg := JoinNetworkMessage{
		ImNewHere: c.Me.GetConfig(),
		Header:    c.GetHeader(),
	}
	body, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to create json message body")
		return err
	}
	resp, err := http.Post(parsedURL.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("Something went wrong connecting to the Master", c.Me)
		log.Println(err)
		return err
	} else if resp.StatusCode >= 400 {
		return errors.New("Registration request was rejected by the Master")
	}
	log.Println("Registration request sent to the Master successfully.")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var t PeerUpdateMessage
	err = decoder.Decode(&t)
	if err != nil {
		log.Println("Failed to decode response from Master Microcontroller")
		log.Println(err)
		return err
	}
	//Update self with data from the master
	c.Load(t.Cluster)

	return nil
}

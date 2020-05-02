package cluster

import (
	"bytes"
	"encoding/json"
	mc "firecontroller/microcontroller"
	"firecontroller/utilities"
	"log"
	"net/http"
	"time"
)

//PeerErrorMessage -
type PeerErrorMessage struct {
	Panic        bool
	DeregisterMe mc.Microcontroller
	PeerInfoMessage
}

//PeerInfoMessage -
type PeerInfoMessage struct {
	Messages []string
	Header   GoFireHeader
}

//JoinNetworkMessage is the registration request
type JoinNetworkMessage struct {
	ImNewHere mc.Config
	Header    GoFireHeader
}

//PeerUpdateMessage contains a source and cluster info
type PeerUpdateMessage struct {
	Cluster Config
	Header  GoFireHeader
}

//CommandMessage -
type CommandMessage struct {
	Command       string
	ComponentType string
}

//GoFireHeader -
type GoFireHeader struct {
	Source  mc.Config
	Created time.Time
}

//GetHeader -
func (c Cluster) GetHeader() GoFireHeader {
	return GoFireHeader{
		Source:  c.Me.GetConfig(),
		Created: time.Now(),
	}
}

//ClusterError - Meant for Errors that should stop the entire cluster or deregister this micro from the cluster
func (c Cluster) ClusterError(panicAfterWarning bool, panicCluster bool, MicrocontrollerToRemove mc.Microcontroller, notGoodThings ...string) {
	var message PeerErrorMessage
	message.Messages = notGoodThings
	message.Panic = panicCluster
	message.DeregisterMe = MicrocontrollerToRemove
	message.Header = c.GetHeader()
	c.UpdatePeers("/errors", message, []mc.Config{c.Me.GetConfig()})
	if panicAfterWarning {
		panic(notGoodThings)
	}
}

// UpdatePeers will take a byte slice and POST it to each microcontroller

func (c Cluster) UpdatePeers(urlPath string, message interface{}, exclude []mc.Config) error {
	for i := 0; i < len(c.Microcontrollers); i++ {
		if !isExcluded(c.Microcontrollers[i], exclude) {
			body, err := utilities.JSON(message)
			if err != nil {
				log.Println("Failed to convert cluster to json: ", c)
				return err
			}
			currURL := "http://" + c.Microcontrollers[i].ToFullAddress() + "/v1" + urlPath

			resp, err := http.Post(currURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				log.Println("WARNING: Failed to POST to Peer: ", c.Microcontrollers[i].String(), currURL)
				log.Println(err)
			} else {
				defer resp.Body.Close()
				var result string
				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&result)
				log.Println("Peer Update Response:", result)
			}
		}
	}
	return nil
}

//ReceiveError -
func (c *Cluster) ReceiveError(msg PeerErrorMessage) {
	//log msgs to console
	for msg := range msg.Messages {
		log.Println(msg)
	}
	if msg.Panic {
		panic(map[string]interface{}{
			"Cluster": c,
			"Message": msg,
		})
	}
	// TODO do better with this check
	if msg.DeregisterMe.Host != "" {
		//Deregister Microcontroller
		log.Println("Deregistering Microcontroller From Cluster: ", msg.DeregisterMe.String())
		c.RemoveMicrocontroller(msg.DeregisterMe)
		c.SendClusterUpdate([]mc.Config{})
	}
}

package cluster

import (
	"bytes"
	"encoding/json"
	mc "firecontroller/microcontroller"
	"log"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of control devices
type Cluster struct {
	Name         string
	SlaveDevices []mc.Microcontroller
	MasterDevice mc.Microcontroller
	Me           mc.Microcontroller
}

//Start registers this device, retrieves cluster config, loads local components and verifies peers
func (c *Cluster) Start() {
	gofireMaster := viper.GetBool("GOFIRE_MASTER")
	if gofireMaster {
		log.Println("Master Mode Enabled!")
		c.KingMe()
	} else {
		log.Println("Slave Mode Enabled.")
		c.ALifeOfServitude()
	}
}

// UpdatePeers will take a byte slice and POST it to each device
func (c *Cluster) UpdatePeers(urlPath string, message PeerUpdateMessage, excludeDevices []mc.Microcontroller) error {
	for i := 0; i < len(c.SlaveDevices); i++ {
		if !isExcluded(c.SlaveDevices[i], excludeDevices) {
			body, err := json.Marshal(message)
			if err != nil {
				log.Println("Failed to convert cluster to json: ", c)
				return err
			}
			currURL := "http://" + c.SlaveDevices[i].ToFullAddress() + urlPath

			resp, err := http.Post(currURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				log.Println("WARNING: Failed to POST to Peer: ", c.SlaveDevices[i].String(), currURL)
				log.Println(err)
			} else {
				defer resp.Body.Close()
				var result string
				decoder := json.NewDecoder(resp.Body)
				decoder.Decode(&result)
				log.Println("Result:", result)
			}
		}
	}
	return nil
}

//NewDevice -
func (c *Cluster) NewDevice(host string, port string) (mc.Microcontroller, error) {
	dvc := mc.Microcontroller{
		ID:   c.generateUniqueID(),
		Host: host,
		Port: port,
	}
	err := dvc.LoadSolenoids()
	if err != nil {
		return mc.Microcontroller{}, err
	}
	return dvc, nil

}

//******************************************************************************************************
//*******Master Only Methods****************************************************************************
//******************************************************************************************************

//KingMe makes this device the master
func (c *Cluster) KingMe() {
	me, err := c.NewDevice(viper.GetString("GOFIRE_MASTER_HOST"), viper.GetString("GOFIRE_MASTER_PORT"))
	if err != nil {
		log.Println("Failed to Create New Device:", err.Error())
	}
	c.Me = me
	c.MasterDevice = me
	//The Master waits ...
}

//AddDevice attempts to add a device to the cluster and returns the response data. This should only be run by the master.
func (c *Cluster) AddDevice(newMC mc.Microcontroller) (response PeerUpdateMessage, err error) {
	newMC.ID = c.generateUniqueID()
	c.SlaveDevices = append(c.SlaveDevices, newMC)
	c.PrintClusterInfo()

	response = PeerUpdateMessage{
		Source:  c.Me,
		Cluster: *c,
	}

	exclusions := []mc.Microcontroller{newMC, c.Me}
	err = c.UpdatePeers("/", response, exclusions)
	if err != nil {
		log.Println("Unexpected Error during attempt to contact all peers: ", err)
		return PeerUpdateMessage{}, err
	}

	return response, nil
}

//******************************************************************************************************
//*******Slave Only Methods*****************************************************************************
//******************************************************************************************************

//ALifeOfServitude is all that awaits this device
func (c *Cluster) ALifeOfServitude() {
	me, err := c.NewDevice(viper.GetString("GOFIRE_HOST"), viper.GetString("GOFIRE_PORT"))
	if err != nil {
		log.Println("Failed to Create New Device:", err.Error())
	}
	c.Me = me
	masterHostname := viper.GetString("GOFIRE_MASTER_HOST") + ":" + viper.GetString("GOFIRE_MASTER_PORT")
	//Try and Connect to the Master
	err = test(masterHostname)
	if err != nil {
		log.Println("Failed to Reach Master Device: PANIC")
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
	parsedURL, err := url.Parse("http://" + URL + "/join_network")
	log.Println("Trying to Join: " + parsedURL.String())
	msg := JoinNetworkMessage{
		ImNewHere: c.Me,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to create json message body")
		return err
	}
	resp, err := http.Post(parsedURL.String(), "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Println("[test] Couldn't connect to master.", c.Me.ID)
		log.Println(err)
		return err
	}
	log.Println("Connected to master. Sending message to peers.")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var t PeerUpdateMessage
	err = decoder.Decode(&t)
	if err != nil {
		log.Println("Failed to decode response from Master Device")
		log.Println(err)
		return err
	}
	//Update self with data from the master
	c.LoadCluster(t.Cluster)

	return nil
}

//generateUniqueID returns a unique id for asigning to a new device
func (c *Cluster) generateUniqueID() int {
	randID := rand.Intn(100)
	for len(c.getSlavesByID(randID)) > 0 {
		randID = rand.Intn(100)
		log.Println(len(c.getSlavesByID(randID)))
	}
	return randID
}

// getSlaveByID find all the slave for a given ID
func (c *Cluster) getSlavesByID(targetID int) []mc.Microcontroller {
	var devices []mc.Microcontroller

	for i := 0; i < len(c.SlaveDevices); i++ {
		if c.SlaveDevices[i].ID == targetID {
			return append(devices, c.SlaveDevices[i])
		}
	}

	return devices
}

// GetAllSlavesByIP find all slave device by its IP
func (c *Cluster) GetAllSlavesByIP(host string) []mc.Microcontroller {
	var devices []mc.Microcontroller

	for i := 0; i < len(c.SlaveDevices); i++ {
		if c.SlaveDevices[i].Host == host {
			devices = append(devices, c.SlaveDevices[i])
		}
	}

	return devices
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c *Cluster) PrintClusterInfo() {
	log.Println()
	log.Println("====Master====")
	log.Println(c.MasterDevice)

	log.Println()

	for i := 0; i < len(c.SlaveDevices); i++ {
		log.Println("----Device---")
		log.Println(c.SlaveDevices[i])
	}
	log.Println()
}

//LoadCluster sets all Cluster values except for Me
func (c *Cluster) LoadCluster(cluster Cluster) {
	log.Println("Loading Updated Cluster Data...")
	c.Name = cluster.Name
	c.MasterDevice = cluster.MasterDevice
	c.SlaveDevices = cluster.SlaveDevices
	c.PrintClusterInfo()
}

func isExcluded(m mc.Microcontroller, exclusions []mc.Microcontroller) bool {
	for i := 0; i < len(exclusions); i++ {
		if m.Host == exclusions[i].Host && m.Port == exclusions[i].Port {
			return true
		}
	}
	return false
}

// test check if master exists
func test(URL string) error {
	parsedURL, err := url.Parse("http://" + URL)
	if err != nil {
		log.Println("Failed to Parse URL")
		return err
	}
	_, err = http.Get(parsedURL.String())
	return err
}

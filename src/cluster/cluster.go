package cluster

import (
	"bytes"
	"encoding/json"
	"errors"
	"firecontroller/component"
	"firecontroller/microcontroller"
	mc "firecontroller/microcontroller"
	"firecontroller/utilities"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

//Cluster - This object defines an array of microcontrollers
type Cluster struct {
	Name                 string
	SlaveMicrocontrolers []mc.Microcontroller
	Master               mc.Microcontroller
	Me                   mc.Microcontroller
}

func (c *Cluster) String() string {
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

// UpdatePeers will take a byte slice and POST it to each microcontroller
func (c *Cluster) UpdatePeers(urlPath string, message interface{}, exclude []mc.Microcontroller) error {
	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		if !isExcluded(c.SlaveMicrocontrolers[i], exclude) {
			body, err := utilities.JSON(message)
			if err != nil {
				log.Println("Failed to convert cluster to json: ", c)
				return err
			}
			currURL := "http://" + c.SlaveMicrocontrolers[i].ToFullAddress() + urlPath

			resp, err := http.Post(currURL, "application/json", bytes.NewBuffer(body))
			if err != nil {
				log.Println("WARNING: Failed to POST to Peer: ", c.SlaveMicrocontrolers[i].String(), currURL)
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

//NewMicrocontroller -
func (c *Cluster) NewMicrocontroller(host string, port string) (mc.Microcontroller, error) {
	micro := mc.Microcontroller{
		ID:   c.generateUniqueID(),
		Host: host,
		Port: port,
	}
	err := micro.LoadSolenoids()
	if err != nil {
		return mc.Microcontroller{}, err
	}
	return micro, nil

}

//GetMicrocontrollers returns a map[microcontrollerID]microcontroller of all Microcontrollers in the cluster
func (c *Cluster) GetMicrocontrollers() map[int]microcontroller.Microcontroller {
	micros := make(map[int]microcontroller.Microcontroller)
	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		micros[c.SlaveMicrocontrolers[i].ID] = c.SlaveMicrocontrolers[i]
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
func (c *Cluster) GetComponents() map[string]component.Solenoid {
	components := make(map[string]component.Solenoid, c.countComponents())
	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		for j := 0; j < len(c.SlaveMicrocontrolers[i].Solenoids); j++ {
			key := strconv.Itoa(c.SlaveMicrocontrolers[i].Solenoids[j].UID)
			components[key] = c.SlaveMicrocontrolers[i].Solenoids[j]
		}
	}
	return components
}
func (c *Cluster) countComponents() int {
	count := 0
	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		count += len(c.SlaveMicrocontrolers[i].Solenoids)
	}

	return count
}

//******************************************************************************************************
//*******Master Only Methods****************************************************************************
//******************************************************************************************************

//KingMe makes this microcontroller the master
func (c *Cluster) KingMe() {
	me, err := c.NewMicrocontroller(viper.GetString("GOFIRE_MASTER_HOST"), viper.GetString("GOFIRE_MASTER_PORT"))
	if err != nil {
		log.Println("Failed to Create New Microcontroller:", err.Error())
	}
	c.Me = me
	c.Master = me
	//The master also serves
	c.SlaveMicrocontrolers = append(c.SlaveMicrocontrolers, me)
	//The Master waits ...
}

//AddMicrocontroller attempts to add a microcontroller to the cluster and returns the response data. This should only be run by the master.
func (c *Cluster) AddMicrocontroller(newMC mc.Microcontroller) (response PeerUpdateMessage, err error) {
	newMC.ID = c.generateUniqueID()
	c.SlaveMicrocontrolers = append(c.SlaveMicrocontrolers, newMC)
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

//ALifeOfServitude is all that awaits this microcontroller
func (c *Cluster) ALifeOfServitude() {
	me, err := c.NewMicrocontroller(viper.GetString("GOFIRE_HOST"), viper.GetString("GOFIRE_PORT"))
	if err != nil {
		log.Println("Failed to Create New Microcontroller:", err.Error())
	}
	c.Me = me
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
		log.Println("Failed to decode response from Master Microcontroller")
		log.Println(err)
		return err
	}
	//Update self with data from the master
	c.LoadCluster(t.Cluster)

	return nil
}

//generateUniqueID returns a unique id for asigning to a new microcontroller
func (c *Cluster) generateUniqueID() int {
	limit := viper.GetInt("MICROCONTORLLER_LIMIT")
	randID := rand.Intn(limit)
	for len(c.getSlavesByID(randID)) > 0 {
		randID = rand.Intn(limit)
	}
	return randID
}

// getSlaveByID find all the slave for a given ID
func (c *Cluster) getSlavesByID(targetID int) []mc.Microcontroller {
	var micros []mc.Microcontroller

	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		if c.SlaveMicrocontrolers[i].ID == targetID {
			return append(micros, c.SlaveMicrocontrolers[i])
		}
	}

	return micros
}

// PrintClusterInfo will cleanly print out info about the cluster
func (c *Cluster) PrintClusterInfo() {
	log.Println()
	log.Println("====Master====")
	log.Println(c.Master)

	log.Println()

	for i := 0; i < len(c.SlaveMicrocontrolers); i++ {
		log.Println("----Peer---")
		log.Println(c.SlaveMicrocontrolers[i])
	}
	log.Println()
}

//LoadCluster sets all Cluster values except for Me
func (c *Cluster) LoadCluster(cluster Cluster) {
	log.Println("Loading Updated Cluster Data...")
	c.Name = cluster.Name
	c.Master = cluster.Master
	c.SlaveMicrocontrolers = cluster.SlaveMicrocontrolers
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

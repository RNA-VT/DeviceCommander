package app

import (
	"bytes"
	"encoding/json"
	"firecontroller/nodecluster"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

// The Application encapsulates the required state for the running software
type Application struct {
	Cluster nodecluster.Cluster
	Me      nodecluster.NodeInfo
	Echo    *echo.Echo
}

// TestConnectToMaster check if master exists and get assigned id
func (a *Application) TestConnectToMaster(URL string) error {
	parsedURL, err := url.Parse("http://" + URL)
	if err != nil {
		log.Println("Failed to Parse URL")
		return err
	}
	_, err = http.Get(parsedURL.String())
	return err
}

// JoinNetwork check if master exists and get assigned id and port
func (a *Application) JoinNetwork(URL string, node nodecluster.NodeInfo) error {
	parsedURL, err := url.Parse("http://" + URL + "/join_network")
	log.Println("[test] Test Url: " + parsedURL.String())
	msg := nodecluster.JoinNetworkMessage{
		Node: a.Me,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to create json message body")
		return err
	}
	resp, err := http.Post(parsedURL.String(), "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Println("[test] Couldn't connect to master.", a.Me.NodeID)
		log.Println(err)
		return err
	}
	log.Println("Connected to master. Sending message to node.")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var t nodecluster.AddToClusterMessage
	err = decoder.Decode(&t)
	if err != nil {
		log.Println("Failed to decode response from Master Node")
		log.Println(err)
		return err
	}
	//Update self with data from the master
	a.Me.NodeID = t.Dest.NodeID

	a.Cluster = t.Cluster
	a.Cluster.PrintClusterInfo()

	exclusions := []nodecluster.NodeInfo{t.Dest}
	err = a.Cluster.SendMessageToAllNodes("/", t, exclusions)
	if err != nil {
		log.Println("Unexpected Error during attempt to contact all peers: ", err)
		return err
	}

	return nil
}

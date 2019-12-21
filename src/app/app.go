package app

import (
	"encoding/json"
	"firecontroller/nodecluster"
	"fmt"
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
func (a *Application) TestConnectToMaster(URL string) (resp *http.Response, err error) {
	parsedURL, err := url.Parse("http://" + URL)
	if err != nil {
		fmt.Println("Failed to Parse URL")
		return nil, err
	}
	return http.Get(parsedURL.String())
}

// JoinNetwork check if master exists and get assigned id and port
func (a *Application) JoinNetwork(URL string) (nodecluster.NodeInfo, error) {
	parsedURL, err := url.Parse("http://" + URL + "/join_network")
	fmt.Println("[test] Test Url: " + parsedURL.String())
	resp, err := http.Get(parsedURL.String())

	if err != nil {
		fmt.Println("[test] Couldn't connect to master.", a.Me.NodeID)
		fmt.Println(err)
		return nodecluster.NodeInfo{}, err
	}
	fmt.Println("Connected to master. Sending message to node.")

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var t nodecluster.AddToClusterMessage
	err = decoder.Decode(&t)
	if err != nil {
		fmt.Println("Failed to decode response from Master Node")
		fmt.Println(err)
		return nodecluster.NodeInfo{}, err
	}
	//Update self with data from the master
	a.Me = t.Dest
	a.Cluster = t.Cluster
	a.Cluster.PrintClusterInfo()

	err = a.Cluster.SendMessageToAllNodes("/", t)

	return t.Dest, nil
}

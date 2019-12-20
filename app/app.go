package app

import (
	"encoding/json"
	"firecontroller/nodecluster"
	"fmt"
	"net"
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
func (a *Application) TestConnectToMaster(testIP string) bool {
	masterURL := "http://" + testIP
	parsedURL, err := url.Parse(masterURL)
	fmt.Println("[test] Test Url: " + parsedURL.String())
	resp, err := http.Get(parsedURL.String())

	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("[test] Couldn't connect to cluster.", a.Me.NodeId)
			fmt.Println(err)
		}
	} else {
		fmt.Println(resp)
		return true
	}
	return false
}

// JoinNetwork check if master exists and get assigned id and port
func (a *Application) JoinNetwork(testIP string) (bool, nodecluster.NodeInfo) {
	masterURL := "http://" + testIP + "/join_network"
	parsedURL, err := url.Parse(masterURL)
	fmt.Println("[test] Test Url: " + parsedURL.String())
	resp, err := http.Get(parsedURL.String())

	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("[test] Couldn't connect to cluster.", a.Me.NodeId)
			fmt.Println(err)
		}
	} else {
		fmt.Println("[test] Connected to cluster. Sending message to node.")

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var t nodecluster.AddToClusterMessage
		err = decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		a.Me = t.Dest
		a.Cluster = t.Cluster

		a.Cluster.PrintClusterInfo()

		_ = a.Cluster.SendMessageToAllNodes("/", t)

		return true, t.Dest
	}
	return false, a.Me
}

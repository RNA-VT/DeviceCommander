package app

import (
	"encoding/json"
	"firecontroller/nodecluster"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Application struct {
	Cluster nodecluster.Cluster
	Me      nodecluster.NodeInfo
	Echo    *echo.Echo
}

func (a *Application) ConfigureMasterRoutes() {
	fmt.Println("Configure routes listening on " + a.Me.Port)

	// Middleware
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	// Routes
	a.Echo.GET("/", a.defaultGet)
	a.Echo.GET("/cluster_info", a.getClusterInfo)
	a.Echo.GET("/join_network", a.joinNetwork)

	// Start server
	a.Echo.Logger.Fatal(a.Echo.Start(":" + a.Me.Port))
}

func (a *Application) isMaster() bool {
	if *a.Cluster.MasterIp == a.Me.ToFullAdress() {
		return true
	}
	return false
}

// TestConnectToMaster check if master exists and get assigned id
func (a *Application) TestConnectToMaster() (bool, int) {
	if a.isMaster() {
		return true, a.Me.NodeId
	} else {
		var masterURL string = "http://" + *a.Cluster.MasterIp + "/join_network"
		parsedURL, err := url.Parse(masterURL)
		fmt.Println("[test] Test Url: " + parsedURL.String())
		resp, err := http.Get(parsedURL.String())

		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("[test] Couldn't connect to cluster.", a.Me.NodeId)
				fmt.Println(err)
				return false, a.Me.NodeId
			}
		} else {
			fmt.Println("[test] Connected to cluster. Sending message to node.")

			defer resp.Body.Close()
			// body, err := ioutil.ReadAll(resp.Body)

			// if err != nil {
			// 	fmt.Println(err)
			// }
			decoder := json.NewDecoder(resp.Body)
			var t nodecluster.Cluster
			err = decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			fmt.Println(t)
			// var m Message
			// json.Unmarshal()
			return true, 20
		}
	}
	return false, a.Me.NodeId
}

func (a *Application) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster)
}

func (a *Application) defaultGet(c echo.Context) error {
	fmt.Println("DEFAULT GET")

	return c.JSON(http.StatusOK, a.Cluster.GenerateUniqueID())
}

func (a *Application) joinNetwork(c echo.Context) error {
	fmt.Println("[master] Node asked to join cluster")

	newNode := nodecluster.NodeInfo{
		NodeId:     a.Cluster.GenerateUniqueID(),
		NodeIpAddr: "99.99.99.99",
		Port:       "8002",
	}

	a.Cluster.AddSlaveNode(newNode)

	fmt.Println(newNode)

	return c.JSON(http.StatusOK, a.Cluster)
}

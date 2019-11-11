package app

import (
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

func (a *Application) ConfigureRoutes() {
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

func (a *Application) TestConnectToMaster() bool {
	if a.isMaster() {
		return true
	} else {
		fmt.Println(*a.Cluster.MasterIp + "/join_network")
		var masterUrl string = "http://" + *a.Cluster.MasterIp + "/join_network"
		parsedUrl, err := url.Parse(masterUrl)
		fmt.Println(parsedUrl)
		resp, err := http.Get(parsedUrl.String())

		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Couldn't connect to cluster.", a.Me.NodeId)
				fmt.Println(err)
				return false
			}
		} else {
			fmt.Println("Connected to cluster. Sending message to node.")
			fmt.Println(resp)
			return true
		}
	}
	return false
}

func (a *Application) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster)
}

func (a *Application) defaultGet(c echo.Context) error {
	return c.String(http.StatusOK, "DEFAULT")
}

func (a *Application) joinNetwork(c echo.Context) error {
	fmt.Println("Connected to cluster. Sending message to node.")
	text := "Hi nody.. please add me to the cluster.."

	return c.String(http.StatusOK, "Hello, World!"+text)
}

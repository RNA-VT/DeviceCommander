package app

import (
	"firecontroller/nodecluster"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ConfigureRoutes will use Echo to start listening on the appropriate paths
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

func (a *Application) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster)
}

func (a *Application) defaultGet(c echo.Context) error {
	return c.String(http.StatusOK, "This is a FireController Node")
}

func (a *Application) joinNetwork(c echo.Context) error {
	fmt.Println("[master] Node asked to join cluster")

	newNodeID := a.Cluster.GenerateUniqueID()
	newNodeIP := c.RealIP()
	newNodePort := a.Cluster.GenerateUniquePort(newNodeIP)

	newNode := nodecluster.NodeInfo{
		NodeId:     newNodeID,
		NodeIpAddr: newNodeIP,
		Port:       newNodePort,
	}

	a.Cluster.AddSlaveNode(newNode)
	a.Cluster.PrintClusterInfo()

	message := nodecluster.AddToClusterMessage{
		Source:  a.Me,
		Dest:    newNode,
		Cluster: a.Cluster,
	}

	return c.JSON(http.StatusOK, message)
}

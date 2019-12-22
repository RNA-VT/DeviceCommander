package app

import (
	"encoding/json"
	"firecontroller/nodecluster"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ConfigureRoutes will use Echo to start listening on the appropriate paths
func (a *Application) ConfigureRoutes(listenURL string) {

	// Middleware
	a.Echo.Use(middleware.Logger())
	a.Echo.Use(middleware.Recover())

	// Routes
	a.Echo.GET("/", a.defaultGet)
	a.Echo.POST("/", a.peerUpdate)
	a.Echo.GET("/cluster_info", a.getClusterInfo)
	a.Echo.POST("/join_network", a.joinNetwork)

	log.Println("Configure routes listening on " + listenURL)

	log.Println("***************************************")
	log.Println("~Rejoice~ GoFire Lives Again! ~Rejoice~")
	log.Println("***************************************")

	// Start server
	a.Echo.Logger.Fatal(a.Echo.Start(listenURL))

}

func (a *Application) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster)
}

func (a *Application) defaultGet(c echo.Context) error {
	log.Println("Someone is touching me")

	return c.String(http.StatusOK, "This is a FireController Node")
}

func (a *Application) joinNetwork(c echo.Context) error {
	log.Println("[master] Node asked to join cluster")

	body := c.Request().Body
	decoder := json.NewDecoder(body)
	var msg nodecluster.JoinNetworkMessage
	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}
	requestingNode := msg.Node
	requestingNode.NodeID = a.Cluster.GenerateUniqueID()

	a.Cluster.AddSlaveNode(requestingNode)

	a.Cluster.PrintClusterInfo()

	message := nodecluster.AddToClusterMessage{
		Source:  a.Me,
		Dest:    requestingNode,
		Cluster: a.Cluster,
	}

	return c.JSON(http.StatusOK, message)
}

//PeerUpdate receives new cluster info from the most recently registered peer
func (a *Application) peerUpdate(c echo.Context) error {
	log.Println("Receiving Update from New Peer")
	body := c.Request().Body

	var clustah nodecluster.Cluster
	err := json.NewDecoder(body).Decode(&clustah)
	if err != nil {
		log.Println("Failed to decode Cluster info from new peer")
		//TODO: Add Cluster Info Request to repair Cluster info
		return err
	}
	//TODO: Verify my presence in SlaveList
	//TODO: Verify my Master state
	//TODO: Inform Master of Bad Config

	//Update my cluster
	a.Cluster = clustah
	return c.JSON(http.StatusOK, "Peer Update Successfully Received!")
}

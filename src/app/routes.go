package app

import (
	"encoding/json"
	"firecontroller/cluster"
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

	return c.String(http.StatusOK, "Help Me! I'm trapped in the Server! You're the only one receiving this message.")
}

func (a *Application) joinNetwork(c echo.Context) error {
	log.Println("[master] Device asked to join cluster")

	body := c.Request().Body
	decoder := json.NewDecoder(body)
	var msg cluster.JoinNetworkMessage
	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}

	response, err := a.Cluster.AddDevice(msg.ImNewHere)
	if err != nil {
		log.Println("Error Joining Cluster")
	}

	return c.JSON(http.StatusOK, response)
}

//PeerUpdate receives new cluster info from the most recently registered peer
func (a *Application) peerUpdate(c echo.Context) error {
	log.Println("Receiving Update from New Peer")
	body := c.Request().Body

	var clustahUpdate cluster.PeerUpdateMessage
	err := json.NewDecoder(body).Decode(&clustahUpdate)
	if err != nil {
		log.Println("Failed to decode Cluster info from new peer")
		//TODO: Add Cluster Info Request to repair Cluster info
		return err
	}
	//TODO: Verify my presence in SlaveList
	//TODO: Verify my Master state
	//TODO: Inform Master of Bad Config

	//Update my cluster
	a.Cluster.LoadCluster(clustahUpdate.Cluster)

	log.Println("Peer Update Completed")
	return c.JSON(http.StatusOK, "Peer Update Successfully Received by : "+a.Cluster.Me.String())
}

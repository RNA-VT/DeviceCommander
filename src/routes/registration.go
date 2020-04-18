package routes

import (
	"encoding/json"
	"firecontroller/cluster"
	mc "firecontroller/microcontroller"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a *APIService) addRegistrationRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.POST("/peers", a.peerUpdate)
	api.POST("/join_network", a.joinNetwork)
}

func (a *APIService) joinNetwork(c echo.Context) error {
	log.Println("[master] Microcontroller asked to join cluster")

	body := c.Request().Body
	decoder := json.NewDecoder(body)
	var msg cluster.JoinNetworkMessage
	err := decoder.Decode(&msg)
	if err != nil {
		log.Println("Error decoding Request Body", err)
	}

	a.Cluster.AddMicrocontroller(msg.ImNewHere)
	a.Cluster.SendClusterUpdate([]mc.Config{
		msg.ImNewHere,
	})
	clusterUpdate := cluster.PeerUpdateMessage{
		Header:  a.Cluster.GetHeader(),
		Cluster: a.Cluster.GetConfig(),
	}
	return c.JSON(http.StatusOK, clusterUpdate)
}

//PeerUpdate receives new cluster info from the most recently registered peer
func (a *APIService) peerUpdate(c echo.Context) error {
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
	a.Cluster.Load(clustahUpdate.Cluster)

	log.Println("Peer Update Completed")
	return c.JSON(http.StatusOK, "Peer Update Successfully Received by : "+a.Cluster.Me.Name+" @ "+a.Cluster.Me.ToFullAddress())
}

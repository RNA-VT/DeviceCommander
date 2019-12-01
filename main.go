package main

/* Al useful imports */
import (
	"firecontroller/app"
	"firecontroller/nodecluster"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

/* The entry point for our System */
func main() {
	/* Parse the provided parameters on command line */
	makeMasterOnError := flag.Bool(
		"makeMasterOnError",
		false,
		"make this node master if unable to connect to the cluster ip provided.")
	clusterip := flag.String(
		"clusterip",
		"127.0.0.1:8001",
		"ip address of any node to connnect")
	myport := flag.String(
		"myport",
		"8001",
		"ip address to run this node on. default is 8001.")
	flag.Parse()

	/* Generate id for myself */
	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Intn(100)

	myIp, _ := net.InterfaceAddrs()
	// myIpString := strings.Split(myIp[0].String(), "/")[0]

	me := nodecluster.NodeInfo{
		NodeId:     myid,
		NodeIpAddr: myIp[0].String(),
		Port:       *myport}

	var cluster nodecluster.Cluster

	cluster.AddSlaveNode(me)
	cluster.MasterIp = clusterip

	app := app.Application{
		Cluster: cluster,
		Me:      me,
		Echo:    echo.New()}

	/* Try to connect to the cluster, and send request to cluster if able to connect */
	ableToConnect, newID := app.TestConnectToMaster(clusterip)

	fmt.Println("NEW ID: " + strconv.Itoa(newID))

	/*
	 * Listen for other incoming requests form other nodes to join cluster
	 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
	 */
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			app.Cluster.MasterNode = me
			fmt.Println("Will start this node as master.")
		}
		app.ConfigureRoutes()

	} else {
		fmt.Println("Quitting system. Set makeMasterOnError flag to make the node master.", myid)
	}
}

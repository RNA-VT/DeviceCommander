package main

/* Al useful imports */
import (
	"firecontroller/app"
	"firecontroller/nodecluster"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()
	defaultMaster := viper.GetString("GOFIRE_MASTER")
	port := viper.GetString("GOFIRE_PORT")
	fullHostname := viper.GetString("GOFIRE_HOST") + ":" + port
	fmt.Println(fullHostname)

	// /* Generate id for myself */
	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Intn(100)

	myIP, _ := net.InterfaceAddrs()

	me := nodecluster.NodeInfo{
		NodeID:     myid,
		NodeIPAddr: myIP[0].String(),
		Port:       port,
	}

	var cluster nodecluster.Cluster

	cluster.AddSlaveNode(me)

	app := app.Application{
		Cluster: cluster,
		Me:      me,
		Echo:    echo.New(),
	}

	/* Try to connect to the cluster, and send request to cluster if able to connect */
	ableToConnect := app.TestConnectToMaster(fullHostname)

	ableToConnect, assignedInfo := app.JoinNetwork(fullHostname)
	app.Me = assignedInfo

	// fmt.Println("NEW ID: " + strconv.Itoa(newID))

	/*f
	 * Listen for other incoming requests form other nodes to join cluster
	 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
	 */
	if ableToConnect || (!ableToConnect && defaultMaster == "TRUE") {
		if defaultMaster == "TRUE" {
			app.Cluster.MasterNode = me
			fmt.Println("Will start this node as master.")
		}
		app.ConfigureRoutes()

	} else {
		fmt.Println("Quitting system. Set makeMasterOnError flag to make the node master.", myid)
	}
}

func configureEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	viper.AutomaticEnv()

	viper.SetDefault("ENV", "local")
	viper.SetDefault("GOFIRE_MASTER", false)
	viper.SetDefault("GOFIRE_HOST", "127.0.0.1")
	viper.SetDefault("GOFIRE_PORT", 8001)
}

package main

/* Al useful imports */
import (
	"firecontroller/app"
	"firecontroller/nodecluster"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

/* The entry point for our System */
func main() {
	/* Load Config from Env Vars */
	configureEnvironment()

	gofireMaster := viper.GetBool("GOFIRE_MASTER")

	//Pick Listening Port
	port := "8001"
	host := "node1.mindshark.io"
	if viper.GetBool("GOFIRE_MASTER") {
		host = viper.GetString("GOFIRE_MASTER_HOST")
		port = viper.GetString("GOFIRE_MASTER_PORT")
	} else {
		host = viper.GetString("GOFIRE_HOST")
		port = viper.GetString("GOFIRE_PORT")
	}

	fullHostname := host + ":" + port
	masterHostname := viper.GetString("GOFIRE_MASTER_HOST") + ":" + viper.GetString("GOFIRE_MASTER_PORT")
	// /* Generate id for myself */
	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Intn(100)

	myIP, _ := net.InterfaceAddrs()

	me := nodecluster.NodeInfo{
		NodeID:     myid,
		NodeIPAddr: strings.Split(myIP[0].String(), "/")[0],
		Port:       port,
	}

	var cluster nodecluster.Cluster

	//Add this device to the slave list
	cluster.AddSlaveNode(me)

	app := app.Application{
		Cluster: cluster,
		Me:      me,
		Echo:    echo.New(),
	}

	//check to see if this instance is also the master
	if gofireMaster {
		log.Println("Master Mode Enabled!")
		app.Cluster.MasterNode = me
		app.Cluster.MasterNode.Port = port
		/*
		 * Listen for other incoming requests form other nodes to join cluster
		 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
		 */

	} else {
		log.Println("Slave Mode Enabled.")
		//Try and Connect to the Master
		err := app.TestConnectToMaster(masterHostname)
		if err != nil {
			log.Println("Failed to Reach Master Node: PANIC")
			//TODO: Add Retry or failover maybe? panic for now
			panic(err)
		}
		err = app.JoinNetwork(masterHostname, me)
		if err != nil {
			log.Println("Failed to Join Network: PANIC")
			panic(err)
		}
	}

	app.ConfigureRoutes(fullHostname)
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
	viper.SetDefault("GOFIRE_MASTER_PORT", 8000)
	viper.SetDefault("GOFIRE_MASTER_HOST", "127.0.0.1")
}

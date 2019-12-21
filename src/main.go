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
	gofireMaster := viper.GetBool("GOFIRE_MASTER")
	goFirePort := viper.GetString("GOFIRE_PORT")
	fullHostname := viper.GetString("GOFIRE_HOST") + ":" + goFirePort
	fmt.Println(fullHostname)

	// /* Generate id for myself */
	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Intn(100)

	myIP, _ := net.InterfaceAddrs()

	me := nodecluster.NodeInfo{
		NodeID:     myid,
		NodeIPAddr: myIP[0].String(),
		Port:       goFirePort,
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
		fmt.Println("Master Mode Enabled!")
		app.Cluster.MasterNode = me
		/*
		 * Listen for other incoming requests form other nodes to join cluster
		 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
		 */

	} else {
		fmt.Println("Slave Mode Enabled.")
		//Try and Connect to the Master
		_, err := app.TestConnectToMaster(fullHostname)
		if err != nil {
			fmt.Println("Failed to Reach Master Node: PANIC")
			//TODO: Add Retry or failover maybe? panic for now
			panic(err)
		}
		app.Me, err = app.JoinNetwork(fullHostname)
		if err != nil {
			fmt.Println("Failed to Join Network: PANIC")
			panic(err)
		}
	}

	fmt.Println("***************************************")
	fmt.Println("~Rejoice~ GoFire Lives Again! ~Rejoice~")
	fmt.Println("***************************************")

	app.ConfigureRoutes()
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

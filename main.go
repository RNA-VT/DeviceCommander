package main

/* Al useful imports */
import (
	"flag"
	"fmt"
	"net"
	"strings"

	// "strconv"
	"encoding/json"
	"firecontroller/nodecluster"
	"math/rand"
	"time"
	// "firecontroller/firecomponent"
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
	myid := rand.Intn(99999999)

	myIp, _ := net.InterfaceAddrs()
	me := nodecluster.NodeInfo{
		NodeId:     myid,
		NodeIpAddr: myIp[0].String(),
		Port:       *myport}

	// var slaveNodes []nodecluster.NodeInfo

	var cluster nodecluster.Cluster

	cluster.AddSlaveNode(me)
	cluster.MasterIp = clusterip

	cluster.PrintClusterInfo()

	// cluster := nodecluster.Cluster {
	//   // SlaveNodes: slaveNodes []nodecluster.NodeInfo,
	//   MasterIp: clusterip }

	dest := nodecluster.NodeInfo{
		NodeId:     -1,
		NodeIpAddr: strings.Split(*clusterip, ":")[0],
		Port:       strings.Split(*clusterip, ":")[1]}

	fmt.Println("My details:", me.String())

	/* Try to connect to the cluster, and send request to cluster if able to connect */
	ableToConnect := connectToCluster(me, dest)

	/*
	 * Listen for other incoming requests form other nodes to join cluster
	 * Note: We are not doing anything fancy right now to make this node as master. Not yet!
	 */
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("Will start this node as master.")
		}
		listenOnPort(me, cluster)
	} else {
		fmt.Println("Quitting system. Set makeMasterOnError flag to make the node master.", myid)
	}
}

func connectToCluster(me nodecluster.NodeInfo, dest nodecluster.NodeInfo) bool {
	/* connect to this socket details provided */
	connOut, err := net.DialTimeout(
		"tcp",
		dest.NodeIpAddr+":"+dest.Port,
		time.Duration(10)*time.Second)

	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("Couldn't connect to cluster.", me.NodeId)
			return false
		}
	} else {
		fmt.Println("Connected to cluster. Sending message to node.")
		text := "Hi nody.. please add me to the cluster.."
		requestMessage := nodecluster.GetAddToClusterMessage(me, dest, text)
		json.NewEncoder(connOut).Encode(&requestMessage)

		decoder := json.NewDecoder(connOut)
		var responseMessage nodecluster.AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("Got response:\n" + responseMessage.String())

		return true
	}
	return false
}

func listenOnPort(me nodecluster.NodeInfo, cluster nodecluster.Cluster) {
	/* Listen for incoming messages */
	ln, _ := net.Listen("tcp", fmt.Sprint(":"+me.Port))
	/* accept connection on port */
	/* not sure if looping infinetely on ln.Accept() is good idea */
	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", me.NodeId)
			}
		} else {

			var requestMessage nodecluster.AddToClusterMessage
			json.NewDecoder(connIn).Decode(&requestMessage)
			fmt.Println("Got request:\n" + requestMessage.String())

			text := "Sure buddy.. too easy.."
			responseMessage := nodecluster.GetAddToClusterMessage(me, requestMessage.Source, text)
			json.NewEncoder(connIn).Encode(&responseMessage)
			connIn.Close()
		}
	}
}

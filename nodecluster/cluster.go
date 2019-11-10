package nodecluster

import (

)

type Cluster struct {
  SlaveNodes []NodeInfo  `json:"slaveNodes"`
  MasterIp string `json:"masterIp"`
}

func (cluster *Cluster) AddSlaveNode(node NodeInfo) {

}

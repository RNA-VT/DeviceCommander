package nodecluster

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// "fmt"
// "net"
// "time"
// "encoding/json"

/* A standard format for a Request/Response for adding node to cluster */
type AddToClusterMessage struct {
	Source  NodeInfo
	Dest    NodeInfo
	Cluster Cluster
}

func SendMessageToAllNodes(cluster Cluster, message []byte) {
	var buf io.Reader
	resp, err := http.Post("http://example.com/upload", "image/jpeg", buf)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(body)

}

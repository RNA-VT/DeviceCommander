package rpc

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"

	log "github.com/sirupsen/logrus"
)

type DCServer struct {
	Port   int
	Server rpc.Server
}

func NewDCServer() DCServer {
	return DCServer{
		Port:   9999,
		Server: *rpc.NewServer(),
	}
}

func (s DCServer) RegisterHandler(handler interface{}) DCServer {

	return s
}

func (s DCServer) Listen() DCServer {
	err := s.Server.Register(&Device{})
	if err != nil {
		log.Fatal("Format of service Device isn't correct. ", err)
	}
	// Register a HTTP handler
	s.Server.HandleHTTP("/rpc", "/rpc/debug")

	// start the rpc server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Errorf("Error while starting rpc server: %+v", err)
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// http.ListenAndServe(":9999", nil)

	// defer listener.Close()

	// log.Infof("Start listening on port %d for RPC requests...", s.Port)
	// go func() {
	// 	for {
	// 		s.Server.Accept(listener)
	// 	}
	// }()

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}

	return s
}

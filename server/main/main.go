package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gardusig/grpc_service/server/internal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pandoraproto.RegisterPandoraServiceServer(s, &internal.PandoraService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

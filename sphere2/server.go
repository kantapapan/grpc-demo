package main

import (
	"log"
	"net"

	"github.com/kantapapan/grpc-demo/sphere2/service"

	pb "github.com/kantapapan/grpc-demo/sphere2/pb"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	s := &service.SphereService{}
	// Add Service
	pb.RegisterSphereServiceServer(server, s)
	server.Serve(listenPort)
}

package main

import (
	grpc "google.golang.org/grpc"
	"log"
	"net"
	pb "potato/demo/sample-server/grpc/go/proto"
	"potato/demo/sample-server/logic"
)

func main() {
	server := grpc.NewServer()
	if listener, err := net.Listen("tcp", ":5000"); err != nil {
		log.Fatalf("fail to listen : %v", err)
	} else {
		pb.RegisterSampleServiceServer(server, &logic.Service{})
		if err := server.Serve(listener); err != nil {
			log.Fatalf("fail to start server")
		} else {
			log.Println("service started")
		}

	}
}

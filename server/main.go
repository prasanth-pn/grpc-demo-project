package main

import (
	"log"
	"net"
pb "github.com/prasanth-pn/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9090"
)
type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the serer %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer,&helloServer{})
	log.Printf("server started at : %v",lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %v", err)

	}
}

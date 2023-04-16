package main

import (
	"log"
	"time"

	pb "github.com/prasanth-pn/grpc-demo/proto"
)

func (s *helloServer)SayHelloServerStreaming(req *pb.NamesList,stream pb.GreetService_SayHelloServerStreamingServer)error{
log.Printf("got request with Names:%v",req.Name)
for _,name:=range req.Name{
	res:=&pb.HelloResponse{
		Message: "hello"+name,
	}
	if err:=stream.Send(res);err!=nil{
		return err
	}
	time.Sleep(2* time.Second)
}
return nil
}


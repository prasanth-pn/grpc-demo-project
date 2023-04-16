package main

import (
	"log"
pb "github.com/prasanth-pn/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const(
	port= ":9090"
)
func main(){
	conn,err:=grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("did not connecg :%v",err)

	}
	defer conn.Close()
	client:=pb.NewGreetServiceClient(conn)
	names:=&pb.NamesList{
		Name: []string{"prasanth","akhil","kunalkushwaha"},
	}
	// callSayHello(client)
	//callSayHelloSeverStream(client,names)
	//callSayHelloClientStream(client,names)
	callHelloBidirectionalStream(client,names)
}
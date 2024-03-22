package main

import (
	"github.com/abbasfisal/grpc-tut/hello/hellopb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	//
	s := grpc.NewServer()

	//bind server to our pb
	hellopb.RegisterHelloServiceServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}

type Server struct {
	hellopb.UnimplementedHelloServiceServer
}

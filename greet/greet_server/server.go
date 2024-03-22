package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {

	fmt.Printf("greet fun was invoked %v ", req)

	f := req.GetGreeting().FirstName
	r := "hello " + f + "  , welcome "
	res := &greetpb.GreetingResponse{Result: r}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("greet server is running :50051")

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

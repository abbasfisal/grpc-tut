package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
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

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	fmt.Printf("GreetManyTimes fun was invoked %v \n", req)

	f := req.GetGreeting().FirstName

	for i := 0; i < 10; i++ {
		res := "hello form server streaming : " + strconv.Itoa(i) + " | welcome : " + f + "\n"

		response := greetpb.GreetManyTimesResponse{Result: res}

		stream.Send(&response)
		time.Sleep(1 * time.Second)
	}
	return nil

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

package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	//connect to server
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("client created : ", c)

	req := &greetpb.GreetingRequest{Greeting: &greetpb.Greeting{
		FirstName: "Mohammad",
		LastName:  "Ali",
	}}

	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response from greet : ", response.Result)

}

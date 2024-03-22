package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
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

	//-- unary
	doUnary(c)

	//-- server streaming
	doServerStreaming(c)

	//-- client streaming
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {

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

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{Greeting: &greetpb.Greeting{
		FirstName: "MohammadReza",
		LastName:  "Momeni",
	}}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("response from GreetManyTimes: %v ", msg.GetResult())

	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("client streaming RPC ... ")

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	req := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "abbas",
		}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "naser",
		}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "shahla",
		}},
		&greetpb.LongGreetRequest{Greeting: &greetpb.Greeting{
			FirstName: "mahla",
		}},
	}

	for _, r := range req {
		fmt.Printf("sending request %v \n", r)
		stream.Send(r)
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Long Greet Response : %v", response.GetResult())
}
